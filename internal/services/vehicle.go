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

type VehicleService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.VehicleResponse, error)
	Create(ctx context.Context, req requests.CreateVehicleRequest) (responses.VehicleResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateVehicleRequest) (responses.VehicleResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type vehicleService struct {
	repo       repos.VehicleRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewVehicleService(r repos.VehicleRepository, o helpers.IDObfuscator, v *validator.Validate) VehicleService {
	return &vehicleService{repo: r, obfuscator: o, validator: v}
}

// --- CORE METHODS ---

func (s *vehicleService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "plate_number", "description")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.VehicleResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Vehicles retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *vehicleService) GetByID(ctx context.Context, hashID string) (responses.VehicleResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.VehicleResponse{}, errors.New("invalid vehicle id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.VehicleResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *vehicleService) Create(ctx context.Context, req requests.CreateVehicleRequest) (responses.VehicleResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.VehicleResponse{}, err
	}

	// Decode VehicleTypeID
	vehicleTypeID, err := s.obfuscator.Decode(req.VehicleTypeID)
	if err != nil {
		return responses.VehicleResponse{}, errors.New("invalid vehicle type id")
	}

	var customerID *uint
	if req.CustomerID != "" {
		id, err := s.obfuscator.Decode(req.CustomerID)
		if err != nil {
			return responses.VehicleResponse{}, errors.New("invalid customer id")
		}
		customerID = &id
	}

	vehicle := models.Vehicle{
		CustomerID:    customerID,
		VehicleTypeID: vehicleTypeID,
		PlateNumber:   req.PlateNumber,
		Description:   req.Description,
	}

	if err := s.repo.Create(ctx, &vehicle); err != nil {
		return responses.VehicleResponse{}, err
	}

	return s.mapToResponse(vehicle), nil
}

func (s *vehicleService) Update(ctx context.Context, hashID string, req requests.UpdateVehicleRequest) (responses.VehicleResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.VehicleResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.VehicleResponse{}, errors.New("invalid id format")
	}

	// 1. Ambil data lama
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.VehicleResponse{}, err
	}

	// 2. Patching data (Hanya yang dikirim di request)
	if req.CustomerID != "" {
		customerID, err := s.obfuscator.Decode(req.CustomerID)
		if err != nil {
			return responses.VehicleResponse{}, errors.New("invalid customer id")
		}
		existing.CustomerID = &customerID
	}
	if req.VehicleTypeID != "" {
		vehicleTypeID, err := s.obfuscator.Decode(req.VehicleTypeID)
		if err != nil {
			return responses.VehicleResponse{}, errors.New("invalid vehicle type id")
		}
		existing.VehicleTypeID = vehicleTypeID
	}
	if req.PlateNumber != "" {
		existing.PlateNumber = req.PlateNumber
	}
	if req.Description != "" {
		existing.Description = req.Description
	}

	// 3. Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.VehicleResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *vehicleService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// --- PRIVATE HELPER ---

func (s *vehicleService) mapToResponse(m models.Vehicle) responses.VehicleResponse {
	hID, _ := s.obfuscator.Encode(m.ID)
	hVehicleTypeID, _ := s.obfuscator.Encode(m.VehicleTypeID)
	hCustomerID := ""
	if m.CustomerID != nil {
		hCustomerID, _ = s.obfuscator.Encode(*m.CustomerID)
	}

	lastSeenAt := ""
	if m.LastSeenAt != nil {
		lastSeenAt = m.LastSeenAt.Format("2006-01-02 15:04:05")
	}

	return responses.VehicleResponse{
		ID:            hID,
		CustomerID:    hCustomerID,
		VehicleTypeID: hVehicleTypeID,
		PlateNumber:   m.PlateNumber,
		Description:   m.Description,
		LastSeenAt:    lastSeenAt,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
