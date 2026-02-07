package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/pkg/helpers"
	"time"
)

type DailySettlementService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.DailySettlementResponse, error)
	Create(ctx context.Context, req requests.CreateDailySettlementRequest) (responses.DailySettlementResponse, error)
	Reconcile(ctx context.Context, hashedID string, req requests.ReconcileDailySettlementRequest) (responses.DailySettlementResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type dailySettlementService struct {
	repo       repos.DailySettlementRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewDailySettlementService(r repos.DailySettlementRepository, o helpers.IDObfuscator, v *validator.Validate) DailySettlementService {
	return &dailySettlementService{repo: r, obfuscator: o, validator: v}
}

func (s *dailySettlementService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req)
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.DailySettlementResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Daily settlements retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *dailySettlementService) GetByID(ctx context.Context, hashID string) (responses.DailySettlementResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.DailySettlementResponse{}, errors.New("invalid daily settlement id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.DailySettlementResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *dailySettlementService) Create(ctx context.Context, req requests.CreateDailySettlementRequest) (responses.DailySettlementResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.DailySettlementResponse{}, err
	}

	settlement := &models.DailySettlement{
		Date:              req.Date,
		TotalTransactions: req.TotalTransactions,
		TotalRevenue:      req.TotalRevenue,
		TotalCash:         req.TotalCash,
		TotalQRIS:         req.TotalQRIS,
		TotalRefunds:      req.TotalRefunds,
		TotalLostTickets:  req.TotalLostTickets,
		Variance:          0,
	}

	if err := s.repo.Create(ctx, settlement); err != nil {
		return responses.DailySettlementResponse{}, err
	}

	return s.mapToResponse(*settlement), nil
}

func (s *dailySettlementService) Reconcile(ctx context.Context, hashID string, req requests.ReconcileDailySettlementRequest) (responses.DailySettlementResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.DailySettlementResponse{}, errors.New("invalid daily settlement id")
	}

	settlement, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.DailySettlementResponse{}, err
	}

	reconciledByID, err := s.obfuscator.Decode(req.ReconciledBy)
	if err != nil {
		return responses.DailySettlementResponse{}, errors.New("invalid reconciled_by id")
	}

	// Calculate variance
	expectedTotal := settlement.TotalCash + settlement.TotalQRIS
	variance := expectedTotal - settlement.TotalRevenue

	now := time.Now()
	settlement.Variance = variance
	settlement.ReconciledBy = &reconciledByID
	settlement.ReconciledAt = &now

	if err := s.repo.Update(ctx, settlement); err != nil {
		return responses.DailySettlementResponse{}, err
	}

	return s.mapToResponse(*settlement), nil
}

func (s *dailySettlementService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid daily settlement id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *dailySettlementService) mapToResponse(m models.DailySettlement) responses.DailySettlementResponse {
	reconciledBy := ""
	if m.ReconciledBy != nil {
		encoded, err := s.obfuscator.Encode(*m.ReconciledBy)
		if err == nil {
			reconciledBy = encoded
		}
	}

	reconciledAt := ""
	if m.ReconciledAt != nil {
		reconciledAt = m.ReconciledAt.Format("2006-01-02 15:04:05")
	}

	id, _ := s.obfuscator.Encode(m.ID)

	return responses.DailySettlementResponse{
		ID:                id,
		Date:              m.Date.Format("2006-01-02"),
		TotalTransactions: m.TotalTransactions,
		TotalRevenue:      m.TotalRevenue,
		TotalCash:         m.TotalCash,
		TotalQris:         m.TotalQRIS,
		TotalRefunds:      m.TotalRefunds,
		TotalLostTickets:  m.TotalLostTickets,
		Variance:          m.Variance,
		ReconciledBy:      reconciledBy,
		ReconciledAt:      reconciledAt,
		CreatedAt:         m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
