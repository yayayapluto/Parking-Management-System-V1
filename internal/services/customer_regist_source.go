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

type CustomerRegistSourceService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.CustomerRegistSourceResponse, error)
	Create(ctx context.Context, req requests.CreateCustomerRegistSourceRequest) (responses.CustomerRegistSourceResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateCustomerRegistSourceRequest) (responses.CustomerRegistSourceResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type customerRegistSourceService struct {
	repo       repos.CustomerRegistSourceRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewCustomerRegistSourceService(r repos.CustomerRegistSourceRepository, o helpers.IDObfuscator, v *validator.Validate) CustomerRegistSourceService {
	return &customerRegistSourceService{
		repo:       r,
		obfuscator: o,
		validator:  v,
	}
}

func (s *customerRegistSourceService) GetAll(ctx context.Context, p requests.PaginationRequest) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, p, "name")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.CustomerRegistSourceResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Registration sources retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(p, total),
	}, nil
}

func (s *customerRegistSourceService) GetByID(ctx context.Context, hashedID string) (responses.CustomerRegistSourceResponse, error) {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.CustomerRegistSourceResponse{}, errors.New("invalid registration source id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.CustomerRegistSourceResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *customerRegistSourceService) Create(ctx context.Context, req requests.CreateCustomerRegistSourceRequest) (responses.CustomerRegistSourceResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.CustomerRegistSourceResponse{}, err
	}

	source := &models.CustomerRegistSource{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, source); err != nil {
		return responses.CustomerRegistSourceResponse{}, err
	}

	return s.mapToResponse(*source), nil
}

func (s *customerRegistSourceService) Update(ctx context.Context, hashedID string, req requests.UpdateCustomerRegistSourceRequest) (responses.CustomerRegistSourceResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.CustomerRegistSourceResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.CustomerRegistSourceResponse{}, errors.New("invalid registration source id")
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.CustomerRegistSourceResponse{}, err
	}

	if req.Name != "" {
		existing.Name = req.Name
	}

	if req.Description != "" {
		existing.Description = req.Description
	}

	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.CustomerRegistSourceResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *customerRegistSourceService) Delete(ctx context.Context, hashedID string) error {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return errors.New("invalid registration source id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *customerRegistSourceService) mapToResponse(m models.CustomerRegistSource) responses.CustomerRegistSourceResponse {
	hId, _ := s.obfuscator.Encode(m.ID)
	return responses.CustomerRegistSourceResponse{
		ID:          hId,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
