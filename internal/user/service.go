package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals"
	approvalDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/contracts"
	rolesDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/mappers"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
	context2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	RegisterNewUser(ctx context.Context, user *domain.User, req dto.RegisterRequest) (*entity.User, error)
	GenerateToken(user *entity.User) (string, error)
	UpdateUserDetails(ctx context.Context, req dto.UpdateDetailsRequest) (*entity.User, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error)
	AuthenticateAndGenerateToken(ctx context.Context, email, password string) (*AuthResult, error)
	SetMinistryService(ms contracts.MinistryService)
}

type userService struct {
	logger           *zap.Logger
	userRepo         userStorage.UserRepository
	jwtSecret        string
	ministryService  contracts.MinistryService
	approvalsService approvals.ApprovalService
}

func NewUserService(
	logger *zap.Logger,
	userRepo userStorage.UserRepository,
	jwtSecret string,
	ministryService contracts.MinistryService,
	approvalsService approvals.ApprovalService,
) UserService {
	return &userService{
		logger:           logger,
		userRepo:         userRepo,
		jwtSecret:        jwtSecret,
		ministryService:  ministryService,
		approvalsService: approvalsService,
	}
}

type AuthResult struct {
	User  *entity.User
	Token string
}

var (
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrInvalidLoginDetails = errors.New("invalid email or password")
)

func (s *userService) SetMinistryService(ms contracts.MinistryService) {
	s.ministryService = ms
}

func (s *userService) AuthenticateAndGenerateToken(ctx context.Context, email, password string) (*AuthResult, error) {
	user, err := s.AuthenticateUser(ctx, email, password)
	if err != nil {
		return nil, ErrInvalidLoginDetails
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &AuthResult{
		User:  user,
		Token: token,
	}, nil
}

// AuthenticateUser checks credentials and returns the user if valid, otherwise an error.
func (s *userService) AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// RegisterNewUser inserts a new user into the user table and applies any roles that were provided by the user.
func (s *userService) RegisterNewUser(ctx context.Context, user *domain.User, req dto.RegisterRequest) (*entity.User, error) {
	s.logger.Info("Registering new user")

	user.PhoneNumber = NormalizePhoneNumber(user.PhoneNumber)

	userEntity := mappers.ToUserEntity(*user)

	userID, err := s.userRepo.InsertUser(ctx, userEntity)
	if err != nil {
		// Check if the error is a unique constraint violation (email already exists)
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			// PostgreSQL error code 23505 = unique_violation
			return nil, ErrEmailAlreadyExists
		}

		return nil, err
	}

	var uID string
	if userID != nil {
		uID = *userID
	}

	createApproval := func(requestedRole, approvalType string) error {
		approval := &approvalDomain.Approval{
			RequesterID:   uID,
			ApproverID:    nil,
			RequestedRole: requestedRole,
			Type:          approvalType,
			Status:        approvalDomain.Pending,
		}
		return s.approvalsService.CreateNewApproval(ctx, approval)
	}

	if req.IsLeader {
		if err = createApproval(rolesDomain.RoleLeader, approvalDomain.LeaderStatusConfirmation); err != nil {
			return nil, err
		}
	}

	if req.IsPrimary {
		if err = createApproval(rolesDomain.RolePrimary, approvalDomain.PrimaryStatusConfirmation); err != nil {
			return nil, err
		}
	}

	if req.IsPastor {
		if err = createApproval(rolesDomain.RolePastor, approvalDomain.PastorStatusConfirmation); err != nil {
			return nil, err
		}
	}

	if req.IsMinistryLeader {
		if req.MinistryID == nil {
			return nil, fmt.Errorf("ministry_id is required for ministry leader role")
		}

		ministry, err := s.ministryService.GetByID(ctx, *req.MinistryID)
		if err != nil {
			return nil, err
		}

		if err := createApproval(ministry.Name, approvalDomain.MinistryLeaderStatusConfirmation); err != nil {
			return nil, err
		}
	}
	u, err := s.GetUserById(ctx, uID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return u, nil
}

// GetUserByEmail retrieves a user by their email.
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	s.logger.Info("Getting user by email")

	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GenerateToken generates a JWT for the authenticated user.
func (s *userService) GenerateToken(user *entity.User) (string, error) {
	s.logger.Info("Generating token")
	// Define claims; you can add custom claims as needed.
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func (s *userService) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	s.logger.Info("Getting user by id")

	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserDetails updates a user in the database.
func (s *userService) UpdateUserDetails(ctx context.Context, req dto.UpdateDetailsRequest) (*entity.User, error) {
	s.logger.Info("Updating user details")

	// Get the user ID from the context
	userID, err := context2.GetUserIDString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from context: %w", err)
	}

	// Retrieve the current user from the database.
	currentUser, err := s.GetUserById(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	// Update user fields based on the request.
	s.updateUserFields(currentUser, req)

	user, err := s.userRepo.UpdateUser(ctx, currentUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by their ID.
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	s.logger.Info("Deleting user")

	if err := s.userRepo.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("service error deleting user: %w", err)
	}

	return nil
}

// updateUserFields updates the provided user entity with the values from the update request.
func (s *userService) updateUserFields(user *entity.User, req dto.UpdateDetailsRequest) {
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}

	if req.LastName != "" {
		user.LastName = req.LastName
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.PhoneNumber != "" {
		user.Phone = null.StringFrom(NormalizePhoneNumber(req.PhoneNumber))
	}

	if req.Birthday != "" {
		parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			s.logger.Sugar().Errorw("failed to parse birthday", "error", err)
		} else {
			user.Birthday = null.TimeFrom(parsedBirthday)
		}
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			s.logger.Sugar().Errorw("failed to hash new password", "error", err)
		} else {
			user.HashedPassword = string(hashedPassword)
		}
	}

	if req.CellLeaderID != nil {
		user.CellLeaderID = null.StringFrom(*req.CellLeaderID)
	}

	if req.OutreachID != "" {
		user.OutreachID = null.StringFrom(req.OutreachID)
	}
}

func NormalizePhoneNumber(phone string) string {
	if strings.HasPrefix(phone, "07") {
		// Convert UK local (07...) to E.164 (+44...)
		return "+44" + phone[1:]
	}

	return phone
}
