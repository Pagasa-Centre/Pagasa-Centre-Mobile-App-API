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

func testMedia(t *testing.T) {
	t.Parallel()

	query := Media()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMediaDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
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

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Media().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MediumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MediumExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Medium exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MediumExists to return true, but got false.")
	}
}

func testMediaFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mediumFound, err := FindMedium(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mediumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMediaBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Media().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMediaOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Media().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMediaAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mediumOne := &Medium{}
	mediumTwo := &Medium{}
	if err = randomize.Struct(seed, mediumOne, mediumDBTypes, false, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}
	if err = randomize.Struct(seed, mediumTwo, mediumDBTypes, false, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mediumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mediumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Media().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMediaCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mediumOne := &Medium{}
	mediumTwo := &Medium{}
	if err = randomize.Struct(seed, mediumOne, mediumDBTypes, false, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}
	if err = randomize.Struct(seed, mediumTwo, mediumDBTypes, false, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mediumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mediumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mediumBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func mediumAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Medium) error {
	*o = Medium{}
	return nil
}

func testMediaHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Medium{}
	o := &Medium{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mediumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Medium object: %s", err)
	}

	AddMediumHook(boil.BeforeInsertHook, mediumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mediumBeforeInsertHooks = []MediumHook{}

	AddMediumHook(boil.AfterInsertHook, mediumAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mediumAfterInsertHooks = []MediumHook{}

	AddMediumHook(boil.AfterSelectHook, mediumAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mediumAfterSelectHooks = []MediumHook{}

	AddMediumHook(boil.BeforeUpdateHook, mediumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mediumBeforeUpdateHooks = []MediumHook{}

	AddMediumHook(boil.AfterUpdateHook, mediumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mediumAfterUpdateHooks = []MediumHook{}

	AddMediumHook(boil.BeforeDeleteHook, mediumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mediumBeforeDeleteHooks = []MediumHook{}

	AddMediumHook(boil.AfterDeleteHook, mediumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mediumAfterDeleteHooks = []MediumHook{}

	AddMediumHook(boil.BeforeUpsertHook, mediumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mediumBeforeUpsertHooks = []MediumHook{}

	AddMediumHook(boil.AfterUpsertHook, mediumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mediumAfterUpsertHooks = []MediumHook{}
}

func testMediaInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMediaInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mediumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMediaReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
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

func testMediaReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MediumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMediaSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Media().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mediumDBTypes = map[string]string{`ID`: `integer`, `Title`: `text`, `Description`: `text`, `YoutubeVideoID`: `character varying`, `Category`: `character varying`, `PublishedAt`: `timestamp without time zone`, `ThumbnailURL`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_             = bytes.MinRead
)

func testMediaUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mediumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mediumAllColumns) == len(mediumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMediaSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mediumAllColumns) == len(mediumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Medium{}
	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mediumDBTypes, true, mediumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mediumAllColumns, mediumPrimaryKeyColumns) {
		fields = mediumAllColumns
	} else {
		fields = strmangle.SetComplement(
			mediumAllColumns,
			mediumPrimaryKeyColumns,
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

	slice := MediumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMediaUpsert(t *testing.T) {
	t.Parallel()

	if len(mediumAllColumns) == len(mediumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Medium{}
	if err = randomize.Struct(seed, &o, mediumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Medium: %s", err)
	}

	count, err := Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mediumDBTypes, false, mediumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Medium struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Medium: %s", err)
	}

	count, err = Media().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
