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

type PermissionService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.PermissionResponse, error)
	Create(ctx context.Context, req requests.CreatePermissionRequest) (responses.PermissionResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdatePermissionRequest) (responses.PermissionResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type permissionService struct {
	repo       repos.PermissionRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewPermissionService(r repos.PermissionRepository, o helpers.IDObfuscator, v *validator.Validate) PermissionService {
	return &permissionService{repo: r, obfuscator: o, validator: v}
}

func (s *permissionService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "name", "module")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.PermissionResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Permissions retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *permissionService) GetByID(ctx context.Context, hashID string) (responses.PermissionResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.PermissionResponse{}, errors.New("invalid permission id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.PermissionResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *permissionService) Create(ctx context.Context, req requests.CreatePermissionRequest) (responses.PermissionResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.PermissionResponse{}, err
	}

	permission := &models.Permission{
		Name:        req.Name,
		Module:      req.Module,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, permission); err != nil {
		return responses.PermissionResponse{}, err
	}

	return s.mapToResponse(*permission), nil
}

func (s *permissionService) Update(ctx context.Context, hashID string, req requests.UpdatePermissionRequest) (responses.PermissionResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.PermissionResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.PermissionResponse{}, errors.New("invalid permission id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.PermissionResponse{}, err
	}

	if req.Name != "" {
		model.Name = req.Name
	}
	if req.Module != "" {
		model.Module = req.Module
	}
	if req.Description != "" {
		model.Description = req.Description
	}

	if err := s.repo.Update(ctx, model); err != nil {
		return responses.PermissionResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *permissionService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid permission id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *permissionService) mapToResponse(m models.Permission) responses.PermissionResponse {
	id, _ := s.obfuscator.Encode(m.ID)
	return responses.PermissionResponse{
		ID:          id,
		Name:        m.Name,
		Module:      m.Module,
		Description: m.Description,
		CreatedBy:   m.CreatedBy,
		UpdatedBy:   m.UpdatedBy,
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
