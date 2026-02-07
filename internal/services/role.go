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

type RoleService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.RoleResponse, error)
	Create(ctx context.Context, req requests.CreateRoleRequest) (responses.RoleResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateRoleRequest) (responses.RoleResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type roleService struct {
	repo       repos.RoleRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewRoleService(r repos.RoleRepository, o helpers.IDObfuscator, v *validator.Validate) RoleService {
	return &roleService{repo: r, obfuscator: o, validator: v}
}

// GetAll retrieves all roles with pagination
func (s *roleService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "name")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.RoleResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Roles retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

// GetByID retrieves a role by hashed ID
func (s *roleService) GetByID(ctx context.Context, hashID string) (responses.RoleResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.RoleResponse{}, errors.New("invalid role id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.RoleResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

// Create creates a new role
func (s *roleService) Create(ctx context.Context, req requests.CreateRoleRequest) (responses.RoleResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.RoleResponse{}, err
	}

	// Decode PermissionID
	permissionID, err := s.obfuscator.Decode(req.PermissionID)
	if err != nil {
		return responses.RoleResponse{}, errors.New("invalid permission id")
	}

	role := models.Role{
		Name:        req.Name,
		PermissionID: permissionID,
		CreatedBy:   req.CreatedBy,
	}

	if err := s.repo.Create(ctx, &role); err != nil {
		return responses.RoleResponse{}, err
	}

	return s.mapToResponse(role), nil
}

// Update updates an existing role
func (s *roleService) Update(ctx context.Context, hashID string, req requests.UpdateRoleRequest) (responses.RoleResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.RoleResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.RoleResponse{}, errors.New("invalid id format")
	}

	// Get existing role
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.RoleResponse{}, err
	}

	// Patch data (only update provided fields)
	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.PermissionID != "" {
		permissionID, err := s.obfuscator.Decode(req.PermissionID)
		if err != nil {
			return responses.RoleResponse{}, errors.New("invalid permission id")
		}
		existing.PermissionID = permissionID
	}
	if req.UpdatedBy != "" {
		existing.UpdatedBy = req.UpdatedBy
	}

	// Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.RoleResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

// Delete deletes a role
func (s *roleService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// mapToResponse converts a Role model to RoleResponse DTO
func (s *roleService) mapToResponse(m models.Role) responses.RoleResponse {
	hID, _ := s.obfuscator.Encode(m.ID)
	hPermissionID, _ := s.obfuscator.Encode(m.PermissionID)

	return responses.RoleResponse{
		ID:          hID,
		Name:        m.Name,
		PermissionID: hPermissionID,
		CreatedBy:   m.CreatedBy,
		UpdatedBy:   m.UpdatedBy,
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
