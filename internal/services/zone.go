package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/pkg/helpers"
)

type ZoneService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.ZoneResponse, error)
	Create(ctx context.Context, req requests.CreateZoneRequest) (responses.ZoneResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateZoneRequest) (responses.ZoneResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type zoneService struct {
	repo       repos.ZoneRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewZoneService(r repos.ZoneRepository, o helpers.IDObfuscator, v *validator.Validate) ZoneService {
	return &zoneService{repo: r, obfuscator: o, validator: v}
}

// --- CORE METHODS ---

func (s *zoneService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	// Panggil repo dengan searchable columns: name dan location
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "name", "location")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.ZoneResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Zones retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *zoneService) GetByID(ctx context.Context, hashID string) (responses.ZoneResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ZoneResponse{}, errors.New("invalid zone id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ZoneResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *zoneService) Create(ctx context.Context, req requests.CreateZoneRequest) (responses.ZoneResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ZoneResponse{}, err
	}

	// Decode ZoneTypeID
	zoneTypeID, err := s.obfuscator.Decode(req.ZoneTypeID)
	if err != nil {
		return responses.ZoneResponse{}, errors.New("invalid zone type id")
	}

	var operatingHourStart, operatingHourEnd sql.NullString
	if req.OperatingHourStart != "" {
		operatingHourStart = sql.NullString{String: req.OperatingHourStart, Valid: true}
	}
	if req.OperatingHourEnd != "" {
		operatingHourEnd = sql.NullString{String: req.OperatingHourEnd, Valid: true}
	}

	zone := models.Zone{
		Name:               req.Name,
		ZoneTypeID:         zoneTypeID,
		Location:           req.Location,
		MaximumCapacity:    req.MaximumCapacity,
		DefaultFee:         req.DefaultFee,
		Is24Hours:          req.Is24Hours,
		OperatingHourStart: operatingHourStart,
		OperatingHourEnd:   operatingHourEnd,
		IsActive:           true,
		IsInMaintenance:    false,
		Description:        req.Description,
	}

	if err := s.repo.Create(ctx, &zone); err != nil {
		return responses.ZoneResponse{}, err
	}

	return s.mapToResponse(zone), nil
}

func (s *zoneService) Update(ctx context.Context, hashID string, req requests.UpdateZoneRequest) (responses.ZoneResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ZoneResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ZoneResponse{}, errors.New("invalid id format")
	}

	// 1. Ambil data lama
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ZoneResponse{}, err
	}

	// 2. Patching data (Hanya yang dikirim di request)
	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.ZoneTypeID != "" {
		zoneTypeID, err := s.obfuscator.Decode(req.ZoneTypeID)
		if err != nil {
			return responses.ZoneResponse{}, errors.New("invalid zone type id")
		}
		existing.ZoneTypeID = zoneTypeID
	}
	if req.Location != "" {
		existing.Location = req.Location
	}
	if req.MaximumCapacity != nil {
		existing.MaximumCapacity = *req.MaximumCapacity
	}
	if req.DefaultFee != nil {
		existing.DefaultFee = *req.DefaultFee
	}
	if req.Is24Hours != nil {
		existing.Is24Hours = *req.Is24Hours
	}
	if req.OperatingHourStart != "" {
		existing.OperatingHourStart = sql.NullString{String: req.OperatingHourStart, Valid: true}
	}
	if req.OperatingHourEnd != "" {
		existing.OperatingHourEnd = sql.NullString{String: req.OperatingHourEnd, Valid: true}
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}
	if req.IsInMaintenance != nil {
		existing.IsInMaintenance = *req.IsInMaintenance
	}
	if req.Description != "" {
		existing.Description = req.Description
	}

	// 3. Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.ZoneResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *zoneService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// --- PRIVATE HELPER ---

func (s *zoneService) mapToResponse(m models.Zone) responses.ZoneResponse {
	hID, _ := s.obfuscator.Encode(m.ID)
	hZoneTypeID, _ := s.obfuscator.Encode(m.ZoneTypeID)

	operatingHourStart := ""
	if m.OperatingHourStart.Valid {
		operatingHourStart = m.OperatingHourStart.String
	}

	operatingHourEnd := ""
	if m.OperatingHourEnd.Valid {
		operatingHourEnd = m.OperatingHourEnd.String
	}

	return responses.ZoneResponse{
		ID:                 hID,
		Name:               m.Name,
		ZoneTypeID:         hZoneTypeID,
		Location:           m.Location,
		MaximumCapacity:    m.MaximumCapacity,
		DefaultFee:         m.DefaultFee,
		Is24Hours:          m.Is24Hours,
		OperatingHourStart: operatingHourStart,
		OperatingHourEnd:   operatingHourEnd,
		IsActive:           m.IsActive,
		IsInMaintenance:    m.IsInMaintenance,
		Description:        m.Description,
		CreatedAt:          m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:          m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
