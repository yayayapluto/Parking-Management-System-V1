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

type CustomerService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.CustomerResponse, error)
	Create(ctx context.Context, req requests.CreateCustomerRequest) (responses.CustomerResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateCustomerRequest) (responses.CustomerResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type customerService struct {
	repo       repos.CustomerRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewCustomerService(r repos.CustomerRepository, o helpers.IDObfuscator, v *validator.Validate) CustomerService {
	return &customerService{repo: r, obfuscator: o, validator: v}
}

// --- CORE METHODS ---

func (s *customerService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "name", "phone", "rfid_uid")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.CustomerResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Customers retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *customerService) GetByID(ctx context.Context, hashID string) (responses.CustomerResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.CustomerResponse{}, errors.New("invalid customer id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.CustomerResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *customerService) Create(ctx context.Context, req requests.CreateCustomerRequest) (responses.CustomerResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.CustomerResponse{}, err
	}

	var registrationSourceID *uint
	if req.RegistrationSourceID != "" {
		id, err := s.obfuscator.Decode(req.RegistrationSourceID)
		if err != nil {
			return responses.CustomerResponse{}, errors.New("invalid registration source id")
		}
		registrationSourceID = &id
	}

	customer := models.Customer{
		RfidUID:              req.RfidUID,
		Name:                 req.Name,
		Phone:                req.Phone,
		RegistrationSourceID: registrationSourceID,
		IsRegistered:         false,
		TotalVisits:          0,
		TotalSpent:           0,
	}

	if err := s.repo.Create(ctx, &customer); err != nil {
		return responses.CustomerResponse{}, err
	}

	return s.mapToResponse(customer), nil
}

func (s *customerService) Update(ctx context.Context, hashID string, req requests.UpdateCustomerRequest) (responses.CustomerResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.CustomerResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.CustomerResponse{}, errors.New("invalid id format")
	}

	// 1. Ambil data lama
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.CustomerResponse{}, err
	}

	// 2. Patching data (Hanya yang dikirim di request)
	if req.RfidUID != "" {
		existing.RfidUID = req.RfidUID
	}
	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.Phone != "" {
		existing.Phone = req.Phone
	}
	if req.RegistrationSourceID != "" {
		registrationSourceID, err := s.obfuscator.Decode(req.RegistrationSourceID)
		if err != nil {
			return responses.CustomerResponse{}, errors.New("invalid registration source id")
		}
		existing.RegistrationSourceID = &registrationSourceID
	}
	if req.IsRegistered != nil {
		existing.IsRegistered = *req.IsRegistered
	}

	// 3. Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.CustomerResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

func (s *customerService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// --- PRIVATE HELPER ---

func (s *customerService) mapToResponse(m models.Customer) responses.CustomerResponse {
	hID, _ := s.obfuscator.Encode(m.ID)
	hRegistrationSourceID := ""
	if m.RegistrationSourceID != nil {
		hRegistrationSourceID, _ = s.obfuscator.Encode(*m.RegistrationSourceID)
	}

	registeredAt := ""
	if m.RegisteredAt != nil {
		registeredAt = m.RegisteredAt.Format("2006-01-02 15:04:05")
	}

	return responses.CustomerResponse{
		ID:                   hID,
		RfidUID:              m.RfidUID,
		Name:                 m.Name,
		Phone:                m.Phone,
		RegistrationSourceID: hRegistrationSourceID,
		IsRegistered:         m.IsRegistered,
		RegisteredAt:         registeredAt,
		TotalVisits:          m.TotalVisits,
		TotalSpent:           m.TotalSpent,
		CreatedAt:            m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:            m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
