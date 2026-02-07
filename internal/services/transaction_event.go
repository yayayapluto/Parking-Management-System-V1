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

type TransactionEventService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.TransactionEventResponse, error)
	Create(ctx context.Context, req requests.CreateTransactionEventRequest) (responses.TransactionEventResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type transactionEventService struct {
	repo       repos.TransactionEventRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewTransactionEventService(r repos.TransactionEventRepository, o helpers.IDObfuscator, v *validator.Validate) TransactionEventService {
	return &transactionEventService{repo: r, obfuscator: o, validator: v}
}

func (s *transactionEventService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "event_type", "plate_detected")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.TransactionEventResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Transaction events retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *transactionEventService) GetByID(ctx context.Context, hashID string) (responses.TransactionEventResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.TransactionEventResponse{}, errors.New("invalid transaction event id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.TransactionEventResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *transactionEventService) Create(ctx context.Context, req requests.CreateTransactionEventRequest) (responses.TransactionEventResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.TransactionEventResponse{}, err
	}

	transactionID, err := s.obfuscator.Decode(req.TransactionID)
	if err != nil {
		return responses.TransactionEventResponse{}, errors.New("invalid transaction id")
	}

	operatorID, err := s.obfuscator.Decode(req.OperatorID)
	if err != nil {
		return responses.TransactionEventResponse{}, errors.New("invalid operator id")
	}

	event := &models.TransactionEvent{
		TransactionID: transactionID,
		EventType:     req.EventType,
		OperatorID:    operatorID,
		PhotoPath:     req.PhotoPath,
		PlateDetected: req.PlateDetected,
		RfidDetected:  req.RfidDetected,
		Notes:         req.Notes,
	}

	if err := s.repo.Create(ctx, event); err != nil {
		return responses.TransactionEventResponse{}, err
	}

	return s.mapToResponse(*event), nil
}

func (s *transactionEventService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid transaction event id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *transactionEventService) mapToResponse(m models.TransactionEvent) responses.TransactionEventResponse {
	id, _ := s.obfuscator.Encode(m.ID)
	tID, _ := s.obfuscator.Encode(m.TransactionID)
	oID, _ := s.obfuscator.Encode(m.OperatorID)

	return responses.TransactionEventResponse{
		ID:            id,
		TransactionID: tID,
		EventType:     m.EventType,
		OperatorID:    oID,
		PhotoPath:     m.PhotoPath,
		PlateDetected: m.PlateDetected,
		RfidDetected:  m.RfidDetected,
		Notes:         m.Notes,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
