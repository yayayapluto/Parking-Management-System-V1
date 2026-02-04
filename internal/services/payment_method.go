package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/pkg/helpers"
)

type PaymentMethodService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.PaymentMethodResponse, error)
	Create(ctx context.Context, req requests.CreatePaymentMethodRequest) (responses.PaymentMethodResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdatePaymentMethodRequest) (responses.PaymentMethodResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type paymentMethodService struct {
	repo       repos.PaymentMethodRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewPaymentMethodService(r repos.PaymentMethodRepository, o helpers.IDObfuscator, v *validator.Validate) PaymentMethodService {
	return &paymentMethodService{
		repo:       r,
		obfuscator: o,
		validator:  v,
	}
}

func (s *paymentMethodService) GetAll(ctx context.Context, p requests.PaginationRequest) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, p, "name")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.PaymentMethodResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Payment methods retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(p, total),
	}, nil
}

func (s *paymentMethodService) GetByID(ctx context.Context, hashedID string) (responses.PaymentMethodResponse, error) {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.PaymentMethodResponse{}, errors.New("invalid payment method id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *paymentMethodService) Create(ctx context.Context, req requests.CreatePaymentMethodRequest) (responses.PaymentMethodResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	pm := &models.PaymentMethod{
		Code:     req.Code,
		Name:     req.Name,
		Config:   pq.StringArray(req.Config),
		IsActive: req.IsActive,
	}

	if err := s.repo.Create(ctx, pm); err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	return s.mapToResponse(*pm), nil
}

func (s *paymentMethodService) Update(ctx context.Context, hashedID string, req requests.UpdatePaymentMethodRequest) (responses.PaymentMethodResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return responses.PaymentMethodResponse{}, errors.New("invalid payment method id")
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	if req.Name != "" {
		existing.Name = req.Name
	}

	// Untuk slice/array, biasanya kita ganti seluruhnya jika diinput
	if req.Config != nil {
		existing.Config = pq.StringArray(req.Config)
	}

	existing.IsActive = req.IsActive

	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *paymentMethodService) Delete(ctx context.Context, hashedID string) error {
	id, err := s.obfuscator.Decode(hashedID)
	if err != nil {
		return errors.New("invalid payment method id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *paymentMethodService) mapToResponse(m models.PaymentMethod) responses.PaymentMethodResponse {
	hId, _ := s.obfuscator.Encode(m.ID)
	return responses.PaymentMethodResponse{
		ID:        hId,
		Code:      m.Code,
		Name:      m.Name,
		Config:    []string(m.Config),
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
