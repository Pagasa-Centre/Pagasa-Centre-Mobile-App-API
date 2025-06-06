package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type RolesRepository interface {
	AssignRole(ctx context.Context, userID, role string) error
	GetUserRoles(ctx context.Context, userID string) ([]string, error)
	AssignRoleTx(ctx context.Context, tx *sqlx.Tx, userID, roleName string, ministryID *string) error
}

type repository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) RolesRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) AssignRoleTx(ctx context.Context, tx *sqlx.Tx, userID, roleName string, ministryID *string) error {
	var minID string
	if ministryID != nil {
		minID = *ministryID
	}

	roleID, err := r.getRoleIDTx(ctx, tx, roleName)
	if err != nil {
		return err
	}

	if strings.Contains(roleName, "Ministry Leader") {
		ministryLeader := entity.MinistryLeader{
			UserID:     userID,
			MinistryID: minID,
		}

		err = ministryLeader.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	return r.assignRoleTx(ctx, tx, userID, roleID)
}

func (r *repository) getRoleIDTx(ctx context.Context, tx *sqlx.Tx, roleName string) (string, error) {
	role, err := entity.Roles(entity.RoleWhere.RoleName.EQ(roleName)).One(ctx, tx)
	if err != nil {
		return "", err
	}

	return role.ID, nil
}

func (r *repository) assignRoleTx(ctx context.Context, tx *sqlx.Tx, userID, roleID string) error {
	userRole := &entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	return userRole.Insert(ctx, tx, boil.Infer())
}

// AssignRole assigns the provided role to the given user.
func (r *repository) AssignRole(ctx context.Context, userID, role string) error {
	roleID, err := r.getRoleID(ctx, role)
	if err != nil {
		return err
	}

	return r.assignRole(ctx, userID, roleID)
}

// getRoleID retrieves the role ID for a given role name using the ORM.
func (r *repository) getRoleID(ctx context.Context, roleName string) (string, error) {
	role, err := entity.Roles(entity.RoleWhere.RoleName.EQ(roleName)).One(ctx, r.db)
	if err != nil {
		return "", fmt.Errorf("failed to get role id for '%s': %w", roleName, err)
	}

	return role.ID, nil
}

// assignRole inserts a record into the user_roles join table using the ORM.
func (r *repository) assignRole(ctx context.Context, userID, roleID string) error {
	userRole := &entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	if err := userRole.Insert(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to assign role %s to user %s: %w", roleID, userID, err)
	}

	return nil
}

// GetUserRoles retrieves all role names assigned to the given user ID.
func (r *repository) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	userRoles, err := entity.UserRoles(
		entity.UserRoleWhere.UserID.EQ(userID),
	).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles for user %s: %w", userID, err)
	}

	var roles []string

	for _, ur := range userRoles {
		role, err := entity.Roles(
			entity.RoleWhere.ID.EQ(ur.RoleID),
		).One(ctx, r.db)
		if err != nil {
			return nil, fmt.Errorf("failed to find role %s: %w", ur.RoleID, err)
		}

		roles = append(roles, role.RoleName)
	}

	return roles, nil
}
