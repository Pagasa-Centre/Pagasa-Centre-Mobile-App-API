// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserRoles(t *testing.T) {
	t.Parallel()

	query := UserRoles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserRoles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserRoleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserRoleExists(ctx, tx, o.UserID, o.RoleID)
	if err != nil {
		t.Errorf("Unable to check if UserRole exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserRoleExists to return true, but got false.")
	}
}

func testUserRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userRoleFound, err := FindUserRole(ctx, tx, o.UserID, o.RoleID)
	if err != nil {
		t.Error(err)
	}

	if userRoleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserRoles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserRoles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userRoleOne := &UserRole{}
	userRoleTwo := &UserRole{}
	if err = randomize.Struct(seed, userRoleOne, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}
	if err = randomize.Struct(seed, userRoleTwo, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userRoleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userRoleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserRoles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userRoleOne := &UserRole{}
	userRoleTwo := &UserRole{}
	if err = randomize.Struct(seed, userRoleOne, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}
	if err = randomize.Struct(seed, userRoleTwo, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userRoleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userRoleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userRoleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func userRoleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserRole) error {
	*o = UserRole{}
	return nil
}

func testUserRolesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserRole{}
	o := &UserRole{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userRoleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserRole object: %s", err)
	}

	AddUserRoleHook(boil.BeforeInsertHook, userRoleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userRoleBeforeInsertHooks = []UserRoleHook{}

	AddUserRoleHook(boil.AfterInsertHook, userRoleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userRoleAfterInsertHooks = []UserRoleHook{}

	AddUserRoleHook(boil.AfterSelectHook, userRoleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userRoleAfterSelectHooks = []UserRoleHook{}

	AddUserRoleHook(boil.BeforeUpdateHook, userRoleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userRoleBeforeUpdateHooks = []UserRoleHook{}

	AddUserRoleHook(boil.AfterUpdateHook, userRoleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userRoleAfterUpdateHooks = []UserRoleHook{}

	AddUserRoleHook(boil.BeforeDeleteHook, userRoleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userRoleBeforeDeleteHooks = []UserRoleHook{}

	AddUserRoleHook(boil.AfterDeleteHook, userRoleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userRoleAfterDeleteHooks = []UserRoleHook{}

	AddUserRoleHook(boil.BeforeUpsertHook, userRoleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userRoleBeforeUpsertHooks = []UserRoleHook{}

	AddUserRoleHook(boil.AfterUpsertHook, userRoleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userRoleAfterUpsertHooks = []UserRoleHook{}
}

func testUserRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userRoleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserRoleToOneRoleUsingRole(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserRole
	var foreign Role

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RoleID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Role().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddRoleHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Role) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserRoleSlice{&local}
	if err = local.L.LoadRole(ctx, tx, false, (*[]*UserRole)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Role = nil
	if err = local.L.LoadRole(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserRoleToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserRole
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserRoleSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserRole)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserRoleToOneSetOpRoleUsingRole(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserRole
	var b, c Role

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userRoleDBTypes, false, strmangle.SetComplement(userRolePrimaryKeyColumns, userRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Role{&b, &c} {
		err = a.SetRole(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Role != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserRoles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoleID != x.ID {
			t.Error("foreign key was wrong value", a.RoleID)
		}

		if exists, err := UserRoleExists(ctx, tx, a.UserID, a.RoleID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testUserRoleToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserRole
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userRoleDBTypes, false, strmangle.SetComplement(userRolePrimaryKeyColumns, userRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserRoles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := UserRoleExists(ctx, tx, a.UserID, a.RoleID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testUserRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserRoleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserRoles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userRoleDBTypes = map[string]string{`UserID`: `uuid`, `RoleID`: `uuid`, `AssignedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testUserRolesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userRolePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userRoleAllColumns) == len(userRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userRoleAllColumns) == len(userRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserRole{}
	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userRoleDBTypes, true, userRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userRoleAllColumns, userRolePrimaryKeyColumns) {
		fields = userRoleAllColumns
	} else {
		fields = strmangle.SetComplement(
			userRoleAllColumns,
			userRolePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UserRoleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(userRoleAllColumns) == len(userRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserRole{}
	if err = randomize.Struct(seed, &o, userRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserRole: %s", err)
	}

	count, err := UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userRoleDBTypes, false, userRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserRole struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserRole: %s", err)
	}

	count, err = UserRoles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
