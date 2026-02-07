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
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
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
	return &paymentMethodService{repo: r, obfuscator: o, validator: v}
}

// --- CORE METHODS ---

func (s *paymentMethodService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "code", "name")
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
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *paymentMethodService) GetByID(ctx context.Context, hashID string) (responses.PaymentMethodResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
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

	paymentMethod := models.PaymentMethod{
		Code:     req.Code,
		Name:     req.Name,
		Config:   pq.StringArray(req.Config),
		IsActive: req.IsActive,
	}

	if err := s.repo.Create(ctx, &paymentMethod); err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	return s.mapToResponse(paymentMethod), nil
}

func (s *paymentMethodService) Update(ctx context.Context, hashID string, req requests.UpdatePaymentMethodRequest) (responses.PaymentMethodResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.PaymentMethodResponse{}, errors.New("invalid id format")
	}

	// 1. Ambil data lama
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	// 2. Patching data (Hanya yang dikirim di request)
	if req.Name != "" {
		existing.Name = req.Name
	}
	if len(req.Config) > 0 {
		existing.Config = pq.StringArray(req.Config)
	}
	existing.IsActive = req.IsActive

	// 3. Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.PaymentMethodResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *paymentMethodService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// --- PRIVATE HELPER ---

func (s *paymentMethodService) mapToResponse(m models.PaymentMethod) responses.PaymentMethodResponse {
	hID, _ := s.obfuscator.Encode(m.ID)

	config := []string{}
	if m.Config != nil {
		config = m.Config
	}

	return responses.PaymentMethodResponse{
		ID:        hID,
		Code:      m.Code,
		Name:      m.Name,
		Config:    config,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
