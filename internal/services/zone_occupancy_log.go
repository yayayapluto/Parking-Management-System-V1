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

type ZoneOccupancyLogService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.ZoneOccupancyLogResponse, error)
	Create(ctx context.Context, req requests.CreateZoneOccupancyLogRequest) (responses.ZoneOccupancyLogResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type zoneOccupancyLogService struct {
	repo       repos.ZoneOccupancyLogRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewZoneOccupancyLogService(r repos.ZoneOccupancyLogRepository, o helpers.IDObfuscator, v *validator.Validate) ZoneOccupancyLogService {
	return &zoneOccupancyLogService{repo: r, obfuscator: o, validator: v}
}

func (s *zoneOccupancyLogService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	// Sesuaikan column search dengan field yang ada di model baru: "event_type"
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "event_type")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.ZoneOccupancyLogResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Zone occupancy logs retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *zoneOccupancyLogService) GetByID(ctx context.Context, hashID string) (responses.ZoneOccupancyLogResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.ZoneOccupancyLogResponse{}, errors.New("invalid zone occupancy log id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.ZoneOccupancyLogResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *zoneOccupancyLogService) Create(ctx context.Context, req requests.CreateZoneOccupancyLogRequest) (responses.ZoneOccupancyLogResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.ZoneOccupancyLogResponse{}, err
	}

	// Decode all Hashed IDs from request
	zoneID, err := s.obfuscator.Decode(req.ZoneID)
	if err != nil {
		return responses.ZoneOccupancyLogResponse{}, errors.New("invalid zone id")
	}

	transactionID, err := s.obfuscator.Decode(req.TransactionID)
	if err != nil {
		return responses.ZoneOccupancyLogResponse{}, errors.New("invalid transaction id")
	}

	operatorID, err := s.obfuscator.Decode(req.OperatorID)
	if err != nil {
		return responses.ZoneOccupancyLogResponse{}, errors.New("invalid operator id")
	}

	log := &models.ZoneOccupancyLog{
		ZoneID:         zoneID,
		TransactionID:  transactionID,
		OperatorID:     operatorID,
		OccupiedCount:  req.OccupiedCount,
		AvailableSlots: req.AvailableSlots,
		EventType:      req.EventType,
		Notes:          req.Notes,
	}

	if err := s.repo.Create(ctx, log); err != nil {
		return responses.ZoneOccupancyLogResponse{}, err
	}

	return s.mapToResponse(*log), nil
}

func (s *zoneOccupancyLogService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid zone occupancy log id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *zoneOccupancyLogService) mapToResponse(m models.ZoneOccupancyLog) responses.ZoneOccupancyLogResponse {
	// Handle multiple return values (string, error) from Encode
	id, _ := s.obfuscator.Encode(m.ID)
	zID, _ := s.obfuscator.Encode(m.ZoneID)
	tID, _ := s.obfuscator.Encode(m.TransactionID)
	oID, _ := s.obfuscator.Encode(m.OperatorID)

	return responses.ZoneOccupancyLogResponse{
		ID:             id,
		ZoneID:         zID,
		TransactionID:  tID,
		OperatorID:     oID,
		OccupiedCount:  m.OccupiedCount,
		AvailableSlots: m.AvailableSlots,
		EventType:      m.EventType,
		Notes:          m.Notes,
		CreatedAt:      m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
