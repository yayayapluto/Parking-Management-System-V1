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

type ReportCacheService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.ReportCacheResponse, error)
	Create(ctx context.Context, req requests.CreateReportCacheRequest) (responses.ReportCacheResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateReportCacheRequest) (responses.ReportCacheResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type reportCacheService struct {
	repo       repos.ReportCacheRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewReportCacheService(r repos.ReportCacheRepository, o helpers.IDObfuscator, v *validator.Validate) ReportCacheService {
	return &reportCacheService{repo: r, obfuscator: o, validator: v}
}

func (s *reportCacheService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "report_type", "cache_key")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.ReportCacheResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Report caches retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *reportCacheService) GetByID(ctx context.Context, hashID string) (responses.ReportCacheResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ReportCacheResponse{}, errors.New("invalid report cache id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ReportCacheResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *reportCacheService) Create(ctx context.Context, req requests.CreateReportCacheRequest) (responses.ReportCacheResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ReportCacheResponse{}, err
	}

	expiresAt, err := time.Parse(time.RFC3339, req.ExpiresAt)
	if err != nil {
		return responses.ReportCacheResponse{}, errors.New("invalid expires_at format, expected RFC3339")
	}

	cache := &models.ReportCache{
		ReportType: req.ReportType,
		CacheKey:   req.CacheKey,
		CacheData:  req.CacheData,
		ExpiresAt:  expiresAt,
	}

	if err := s.repo.Create(ctx, cache); err != nil {
		return responses.ReportCacheResponse{}, err
	}

	return s.mapToResponse(*cache), nil
}

func (s *reportCacheService) Update(ctx context.Context, hashID string, req requests.UpdateReportCacheRequest) (responses.ReportCacheResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ReportCacheResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ReportCacheResponse{}, errors.New("invalid report cache id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ReportCacheResponse{}, err
	}

	if req.ReportType != "" {
		model.ReportType = req.ReportType
	}
	if req.CacheKey != "" {
		model.CacheKey = req.CacheKey
	}
	if req.CacheData != "" {
		model.CacheData = req.CacheData
	}
	if req.ExpiresAt != "" {
		expiresAt, err := time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return responses.ReportCacheResponse{}, errors.New("invalid expires_at format, expected RFC3339")
		}
		model.ExpiresAt = expiresAt
	}

	if err := s.repo.Update(ctx, model); err != nil {
		return responses.ReportCacheResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *reportCacheService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid report cache id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *reportCacheService) mapToResponse(m models.ReportCache) responses.ReportCacheResponse {
	id, _ := s.obfuscator.Encode(m.ID)
	return responses.ReportCacheResponse{
		ID:         id,
		ReportType: m.ReportType,
		CacheKey:   m.CacheKey,
		CacheData:  m.CacheData,
		ExpiresAt:  m.ExpiresAt.Format("2006-01-02 15:04:05"),
		CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
