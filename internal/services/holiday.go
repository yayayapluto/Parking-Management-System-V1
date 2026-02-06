package services

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/pkg/helpers"
)

type HolidayService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.HolidayResponse, error)
	Create(ctx context.Context, req requests.CreateHolidayRequest) (responses.HolidayResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateHolidayRequest) (responses.HolidayResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type holidayService struct {
	repo       repos.HolidayRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewHolidayService(r repos.HolidayRepository, o helpers.IDObfuscator, v *validator.Validate) HolidayService {
	return &holidayService{
		repo:       r,
		obfuscator: o,
		validator:  v,
	}
}

func (s *holidayService) GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, p, "date")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.HolidayResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Holidays retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(p, total, url),
	}, nil
}

func (s *holidayService) GetByID(ctx context.Context, hashedID string) (responses.HolidayResponse, error) {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.HolidayResponse{}, errors.New("invalid holiday id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.HolidayResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *holidayService) Create(ctx context.Context, req requests.CreateHolidayRequest) (responses.HolidayResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.HolidayResponse{}, err
	}

	// Parsing string date ke time.Time
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return responses.HolidayResponse{}, errors.New("invalid date format, use YYYY-MM-DD")
	}

	holiday := &models.Holiday{
		Date:         parsedDate,
		Name:         req.Name,
		AffectsRates: req.AffectsRates,
	}

	if err := s.repo.Create(ctx, holiday); err != nil {
		return responses.HolidayResponse{}, err
	}

	return s.mapToResponse(*holiday), nil
}

func (s *holidayService) Update(ctx context.Context, hashedID string, req requests.UpdateHolidayRequest) (responses.HolidayResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.HolidayResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.HolidayResponse{}, errors.New("invalid holiday id")
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.HolidayResponse{}, err
	}

	if req.Name != "" {
		existing.Name = req.Name
	}

	// Untuk float, kita asumsikan 0 atau lebih adalah input valid sesuai validator
	existing.AffectsRates = req.AffectsRates

	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.HolidayResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *holidayService) Delete(ctx context.Context, hashedID string) error {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return errors.New("invalid holiday id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *holidayService) mapToResponse(m models.Holiday) responses.HolidayResponse {
	hId, _ := s.obfuscator.Encode(m.ID)
	return responses.HolidayResponse{
		ID:           hId,
		Date:         m.Date.Format("2006-01-02"),
		Name:         m.Name,
		AffectsRates: m.AffectsRates,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
