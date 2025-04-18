package approvals

import (
	"context"
	"fmt"
	"slices"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals/dto"
	dto2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/contracts"
	context2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
)

type ApprovalService interface {
	CreateNewApproval(ctx context.Context, Approval *domain.Approval) error
	GetAll(ctx context.Context) ([]dto.Approval, error)
	SetUserService(us contracts.UserService)
	UpdateApprovalStatus(ctx context.Context, approvalID, status string) error
}

type service struct {
	logger       *zap.Logger
	approvalRepo storage.ApprovalRepository
	userService  contracts.UserService
	roleService  roles.RolesService
}

func NewApprovalService(
	logger *zap.Logger,
	approvalRepo storage.ApprovalRepository,
	userService contracts.UserService,
	roleService roles.RolesService,
) ApprovalService {
	return &service{
		logger:       logger,
		approvalRepo: approvalRepo,
		userService:  userService,
		roleService:  roleService,
	}
}

func (s *service) CreateNewApproval(ctx context.Context, approval *domain.Approval) error {
	s.logger.Info("Creating new approval")

	entity := mappers.ToApprovalEntity(approval)

	err := s.approvalRepo.Insert(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context) ([]dto.Approval, error) {
	// Get the user ID from the context
	userID, err := context2.GetUserIDString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from context: %w", err)
	}

	_, err = s.userService.GetUserById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user does not exist or has been deleted: %w", err)
	}

	// Get user roles
	userRoles, err := s.roleService.GetUserRoles(ctx, userID)
	if err != nil {
		return nil, err
	}

	var approvalsSlice entity.ApprovalSlice
	if slices.Contains(userRoles, "Admin") {
		// 1. Get all approvals by requesterID
		approvalsSlice, err = s.approvalRepo.GetAllPendingApprovals(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get approvals: %w", err)
		}
	} else {
		// 1. Get all approvals by requesterID
		approvalsSlice, err = s.approvalRepo.GetAllPendingApprovalsByUserID(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("failed to get approvals: %w", err)
		}
	}

	// 2. For each Approval, get userDetails and add to response
	var approvals []dto.Approval

	for _, apr := range approvalsSlice {
		u, err := s.userService.GetUserById(ctx, apr.RequesterID)
		if err != nil {
			return nil, fmt.Errorf("user does not exist or has been deleted: %w", err)
		}

		var approvalType string
		if apr.Type.Valid {
			approvalType = apr.Type.String
		}

		approval := dto.Approval{
			ID:            apr.ID,
			RequestedRole: apr.RequestedRole,
			Type:          approvalType,
			Status:        apr.Status,
			RequesterDetails: dto2.UserDetails{
				FirstName:   u.FirstName,
				LastName:    u.LastName,
				Email:       u.Email,
				PhoneNumber: u.PhoneNumber,
			},
		}
		approvals = append(approvals, approval)
	}

	return approvals, nil
}

func (s *service) SetUserService(us contracts.UserService) {
	s.userService = us
}

func (s *service) UpdateApprovalStatus(ctx context.Context, approvalID, status string) error {
	s.logger.Info("Updating approval status")

	// 1. Get approval by id
	approval, err := s.approvalRepo.GetApprovalByID(ctx, approvalID)
	if err != nil {
		return err
	}

	switch status {
	case domain.Approved:
		approval.Status = domain.Approved
	case domain.Rejected:
		approval.Status = domain.Rejected
	default:
		return fmt.Errorf("invalid approval status: %s", status)
	}
	// 2. Update Status
	err = s.approvalRepo.Update(ctx, approval)
	if err != nil {
		return err
	}

	if approval.Status != domain.Approved {
		return nil
	}

	// 3. Check type and depending on type assign roles
	err = s.roleService.AssignRole(ctx, approval.RequesterID, approval.RequestedRole)
	if err != nil {
		return err
	}

	// 4. Send a message to applicant/requester?

	return nil
}
