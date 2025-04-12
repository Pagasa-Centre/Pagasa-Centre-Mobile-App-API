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

func testCellGroups(t *testing.T) {
	t.Parallel()

	query := CellGroups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCellGroupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
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

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCellGroupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CellGroups().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCellGroupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CellGroupSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCellGroupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CellGroupExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CellGroup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CellGroupExists to return true, but got false.")
	}
}

func testCellGroupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	cellGroupFound, err := FindCellGroup(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if cellGroupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCellGroupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CellGroups().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCellGroupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CellGroups().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCellGroupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cellGroupOne := &CellGroup{}
	cellGroupTwo := &CellGroup{}
	if err = randomize.Struct(seed, cellGroupOne, cellGroupDBTypes, false, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}
	if err = randomize.Struct(seed, cellGroupTwo, cellGroupDBTypes, false, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cellGroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cellGroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CellGroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCellGroupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cellGroupOne := &CellGroup{}
	cellGroupTwo := &CellGroup{}
	if err = randomize.Struct(seed, cellGroupOne, cellGroupDBTypes, false, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}
	if err = randomize.Struct(seed, cellGroupTwo, cellGroupDBTypes, false, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cellGroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cellGroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func cellGroupBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func cellGroupAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CellGroup) error {
	*o = CellGroup{}
	return nil
}

func testCellGroupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CellGroup{}
	o := &CellGroup{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cellGroupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CellGroup object: %s", err)
	}

	AddCellGroupHook(boil.BeforeInsertHook, cellGroupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cellGroupBeforeInsertHooks = []CellGroupHook{}

	AddCellGroupHook(boil.AfterInsertHook, cellGroupAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cellGroupAfterInsertHooks = []CellGroupHook{}

	AddCellGroupHook(boil.AfterSelectHook, cellGroupAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cellGroupAfterSelectHooks = []CellGroupHook{}

	AddCellGroupHook(boil.BeforeUpdateHook, cellGroupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cellGroupBeforeUpdateHooks = []CellGroupHook{}

	AddCellGroupHook(boil.AfterUpdateHook, cellGroupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cellGroupAfterUpdateHooks = []CellGroupHook{}

	AddCellGroupHook(boil.BeforeDeleteHook, cellGroupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cellGroupBeforeDeleteHooks = []CellGroupHook{}

	AddCellGroupHook(boil.AfterDeleteHook, cellGroupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cellGroupAfterDeleteHooks = []CellGroupHook{}

	AddCellGroupHook(boil.BeforeUpsertHook, cellGroupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cellGroupBeforeUpsertHooks = []CellGroupHook{}

	AddCellGroupHook(boil.AfterUpsertHook, cellGroupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cellGroupAfterUpsertHooks = []CellGroupHook{}
}

func testCellGroupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCellGroupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(cellGroupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCellGroupToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CellGroup
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"cell_group_members\" (\"cell_group_id\", \"user_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"cell_group_members\" (\"cell_group_id\", \"user_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CellGroupSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*CellGroup)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCellGroupToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CellGroup
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cellGroupDBTypes, false, strmangle.SetComplement(cellGroupPrimaryKeyColumns, cellGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.CellGroups[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.CellGroups[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCellGroupToManySetOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CellGroup
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cellGroupDBTypes, false, strmangle.SetComplement(cellGroupPrimaryKeyColumns, cellGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUsers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUsers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.CellGroups) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.CellGroups) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.CellGroups[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.CellGroups[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Users[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Users[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCellGroupToManyRemoveOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CellGroup
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cellGroupDBTypes, false, strmangle.SetComplement(cellGroupPrimaryKeyColumns, cellGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUsers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUsers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.CellGroups) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.CellGroups) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.CellGroups[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.CellGroups[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Users) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Users[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Users[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCellGroupToOneUserUsingLeader(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CellGroup
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.LeaderID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Leader().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := CellGroupSlice{&local}
	if err = local.L.LoadLeader(ctx, tx, false, (*[]*CellGroup)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Leader == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Leader = nil
	if err = local.L.LoadLeader(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Leader == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testCellGroupToOneSetOpUserUsingLeader(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CellGroup
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cellGroupDBTypes, false, strmangle.SetComplement(cellGroupPrimaryKeyColumns, cellGroupColumnsWithoutDefault)...); err != nil {
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
		err = a.SetLeader(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Leader != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.LeaderCellGroups[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.LeaderID, x.ID) {
			t.Error("foreign key was wrong value", a.LeaderID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LeaderID))
		reflect.Indirect(reflect.ValueOf(&a.LeaderID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.LeaderID, x.ID) {
			t.Error("foreign key was wrong value", a.LeaderID, x.ID)
		}
	}
}

func testCellGroupToOneRemoveOpUserUsingLeader(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CellGroup
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cellGroupDBTypes, false, strmangle.SetComplement(cellGroupPrimaryKeyColumns, cellGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetLeader(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveLeader(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Leader().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Leader != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.LeaderID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.LeaderCellGroups) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCellGroupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
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

func testCellGroupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CellGroupSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCellGroupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CellGroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cellGroupDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`, `LeaderID`: `uuid`}
	_                = bytes.MinRead
)

func testCellGroupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(cellGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(cellGroupAllColumns) == len(cellGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCellGroupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cellGroupAllColumns) == len(cellGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CellGroup{}
	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cellGroupDBTypes, true, cellGroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cellGroupAllColumns, cellGroupPrimaryKeyColumns) {
		fields = cellGroupAllColumns
	} else {
		fields = strmangle.SetComplement(
			cellGroupAllColumns,
			cellGroupPrimaryKeyColumns,
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

	slice := CellGroupSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCellGroupsUpsert(t *testing.T) {
	t.Parallel()

	if len(cellGroupAllColumns) == len(cellGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CellGroup{}
	if err = randomize.Struct(seed, &o, cellGroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CellGroup: %s", err)
	}

	count, err := CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, cellGroupDBTypes, false, cellGroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CellGroup struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CellGroup: %s", err)
	}

	count, err = CellGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
