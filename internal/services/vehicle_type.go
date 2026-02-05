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

type VehicleTypeService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.VehicleTypeResponse, error)
	Create(ctx context.Context, req requests.CreateVehicleTypeRequest) (responses.VehicleTypeResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateVehicleTypeRequest) (responses.VehicleTypeResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type vehicleTypeService struct {
	repo       repos.VehicleTypeRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewVehicleTypeService(r repos.VehicleTypeRepository, o helpers.IDObfuscator, v *validator.Validate) VehicleTypeService {
	return &vehicleTypeService{repo: r, obfuscator: o, validator: v}
}

// --- CORE METHODS ---

func (s *vehicleTypeService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	// Panggil repo dengan searchable columns: code dan name
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "code", "name")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.VehicleTypeResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Vehicle types retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *vehicleTypeService) GetByID(ctx context.Context, hashID string) (responses.VehicleTypeResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.VehicleTypeResponse{}, errors.New("invalid vehicle type id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.VehicleTypeResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *vehicleTypeService) Create(ctx context.Context, req requests.CreateVehicleTypeRequest) (responses.VehicleTypeResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.VehicleTypeResponse{}, err
	}

	vehicleType := models.VehicleType{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}

	if err := s.repo.Create(ctx, &vehicleType); err != nil {
		return responses.VehicleTypeResponse{}, err
	}

	return s.mapToResponse(vehicleType), nil
}

func (s *vehicleTypeService) Update(ctx context.Context, hashID string, req requests.UpdateVehicleTypeRequest) (responses.VehicleTypeResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.VehicleTypeResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.VehicleTypeResponse{}, errors.New("invalid id format")
	}

	// 1. Ambil data lama
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.VehicleTypeResponse{}, err
	}

	// 2. Patching data (Hanya yang dikirim di request)
	if req.Code != "" {
		existing.Code = req.Code
	}
	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	// 3. Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.VehicleTypeResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *vehicleTypeService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// --- PRIVATE HELPER ---

func (s *vehicleTypeService) mapToResponse(m models.VehicleType) responses.VehicleTypeResponse {
	hID, _ := s.obfuscator.Encode(m.ID)
	return responses.VehicleTypeResponse{
		ID:          hID,
		Code:        m.Code,
		Name:        m.Name,
		Description: m.Description,
		IsActive:    m.IsActive,
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
