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

type OCRLogService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.OCRLogResponse, error)
	Create(ctx context.Context, req requests.CreateOCRLogRequest) (responses.OCRLogResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type ocrLogService struct {
	repo       repos.OCRLogRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewOCRLogService(r repos.OCRLogRepository, o helpers.IDObfuscator, v *validator.Validate) OCRLogService {
	return &ocrLogService{repo: r, obfuscator: o, validator: v}
}

func (s *ocrLogService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "ocr_type", "detected_text", "status")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.OCRLogResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "OCR logs retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *ocrLogService) GetByID(ctx context.Context, hashID string) (responses.OCRLogResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.OCRLogResponse{}, errors.New("invalid ocr log id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.OCRLogResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *ocrLogService) Create(ctx context.Context, req requests.CreateOCRLogRequest) (responses.OCRLogResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.OCRLogResponse{}, err
	}

	transactionID, err := s.obfuscator.Decode(req.TransactionID)
	if err != nil {
		return responses.OCRLogResponse{}, errors.New("invalid transaction id")
	}

	log := &models.OCRLog{
		TransactionID:    transactionID,
		OCRType:          req.OCRType,
		ImagePath:        req.ImagePath,
		RawOCRResult:     req.RawOCRResult,
		DetectedText:     req.DetectedText,
		ValidatedPlate:   req.ValidatedPlate,
		ConfidenceScore:  req.ConfidenceScore,
		ProcessingTimeMs: req.ProcessingTimeMs,
		Status:           req.Status,
		ErrorMessage:     req.ErrorMessage,
	}

	if err := s.repo.Create(ctx, log); err != nil {
		return responses.OCRLogResponse{}, err
	}

	return s.mapToResponse(*log), nil
}

func (s *ocrLogService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid ocr log id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *ocrLogService) mapToResponse(m models.OCRLog) responses.OCRLogResponse {
	id, _ := s.obfuscator.Encode(m.ID)
	transactionID, _ := s.obfuscator.Encode(m.TransactionID)
	return responses.OCRLogResponse{
		ID:               id,
		TransactionID:    transactionID,
		OCRType:          m.OCRType,
		ImagePath:        m.ImagePath,
		RawOCRResult:     m.RawOCRResult,
		DetectedText:     m.DetectedText,
		ValidatedPlate:   m.ValidatedPlate,
		ConfidenceScore:  m.ConfidenceScore,
		ProcessingTimeMs: m.ProcessingTimeMs,
		Status:           m.Status,
		ErrorMessage:     m.ErrorMessage,
		CreatedAt:        m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
