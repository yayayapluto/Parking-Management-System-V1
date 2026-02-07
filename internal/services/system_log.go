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

type SystemLogService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.SystemLogResponse, error)
	Create(ctx context.Context, req requests.CreateSystemLogRequest) (responses.SystemLogResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type systemLogService struct {
	repo       repos.SystemLogRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewSystemLogService(r repos.SystemLogRepository, o helpers.IDObfuscator, v *validator.Validate) SystemLogService {
	return &systemLogService{repo: r, obfuscator: o, validator: v}
}

func (s *systemLogService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "service", "message")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.SystemLogResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "System logs retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *systemLogService) GetByID(ctx context.Context, hashID string) (responses.SystemLogResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.SystemLogResponse{}, errors.New("invalid system log id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.SystemLogResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *systemLogService) Create(ctx context.Context, req requests.CreateSystemLogRequest) (responses.SystemLogResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.SystemLogResponse{}, err
	}

	log := &models.SystemLog{
		Level:      req.Level,
		Service:    req.Service,
		Message:    req.Message,
		Context:    req.Context,
		StackTrace: req.StackTrace,
	}

	if err := s.repo.Create(ctx, log); err != nil {
		return responses.SystemLogResponse{}, err
	}

	return s.mapToResponse(*log), nil
}

func (s *systemLogService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid system log id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *systemLogService) mapToResponse(m models.SystemLog) responses.SystemLogResponse {
	encodedID, _ := s.obfuscator.Encode(m.ID)
	return responses.SystemLogResponse{
		ID:         encodedID,
		Level:      m.Level,
		Service:    m.Service,
		Message:    m.Message,
		Context:    m.Context,
		StackTrace: m.StackTrace,
		CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
