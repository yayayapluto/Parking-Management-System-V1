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

type ShiftReportService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.ShiftReportResponse, error)
	Create(ctx context.Context, req requests.CreateShiftReportRequest) (responses.ShiftReportResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateShiftReportRequest) (responses.ShiftReportResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type shiftReportService struct {
	repo       repos.ShiftReportRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewShiftReportService(r repos.ShiftReportRepository, o helpers.IDObfuscator, v *validator.Validate) ShiftReportService {
	return &shiftReportService{repo: r, obfuscator: o, validator: v}
}

func (s *shiftReportService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req)
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.ShiftReportResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Shift reports retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *shiftReportService) GetByID(ctx context.Context, hashID string) (responses.ShiftReportResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ShiftReportResponse{}, errors.New("invalid shift report id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ShiftReportResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *shiftReportService) Create(ctx context.Context, req requests.CreateShiftReportRequest) (responses.ShiftReportResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ShiftReportResponse{}, err
	}

	userID, err := s.obfuscator.Decode(req.UserID)
	if err != nil {
		return responses.ShiftReportResponse{}, errors.New("invalid user id")
	}

	// Calculate total revenue
	totalRevenue := req.TotalCash + req.TotalQRIS

	// Calculate discrepancy
	discrepancy := req.ClosingCash - req.OpeningCash - totalRevenue

	report := &models.ShiftReport{
		UserID:            userID,
		ShiftStart:        req.ShiftStart,
		ShiftEnd:          req.ShiftEnd,
		OpeningCash:       req.OpeningCash,
		ClosingCash:       req.ClosingCash,
		TotalTransactions: req.TotalTransactions,
		TotalCash:         req.TotalCash,
		TotalQRIS:         req.TotalQRIS,
		TotalRevenue:      totalRevenue,
		Discrepancy:       discrepancy,
		Notes:             req.Notes,
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return responses.ShiftReportResponse{}, err
	}

	return s.mapToResponse(*report), nil
}

func (s *shiftReportService) Update(ctx context.Context, hashID string, req requests.UpdateShiftReportRequest) (responses.ShiftReportResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ShiftReportResponse{}, errors.New("invalid shift report id")
	}

	report, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ShiftReportResponse{}, err
	}

	if req.ClosingCash != nil {
		report.ClosingCash = *req.ClosingCash
	}
	if req.TotalTransactions != nil {
		report.TotalTransactions = *req.TotalTransactions
	}
	if req.TotalCash != nil {
		report.TotalCash = *req.TotalCash
	}
	if req.TotalQRIS != nil {
		report.TotalQRIS = *req.TotalQRIS
	}
	if req.Notes != "" {
		report.Notes = req.Notes
	}

	// Recalculate totals
	report.TotalRevenue = report.TotalCash + report.TotalQRIS
	report.Discrepancy = report.ClosingCash - report.OpeningCash - report.TotalRevenue

	if err := s.repo.Update(ctx, report); err != nil {
		return responses.ShiftReportResponse{}, err
	}

	return s.mapToResponse(*report), nil
}

func (s *shiftReportService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid shift report id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *shiftReportService) mapToResponse(m models.ShiftReport) responses.ShiftReportResponse {
	// Pecah dulu semua yang pake obfuscator karena return (string, error)
	encodedID, _ := s.obfuscator.Encode(m.ID)
	encodedUserID, _ := s.obfuscator.Encode(m.UserID)

	/**
	Unknown field 'UserID' in struct literal
	dll
	*/
	return responses.ShiftReportResponse{
		ID:                encodedID,
		UserID:            encodedUserID, // Sekarang aman, udah jadi single value string
		ShiftStart:        m.ShiftStart.Format("2006-01-02 15:04:05"),
		ShiftEnd:          m.ShiftEnd.Format("2006-01-02 15:04:05"),
		OpeningCash:       m.OpeningCash,
		ClosingCash:       m.ClosingCash,
		TotalTransactions: m.TotalTransactions,
		TotalCash:         m.TotalCash,
		TotalQRIS:         m.TotalQRIS,
		TotalRevenue:      m.TotalRevenue,
		Discrepancy:       m.Discrepancy,
		Notes:             m.Notes,
		CreatedAt:         m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
