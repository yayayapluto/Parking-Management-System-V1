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
	"time"
)

type RefundService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.RefundResponse, error)
	Create(ctx context.Context, req requests.CreateRefundRequest) (responses.RefundResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateRefundRequest) (responses.RefundResponse, error)
	Approve(ctx context.Context, hashedID string, req requests.ApproveRefundRequest) (responses.RefundResponse, error)
	Reject(ctx context.Context, hashedID string, req requests.RejectRefundRequest) (responses.RefundResponse, error)
	Delete(ctx context.Context, hashedID string) error
}

type refundService struct {
	repo       repos.RefundRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewRefundService(r repos.RefundRepository, o helpers.IDObfuscator, v *validator.Validate) RefundService {
	return &refundService{repo: r, obfuscator: o, validator: v}
}

func (s *refundService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "reason")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.RefundResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Refunds retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

func (s *refundService) GetByID(ctx context.Context, hashID string) (responses.RefundResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid refund id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.RefundResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

func (s *refundService) Create(ctx context.Context, req requests.CreateRefundRequest) (responses.RefundResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.RefundResponse{}, err
	}

	paymentID, err := s.obfuscator.Decode(req.PaymentID)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid payment id")
	}

	transactionID, err := s.obfuscator.Decode(req.TransactionID)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid transaction id")
	}

	refundedByID, err := s.obfuscator.Decode(req.RefundedBy)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid refunded_by id")
	}

	refund := &models.Refund{
		PaymentID:     paymentID,
		TransactionID: transactionID,
		Amount:        req.Amount,
		Reason:        req.Reason,
		RefundedBy:    refundedByID,
		Status:        "requested",
	}

	if err := s.repo.Create(ctx, refund); err != nil {
		return responses.RefundResponse{}, err
	}

	return s.mapToResponse(*refund), nil
}

func (s *refundService) Update(ctx context.Context, hashID string, req requests.UpdateRefundRequest) (responses.RefundResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid refund id")
	}

	refund, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.RefundResponse{}, err
	}

	if req.Amount != nil {
		refund.Amount = *req.Amount
	}
	if req.Reason != "" {
		refund.Reason = req.Reason
	}
	if req.Status != "" {
		refund.Status = req.Status
	}
	if req.ApprovedBy != "" {
		approvedByID, err := s.obfuscator.Decode(req.ApprovedBy)
		if err != nil {
			return responses.RefundResponse{}, errors.New("invalid approved_by id")
		}
		refund.ApprovedBy = &approvedByID
	}

	if err := s.repo.Update(ctx, refund); err != nil {
		return responses.RefundResponse{}, err
	}

	return s.mapToResponse(*refund), nil
}

func (s *refundService) Approve(ctx context.Context, hashID string, req requests.ApproveRefundRequest) (responses.RefundResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid refund id")
	}

	refund, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.RefundResponse{}, err
	}

	if refund.Status != "pending_approval" && refund.Status != "requested" {
		return responses.RefundResponse{}, errors.New("refund cannot be approved in current status")
	}

	approvedByID, err := s.obfuscator.Decode(req.ApprovedBy)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid approved_by id")
	}

	now := time.Now()
	refund.Status = "approved"
	refund.ApprovedBy = &approvedByID
	refund.RefundedAt = &now

	if err := s.repo.Update(ctx, refund); err != nil {
		return responses.RefundResponse{}, err
	}

	return s.mapToResponse(*refund), nil
}

func (s *refundService) Reject(ctx context.Context, hashID string, req requests.RejectRefundRequest) (responses.RefundResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.RefundResponse{}, errors.New("invalid refund id")
	}

	refund, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.RefundResponse{}, err
	}

	if refund.Status != "pending_approval" && refund.Status != "requested" {
		return responses.RefundResponse{}, errors.New("refund cannot be rejected in current status")
	}

	refund.Status = "rejected"
	refund.Reason = req.Reason

	if err := s.repo.Update(ctx, refund); err != nil {
		return responses.RefundResponse{}, err
	}

	return s.mapToResponse(*refund), nil
}

func (s *refundService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid refund id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *refundService) mapToResponse(m models.Refund) responses.RefundResponse {
	refundedAt := ""
	if m.RefundedAt != nil {
		refundedAt = m.RefundedAt.Format("2006-01-02 15:04:05")
	}

	approvedBy := ""
	if m.ApprovedBy != nil {
		approvedByEncoded, _ := s.obfuscator.Encode(*m.ApprovedBy)
		approvedBy = approvedByEncoded
	}

	id, _ := s.obfuscator.Encode(m.ID)
	paymentID, _ := s.obfuscator.Encode(m.PaymentID)
	transactionID, _ := s.obfuscator.Encode(m.TransactionID)
	refundedBy, _ := s.obfuscator.Encode(m.RefundedBy)

	return responses.RefundResponse{
		ID:            id,
		PaymentID:     paymentID,
		TransactionID: transactionID,
		Amount:        m.Amount,
		Reason:        m.Reason,
		RefundedBy:    refundedBy,
		ApprovedBy:    approvedBy,
		Status:        m.Status,
		RefundedAt:    refundedAt,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
