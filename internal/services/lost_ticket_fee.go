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
)

type LostTicketFeeService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.LostTicketFeeResponse, error)
	Create(ctx context.Context, req requests.CreateLostTicketFeeRequest) (responses.LostTicketFeeResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type lostTicketFeeService struct {
	repo       repos.LostTicketFeeRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewLostTicketFeeService(r repos.LostTicketFeeRepository, o helpers.IDObfuscator, v *validator.Validate) LostTicketFeeService {
	return &lostTicketFeeService{repo: r, obfuscator: o, validator: v}
}

func (s *lostTicketFeeService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req)
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.LostTicketFeeResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Lost ticket fees retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *lostTicketFeeService) GetByID(ctx context.Context, hashID string) (responses.LostTicketFeeResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.LostTicketFeeResponse{}, errors.New("invalid lost ticket fee id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.LostTicketFeeResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *lostTicketFeeService) Create(ctx context.Context, req requests.CreateLostTicketFeeRequest) (responses.LostTicketFeeResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.LostTicketFeeResponse{}, err
	}

	transactionID, err := s.obfuscator.Decode(req.TransactionID)
	if err != nil {
		return responses.LostTicketFeeResponse{}, errors.New("invalid transaction id")
	}

	processedByID, err := s.obfuscator.Decode(req.ProcessedBy)
	if err != nil {
		return responses.LostTicketFeeResponse{}, errors.New("invalid processed_by id")
	}

	// Calculate total charged: original fee + lost ticket fee
	totalCharged := req.OriginalFee + req.LostTicketFee

	fee := &models.LostTicketFee{
		TransactionID: transactionID,
		OriginalFee:   req.OriginalFee,
		LostTicketFee: req.LostTicketFee,
		TotalCharged:  totalCharged,
		ProcessedBy:   processedByID,
	}

	if err := s.repo.Create(ctx, fee); err != nil {
		return responses.LostTicketFeeResponse{}, err
	}

	return s.mapToResponse(*fee), nil
}

func (s *lostTicketFeeService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid lost ticket fee id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *lostTicketFeeService) mapToResponse(m models.LostTicketFee) responses.LostTicketFeeResponse {
	id, _ := s.obfuscator.Encode(m.ID)
	transactionID, _ := s.obfuscator.Encode(m.TransactionID)
	processedBy, _ := s.obfuscator.Encode(m.ProcessedBy)
	return responses.LostTicketFeeResponse{
		ID:            id,
		TransactionID: transactionID,
		OriginalFee:   m.OriginalFee,
		LostTicketFee: m.LostTicketFee,
		TotalCharged:  m.TotalCharged,
		ProcessedBy:   processedBy,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
