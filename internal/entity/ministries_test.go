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

func testMinistries(t *testing.T) {
	t.Parallel()

	query := Ministries()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMinistriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
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

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMinistriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Ministries().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMinistriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MinistrySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMinistriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MinistryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Ministry exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MinistryExists to return true, but got false.")
	}
}

func testMinistriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	ministryFound, err := FindMinistry(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if ministryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMinistriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Ministries().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMinistriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Ministries().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMinistriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ministryOne := &Ministry{}
	ministryTwo := &Ministry{}
	if err = randomize.Struct(seed, ministryOne, ministryDBTypes, false, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}
	if err = randomize.Struct(seed, ministryTwo, ministryDBTypes, false, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ministryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ministryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Ministries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMinistriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ministryOne := &Ministry{}
	ministryTwo := &Ministry{}
	if err = randomize.Struct(seed, ministryOne, ministryDBTypes, false, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}
	if err = randomize.Struct(seed, ministryTwo, ministryDBTypes, false, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ministryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ministryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func ministryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func ministryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Ministry) error {
	*o = Ministry{}
	return nil
}

func testMinistriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Ministry{}
	o := &Ministry{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ministryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Ministry object: %s", err)
	}

	AddMinistryHook(boil.BeforeInsertHook, ministryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ministryBeforeInsertHooks = []MinistryHook{}

	AddMinistryHook(boil.AfterInsertHook, ministryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ministryAfterInsertHooks = []MinistryHook{}

	AddMinistryHook(boil.AfterSelectHook, ministryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ministryAfterSelectHooks = []MinistryHook{}

	AddMinistryHook(boil.BeforeUpdateHook, ministryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ministryBeforeUpdateHooks = []MinistryHook{}

	AddMinistryHook(boil.AfterUpdateHook, ministryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ministryAfterUpdateHooks = []MinistryHook{}

	AddMinistryHook(boil.BeforeDeleteHook, ministryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ministryBeforeDeleteHooks = []MinistryHook{}

	AddMinistryHook(boil.AfterDeleteHook, ministryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ministryAfterDeleteHooks = []MinistryHook{}

	AddMinistryHook(boil.BeforeUpsertHook, ministryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ministryBeforeUpsertHooks = []MinistryHook{}

	AddMinistryHook(boil.AfterUpsertHook, ministryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ministryAfterUpsertHooks = []MinistryHook{}
}

func testMinistriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMinistriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(ministryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMinistryToManyMinistryLeaders(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Ministry
	var b, c MinistryLeader

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, ministryLeaderDBTypes, false, ministryLeaderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, ministryLeaderDBTypes, false, ministryLeaderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.MinistryID = a.ID
	c.MinistryID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.MinistryLeaders().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.MinistryID == b.MinistryID {
			bFound = true
		}
		if v.MinistryID == c.MinistryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MinistrySlice{&a}
	if err = a.L.LoadMinistryLeaders(ctx, tx, false, (*[]*Ministry)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MinistryLeaders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MinistryLeaders = nil
	if err = a.L.LoadMinistryLeaders(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MinistryLeaders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMinistryToManyAddOpMinistryLeaders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Ministry
	var b, c, d, e MinistryLeader

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ministryDBTypes, false, strmangle.SetComplement(ministryPrimaryKeyColumns, ministryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MinistryLeader{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, ministryLeaderDBTypes, false, strmangle.SetComplement(ministryLeaderPrimaryKeyColumns, ministryLeaderColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*MinistryLeader{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMinistryLeaders(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.MinistryID {
			t.Error("foreign key was wrong value", a.ID, first.MinistryID)
		}
		if a.ID != second.MinistryID {
			t.Error("foreign key was wrong value", a.ID, second.MinistryID)
		}

		if first.R.Ministry != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Ministry != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MinistryLeaders[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MinistryLeaders[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MinistryLeaders().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testMinistryToOneUserUsingLeader(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Ministry
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
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

	slice := MinistrySlice{&local}
	if err = local.L.LoadLeader(ctx, tx, false, (*[]*Ministry)(&slice), nil); err != nil {
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

func testMinistryToOneOutreachUsingOutreach(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Ministry
	var foreign Outreach

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, ministryDBTypes, false, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, outreachDBTypes, false, outreachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Outreach struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OutreachID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Outreach().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddOutreachHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Outreach) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := MinistrySlice{&local}
	if err = local.L.LoadOutreach(ctx, tx, false, (*[]*Ministry)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Outreach == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Outreach = nil
	if err = local.L.LoadOutreach(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Outreach == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testMinistryToOneSetOpUserUsingLeader(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Ministry
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ministryDBTypes, false, strmangle.SetComplement(ministryPrimaryKeyColumns, ministryColumnsWithoutDefault)...); err != nil {
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

		if x.R.LeaderMinistries[0] != &a {
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

func testMinistryToOneRemoveOpUserUsingLeader(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Ministry
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ministryDBTypes, false, strmangle.SetComplement(ministryPrimaryKeyColumns, ministryColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.LeaderMinistries) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testMinistryToOneSetOpOutreachUsingOutreach(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Ministry
	var b, c Outreach

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ministryDBTypes, false, strmangle.SetComplement(ministryPrimaryKeyColumns, ministryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, outreachDBTypes, false, strmangle.SetComplement(outreachPrimaryKeyColumns, outreachColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, outreachDBTypes, false, strmangle.SetComplement(outreachPrimaryKeyColumns, outreachColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Outreach{&b, &c} {
		err = a.SetOutreach(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Outreach != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Ministries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OutreachID != x.ID {
			t.Error("foreign key was wrong value", a.OutreachID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OutreachID))
		reflect.Indirect(reflect.ValueOf(&a.OutreachID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OutreachID != x.ID {
			t.Error("foreign key was wrong value", a.OutreachID, x.ID)
		}
	}
}

func testMinistriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
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

func testMinistriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MinistrySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMinistriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Ministries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ministryDBTypes = map[string]string{`ID`: `uuid`, `OutreachID`: `uuid`, `Name`: `character varying`, `Description`: `text`, `LeaderID`: `uuid`, `MeetingDay`: `character varying`, `StartTime`: `timestamp without time zone`, `EndTime`: `timestamp without time zone`, `MeetingLocation`: `character varying`}
	_               = bytes.MinRead
)

func testMinistriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(ministryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(ministryAllColumns) == len(ministryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMinistriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ministryAllColumns) == len(ministryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Ministry{}
	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ministryDBTypes, true, ministryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ministryAllColumns, ministryPrimaryKeyColumns) {
		fields = ministryAllColumns
	} else {
		fields = strmangle.SetComplement(
			ministryAllColumns,
			ministryPrimaryKeyColumns,
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

	slice := MinistrySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMinistriesUpsert(t *testing.T) {
	t.Parallel()

	if len(ministryAllColumns) == len(ministryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Ministry{}
	if err = randomize.Struct(seed, &o, ministryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Ministry: %s", err)
	}

	count, err := Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, ministryDBTypes, false, ministryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ministry struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Ministry: %s", err)
	}

	count, err = Ministries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
