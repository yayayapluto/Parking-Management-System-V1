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

type ManualCorrectionService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.ManualCorrectionResponse, error)
	Create(ctx context.Context, req requests.CreateManualCorrectionRequest) (responses.ManualCorrectionResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type manualCorrectionService struct {
	repo       repos.ManualCorrectionRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewManualCorrectionService(r repos.ManualCorrectionRepository, o helpers.IDObfuscator, v *validator.Validate) ManualCorrectionService {
	return &manualCorrectionService{repo: r, obfuscator: o, validator: v}
}

func (s *manualCorrectionService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "field_corrected", "reason")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.ManualCorrectionResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Manual corrections retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *manualCorrectionService) GetByID(ctx context.Context, hashID string) (responses.ManualCorrectionResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ManualCorrectionResponse{}, errors.New("invalid manual correction id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ManualCorrectionResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *manualCorrectionService) Create(ctx context.Context, req requests.CreateManualCorrectionRequest) (responses.ManualCorrectionResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ManualCorrectionResponse{}, err
	}

	transactionID, err := s.obfuscator.Decode(req.TransactionID)
	if err != nil {
		return responses.ManualCorrectionResponse{}, errors.New("invalid transaction id")
	}

	correctedByID, err := s.obfuscator.Decode(req.CorrectedBy)
	if err != nil {
		return responses.ManualCorrectionResponse{}, errors.New("invalid corrected_by id")
	}

	correction := &models.ManualCorrection{
		TransactionID:  transactionID,
		FieldCorrected: req.FieldCorrected,
		OldValue:       req.OldValue,
		NewValue:       req.NewValue,
		Reason:         req.Reason,
		CorrectedBy:    correctedByID,
	}

	if err := s.repo.Create(ctx, correction); err != nil {
		return responses.ManualCorrectionResponse{}, err
	}

	return s.mapToResponse(*correction), nil
}

func (s *manualCorrectionService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid manual correction id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *manualCorrectionService) mapToResponse(m models.ManualCorrection) responses.ManualCorrectionResponse {
	id, _ := s.obfuscator.Encode(m.ID)
	transactionID, _ := s.obfuscator.Encode(m.TransactionID)
	correctedBy, _ := s.obfuscator.Encode(m.CorrectedBy)
	return responses.ManualCorrectionResponse{
		ID:             id,
		TransactionID:  transactionID,
		FieldCorrected: m.FieldCorrected,
		OldValue:       m.OldValue,
		NewValue:       m.NewValue,
		Reason:         m.Reason,
		CorrectedBy:    correctedBy,
		CreatedAt:      m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
