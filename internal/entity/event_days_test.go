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

func testEventDays(t *testing.T) {
	t.Parallel()

	query := EventDays()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEventDaysDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
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

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventDaysQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EventDays().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventDaysSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventDaySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventDaysExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EventDayExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EventDay exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EventDayExists to return true, but got false.")
	}
}

func testEventDaysFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	eventDayFound, err := FindEventDay(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if eventDayFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEventDaysBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EventDays().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEventDaysOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EventDays().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEventDaysAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	eventDayOne := &EventDay{}
	eventDayTwo := &EventDay{}
	if err = randomize.Struct(seed, eventDayOne, eventDayDBTypes, false, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}
	if err = randomize.Struct(seed, eventDayTwo, eventDayDBTypes, false, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventDayOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventDayTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EventDays().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEventDaysCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	eventDayOne := &EventDay{}
	eventDayTwo := &EventDay{}
	if err = randomize.Struct(seed, eventDayOne, eventDayDBTypes, false, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}
	if err = randomize.Struct(seed, eventDayTwo, eventDayDBTypes, false, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventDayOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventDayTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func eventDayBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func eventDayAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EventDay) error {
	*o = EventDay{}
	return nil
}

func testEventDaysHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EventDay{}
	o := &EventDay{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, eventDayDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EventDay object: %s", err)
	}

	AddEventDayHook(boil.BeforeInsertHook, eventDayBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	eventDayBeforeInsertHooks = []EventDayHook{}

	AddEventDayHook(boil.AfterInsertHook, eventDayAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	eventDayAfterInsertHooks = []EventDayHook{}

	AddEventDayHook(boil.AfterSelectHook, eventDayAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	eventDayAfterSelectHooks = []EventDayHook{}

	AddEventDayHook(boil.BeforeUpdateHook, eventDayBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	eventDayBeforeUpdateHooks = []EventDayHook{}

	AddEventDayHook(boil.AfterUpdateHook, eventDayAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	eventDayAfterUpdateHooks = []EventDayHook{}

	AddEventDayHook(boil.BeforeDeleteHook, eventDayBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	eventDayBeforeDeleteHooks = []EventDayHook{}

	AddEventDayHook(boil.AfterDeleteHook, eventDayAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	eventDayAfterDeleteHooks = []EventDayHook{}

	AddEventDayHook(boil.BeforeUpsertHook, eventDayBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	eventDayBeforeUpsertHooks = []EventDayHook{}

	AddEventDayHook(boil.AfterUpsertHook, eventDayAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	eventDayAfterUpsertHooks = []EventDayHook{}
}

func testEventDaysInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventDaysInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(eventDayColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventDayToOneEventUsingEvent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local EventDay
	var foreign Event

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, eventDayDBTypes, false, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.EventID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Event().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddEventHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Event) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := EventDaySlice{&local}
	if err = local.L.LoadEvent(ctx, tx, false, (*[]*EventDay)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Event == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Event = nil
	if err = local.L.LoadEvent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Event == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testEventDayToOneSetOpEventUsingEvent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EventDay
	var b, c Event

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, eventDayDBTypes, false, strmangle.SetComplement(eventDayPrimaryKeyColumns, eventDayColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Event{&b, &c} {
		err = a.SetEvent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Event != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EventDays[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EventID != x.ID {
			t.Error("foreign key was wrong value", a.EventID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EventID))
		reflect.Indirect(reflect.ValueOf(&a.EventID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EventID != x.ID {
			t.Error("foreign key was wrong value", a.EventID, x.ID)
		}
	}
}

func testEventDaysReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
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

func testEventDaysReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventDaySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEventDaysSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EventDays().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	eventDayDBTypes = map[string]string{`ID`: `uuid`, `EventID`: `uuid`, `Date`: `date`, `StartTime`: `time without time zone`, `EndTime`: `time without time zone`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testEventDaysUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(eventDayPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(eventDayAllColumns) == len(eventDayPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEventDaysSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(eventDayAllColumns) == len(eventDayPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EventDay{}
	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventDayDBTypes, true, eventDayPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(eventDayAllColumns, eventDayPrimaryKeyColumns) {
		fields = eventDayAllColumns
	} else {
		fields = strmangle.SetComplement(
			eventDayAllColumns,
			eventDayPrimaryKeyColumns,
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

	slice := EventDaySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEventDaysUpsert(t *testing.T) {
	t.Parallel()

	if len(eventDayAllColumns) == len(eventDayPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EventDay{}
	if err = randomize.Struct(seed, &o, eventDayDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EventDay: %s", err)
	}

	count, err := EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, eventDayDBTypes, false, eventDayPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventDay struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EventDay: %s", err)
	}

	count, err = EventDays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
