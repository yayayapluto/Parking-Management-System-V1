package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/pkg/helpers"
)

type UserActivityLogService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.UserActivityLogResponse, error)
	Create(ctx context.Context, req requests.CreateUserActivityLogRequest) (responses.UserActivityLogResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type userActivityLogService struct {
	repo       repos.UserActivityLogRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewUserActivityLogService(r repos.UserActivityLogRepository, o helpers.IDObfuscator, v *validator.Validate) UserActivityLogService {
	return &userActivityLogService{repo: r, obfuscator: o, validator: v}
}

func (s *userActivityLogService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "action", "module")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.UserActivityLogResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "User activity logs retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *userActivityLogService) GetByID(ctx context.Context, hashID string) (responses.UserActivityLogResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.UserActivityLogResponse{}, errors.New("invalid user activity log id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.UserActivityLogResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *userActivityLogService) Create(ctx context.Context, req requests.CreateUserActivityLogRequest) (responses.UserActivityLogResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.UserActivityLogResponse{}, err
	}

	userID, err := s.obfuscator.Decode(req.UserID)
	if err != nil {
		return responses.UserActivityLogResponse{}, errors.New("invalid user id")
	}

	log := &models.UserActivityLog{
		UserID:      userID,
		Action:      req.Action,
		Module:      req.Module,
		Description: req.Description,
		IPAddress:   req.IPAddress,
		UserAgent:   req.UserAgent,
		RequestData: pq.StringArray(req.RequestData),
	}

	if err := s.repo.Create(ctx, log); err != nil {
		return responses.UserActivityLogResponse{}, err
	}

	return s.mapToResponse(*log), nil
}

func (s *userActivityLogService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid user activity log id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *userActivityLogService) mapToResponse(m models.UserActivityLog) responses.UserActivityLogResponse {
	id, _ := s.obfuscator.Encode(m.ID)
	uID, _ := s.obfuscator.Encode(m.UserID)

	return responses.UserActivityLogResponse{
		ID:          id,
		UserID:      uID,
		Action:      m.Action,
		Module:      m.Module,
		Description: m.Description,
		IPAddress:   m.IPAddress,
		UserAgent:   m.UserAgent,
		RequestData: []string(m.RequestData),
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
