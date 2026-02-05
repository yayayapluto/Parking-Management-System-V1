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

type ZoneTypeService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.ZoneTypeResponse, error)
	Create(ctx context.Context, req requests.CreateZoneTypeRequest) (responses.ZoneTypeResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateZoneTypeRequest) (responses.ZoneTypeResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type zoneTypeService struct {
	repo       repos.ZoneTypeRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func (s *zoneTypeService) GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, p, "name")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.ZoneTypeResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Zone types retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(p, total, url),
	}, nil
}

func (s *zoneTypeService) GetByID(ctx context.Context, hashedID string) (responses.ZoneTypeResponse, error) {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.ZoneTypeResponse{}, errors.New("invalid zone type id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ZoneTypeResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *zoneTypeService) Create(ctx context.Context, req requests.CreateZoneTypeRequest) (responses.ZoneTypeResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ZoneTypeResponse{}, err
	}

	zoneType := &models.ZoneType{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, zoneType); err != nil {
		return responses.ZoneTypeResponse{}, err
	}

	return s.mapToResponse(*zoneType), nil
}

func (s *zoneTypeService) Update(ctx context.Context, hashedID string, req requests.UpdateZoneTypeRequest) (responses.ZoneTypeResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ZoneTypeResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.ZoneTypeResponse{}, errors.New("invalid zone type id")
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ZoneTypeResponse{}, err
	}

	if req.Name != "" {
		existing.Name = req.Name
	}

	if req.Description != "" {
		existing.Description = req.Description
	}

	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.ZoneTypeResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *zoneTypeService) Delete(ctx context.Context, hashedID string) error {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return errors.New("invalid zone type id")
	}

	return s.repo.Delete(ctx, id)
}

func NewZoneTypeService(r repos.ZoneTypeRepository, o helpers.IDObfuscator, v *validator.Validate) ZoneTypeService {
	return &zoneTypeService{
		repo:       r,
		obfuscator: o,
		validator:  v,
	}
}

func (s *zoneTypeService) mapToResponse(m models.ZoneType) responses.ZoneTypeResponse {
	hId, _ := s.obfuscator.Encode(m.ID)
	return responses.ZoneTypeResponse{
		ID:          hId,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
