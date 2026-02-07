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
	"time"
)

type ZoneRateService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.ZoneRateResponse, error)
	Create(ctx context.Context, req requests.CreateZoneRateRequest) (responses.ZoneRateResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateZoneRateRequest) (responses.ZoneRateResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type zoneRateService struct {
	repo       repos.ZoneRateRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewZoneRateService(r repos.ZoneRateRepository, o helpers.IDObfuscator, v *validator.Validate) ZoneRateService {
	return &zoneRateService{repo: r, obfuscator: o, validator: v}
}

// --- CORE METHODS ---

func (s *zoneRateService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	// Panggil repo dengan searchable columns
	mdls, total, err := s.repo.GetWithPagination(ctx, req)
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.ZoneRateResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Zone rates retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *zoneRateService) GetByID(ctx context.Context, hashID string) (responses.ZoneRateResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ZoneRateResponse{}, errors.New("invalid zone rate id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ZoneRateResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *zoneRateService) Create(ctx context.Context, req requests.CreateZoneRateRequest) (responses.ZoneRateResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ZoneRateResponse{}, err
	}

	// Decode ZoneID
	zoneID, err := s.obfuscator.Decode(req.ZoneID)
	if err != nil {
		return responses.ZoneRateResponse{}, errors.New("invalid zone id")
	}

	// Decode VehicleTypeID
	vehicleTypeID, err := s.obfuscator.Decode(req.VehicleTypeID)
	if err != nil {
		return responses.ZoneRateResponse{}, errors.New("invalid vehicle type id")
	}

	// Parse dates
	validFrom, err := time.Parse("2006-01-02", req.ValidFrom)
	if err != nil {
		return responses.ZoneRateResponse{}, errors.New("invalid valid_from date format")
	}

	validTo, err := time.Parse("2006-01-02", req.ValidTo)
	if err != nil {
		return responses.ZoneRateResponse{}, errors.New("invalid valid_to date format")
	}

	var effectiveHourFrom, effectiveHourTo sql.NullString
	if req.EffectiveHourFrom != "" {
		effectiveHourFrom = sql.NullString{String: req.EffectiveHourFrom, Valid: true}
	}
	if req.EffectiveHourTo != "" {
		effectiveHourTo = sql.NullString{String: req.EffectiveHourTo, Valid: true}
	}

	var holidayID *uint
	if req.HolidayID != nil && *req.HolidayID != "" {
		hID, err := s.obfuscator.Decode(*req.HolidayID)
		if err != nil {
			return responses.ZoneRateResponse{}, errors.New("invalid holiday id")
		}
		holidayID = &hID
	}

	zoneRate := models.ZoneRate{
		ZoneID:             zoneID,
		VehicleTypeID:      vehicleTypeID,
		HourlyRate:         req.HourlyRate,
		DailyMaxRate:       req.DailyMaxRate,
		FreeMinutes:        req.FreeMinutes,
		IsWeekend:          req.IsWeekend,
		IsHoliday:          req.IsHoliday,
		HolidayID:          holidayID,
		ValidFrom:          validFrom,
		ValidTo:            validTo,
		EffectiveHourFrom:  effectiveHourFrom,
		EffectiveHourTo:    effectiveHourTo,
		IsActive:           true,
	}

	if err := s.repo.Create(ctx, &zoneRate); err != nil {
		return responses.ZoneRateResponse{}, err
	}

	return s.mapToResponse(zoneRate), nil
}

func (s *zoneRateService) Update(ctx context.Context, hashID string, req requests.UpdateZoneRateRequest) (responses.ZoneRateResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ZoneRateResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ZoneRateResponse{}, errors.New("invalid id format")
	}

	// 1. Ambil data lama
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ZoneRateResponse{}, err
	}

	// 2. Patching data (Hanya yang dikirim di request)
	if req.ZoneID != "" {
		zoneID, err := s.obfuscator.Decode(req.ZoneID)
		if err != nil {
			return responses.ZoneRateResponse{}, errors.New("invalid zone id")
		}
		existing.ZoneID = zoneID
	}

	if req.VehicleTypeID != "" {
		vehicleTypeID, err := s.obfuscator.Decode(req.VehicleTypeID)
		if err != nil {
			return responses.ZoneRateResponse{}, errors.New("invalid vehicle type id")
		}
		existing.VehicleTypeID = vehicleTypeID
	}

	if req.HourlyRate != nil {
		existing.HourlyRate = *req.HourlyRate
	}

	if req.DailyMaxRate != nil {
		existing.DailyMaxRate = *req.DailyMaxRate
	}

	if req.FreeMinutes != nil {
		existing.FreeMinutes = *req.FreeMinutes
	}

	if req.IsWeekend != nil {
		existing.IsWeekend = *req.IsWeekend
	}

	if req.IsHoliday != nil {
		existing.IsHoliday = *req.IsHoliday
	}

	if req.HolidayID != nil {
		if *req.HolidayID == "" {
			existing.HolidayID = nil
		} else {
			hID, err := s.obfuscator.Decode(*req.HolidayID)
			if err != nil {
				return responses.ZoneRateResponse{}, errors.New("invalid holiday id")
			}
			existing.HolidayID = &hID
		}
	}

	if req.ValidFrom != "" {
		validFrom, err := time.Parse("2006-01-02", req.ValidFrom)
		if err != nil {
			return responses.ZoneRateResponse{}, errors.New("invalid valid_from date format")
		}
		existing.ValidFrom = validFrom
	}

	if req.ValidTo != "" {
		validTo, err := time.Parse("2006-01-02", req.ValidTo)
		if err != nil {
			return responses.ZoneRateResponse{}, errors.New("invalid valid_to date format")
		}
		existing.ValidTo = validTo
	}

	if req.EffectiveHourFrom != "" {
		existing.EffectiveHourFrom = sql.NullString{String: req.EffectiveHourFrom, Valid: true}
	}

	if req.EffectiveHourTo != "" {
		existing.EffectiveHourTo = sql.NullString{String: req.EffectiveHourTo, Valid: true}
	}

	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	// 3. Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.ZoneRateResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *zoneRateService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// --- PRIVATE HELPER ---

func (s *zoneRateService) mapToResponse(m models.ZoneRate) responses.ZoneRateResponse {
	hID, _ := s.obfuscator.Encode(m.ID)
	hZoneID, _ := s.obfuscator.Encode(m.ZoneID)
	hVehicleTypeID, _ := s.obfuscator.Encode(m.VehicleTypeID)

	var hHolidayID *string
	if m.HolidayID != nil {
		hID, _ := s.obfuscator.Encode(*m.HolidayID)
		hHolidayID = &hID
	}

	effectiveHourFrom := ""
	if m.EffectiveHourFrom.Valid {
		effectiveHourFrom = m.EffectiveHourFrom.String
	}

	effectiveHourTo := ""
	if m.EffectiveHourTo.Valid {
		effectiveHourTo = m.EffectiveHourTo.String
	}

	return responses.ZoneRateResponse{
		ID:                hID,
		ZoneID:            hZoneID,
		VehicleTypeID:     hVehicleTypeID,
		HourlyRate:        m.HourlyRate,
		DailyMaxRate:      m.DailyMaxRate,
		FreeMinutes:       m.FreeMinutes,
		IsWeekend:         m.IsWeekend,
		IsHoliday:         m.IsHoliday,
		HolidayID:         hHolidayID,
		ValidFrom:         m.ValidFrom.Format("2006-01-02"),
		ValidTo:           m.ValidTo.Format("2006-01-02"),
		EffectiveHourFrom: effectiveHourFrom,
		EffectiveHourTo:   effectiveHourTo,
		IsActive:          m.IsActive,
		CreatedAt:         m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:         m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
