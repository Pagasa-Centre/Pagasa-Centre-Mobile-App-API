// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// GooseDBVersion is an object representing the database table.
type GooseDBVersion struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	VersionID int64     `boil:"version_id" json:"version_id" toml:"version_id" yaml:"version_id"`
	IsApplied bool      `boil:"is_applied" json:"is_applied" toml:"is_applied" yaml:"is_applied"`
	Tstamp    time.Time `boil:"tstamp" json:"tstamp" toml:"tstamp" yaml:"tstamp"`

	R *gooseDBVersionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gooseDBVersionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GooseDBVersionColumns = struct {
	ID        string
	VersionID string
	IsApplied string
	Tstamp    string
}{
	ID:        "id",
	VersionID: "version_id",
	IsApplied: "is_applied",
	Tstamp:    "tstamp",
}

var GooseDBVersionTableColumns = struct {
	ID        string
	VersionID string
	IsApplied string
	Tstamp    string
}{
	ID:        "goose_db_version.id",
	VersionID: "goose_db_version.version_id",
	IsApplied: "goose_db_version.is_applied",
	Tstamp:    "goose_db_version.tstamp",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var GooseDBVersionWhere = struct {
	ID        whereHelperint
	VersionID whereHelperint64
	IsApplied whereHelperbool
	Tstamp    whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"goose_db_version\".\"id\""},
	VersionID: whereHelperint64{field: "\"goose_db_version\".\"version_id\""},
	IsApplied: whereHelperbool{field: "\"goose_db_version\".\"is_applied\""},
	Tstamp:    whereHelpertime_Time{field: "\"goose_db_version\".\"tstamp\""},
}

// GooseDBVersionRels is where relationship names are stored.
var GooseDBVersionRels = struct {
}{}

// gooseDBVersionR is where relationships are stored.
type gooseDBVersionR struct {
}

// NewStruct creates a new relationship struct
func (*gooseDBVersionR) NewStruct() *gooseDBVersionR {
	return &gooseDBVersionR{}
}

// gooseDBVersionL is where Load methods for each relationship are stored.
type gooseDBVersionL struct{}

var (
	gooseDBVersionAllColumns            = []string{"id", "version_id", "is_applied", "tstamp"}
	gooseDBVersionColumnsWithoutDefault = []string{"version_id", "is_applied"}
	gooseDBVersionColumnsWithDefault    = []string{"id", "tstamp"}
	gooseDBVersionPrimaryKeyColumns     = []string{"id"}
	gooseDBVersionGeneratedColumns      = []string{}
)

type (
	// GooseDBVersionSlice is an alias for a slice of pointers to GooseDBVersion.
	// This should almost always be used instead of []GooseDBVersion.
	GooseDBVersionSlice []*GooseDBVersion
	// GooseDBVersionHook is the signature for custom GooseDBVersion hook methods
	GooseDBVersionHook func(context.Context, boil.ContextExecutor, *GooseDBVersion) error

	gooseDBVersionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gooseDBVersionType                 = reflect.TypeOf(&GooseDBVersion{})
	gooseDBVersionMapping              = queries.MakeStructMapping(gooseDBVersionType)
	gooseDBVersionPrimaryKeyMapping, _ = queries.BindMapping(gooseDBVersionType, gooseDBVersionMapping, gooseDBVersionPrimaryKeyColumns)
	gooseDBVersionInsertCacheMut       sync.RWMutex
	gooseDBVersionInsertCache          = make(map[string]insertCache)
	gooseDBVersionUpdateCacheMut       sync.RWMutex
	gooseDBVersionUpdateCache          = make(map[string]updateCache)
	gooseDBVersionUpsertCacheMut       sync.RWMutex
	gooseDBVersionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gooseDBVersionAfterSelectMu sync.Mutex
var gooseDBVersionAfterSelectHooks []GooseDBVersionHook

var gooseDBVersionBeforeInsertMu sync.Mutex
var gooseDBVersionBeforeInsertHooks []GooseDBVersionHook
var gooseDBVersionAfterInsertMu sync.Mutex
var gooseDBVersionAfterInsertHooks []GooseDBVersionHook

var gooseDBVersionBeforeUpdateMu sync.Mutex
var gooseDBVersionBeforeUpdateHooks []GooseDBVersionHook
var gooseDBVersionAfterUpdateMu sync.Mutex
var gooseDBVersionAfterUpdateHooks []GooseDBVersionHook

var gooseDBVersionBeforeDeleteMu sync.Mutex
var gooseDBVersionBeforeDeleteHooks []GooseDBVersionHook
var gooseDBVersionAfterDeleteMu sync.Mutex
var gooseDBVersionAfterDeleteHooks []GooseDBVersionHook

var gooseDBVersionBeforeUpsertMu sync.Mutex
var gooseDBVersionBeforeUpsertHooks []GooseDBVersionHook
var gooseDBVersionAfterUpsertMu sync.Mutex
var gooseDBVersionAfterUpsertHooks []GooseDBVersionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GooseDBVersion) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GooseDBVersion) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GooseDBVersion) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GooseDBVersion) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GooseDBVersion) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GooseDBVersion) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GooseDBVersion) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GooseDBVersion) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GooseDBVersion) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseDBVersionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGooseDBVersionHook registers your hook function for all future operations.
func AddGooseDBVersionHook(hookPoint boil.HookPoint, gooseDBVersionHook GooseDBVersionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		gooseDBVersionAfterSelectMu.Lock()
		gooseDBVersionAfterSelectHooks = append(gooseDBVersionAfterSelectHooks, gooseDBVersionHook)
		gooseDBVersionAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		gooseDBVersionBeforeInsertMu.Lock()
		gooseDBVersionBeforeInsertHooks = append(gooseDBVersionBeforeInsertHooks, gooseDBVersionHook)
		gooseDBVersionBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		gooseDBVersionAfterInsertMu.Lock()
		gooseDBVersionAfterInsertHooks = append(gooseDBVersionAfterInsertHooks, gooseDBVersionHook)
		gooseDBVersionAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		gooseDBVersionBeforeUpdateMu.Lock()
		gooseDBVersionBeforeUpdateHooks = append(gooseDBVersionBeforeUpdateHooks, gooseDBVersionHook)
		gooseDBVersionBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		gooseDBVersionAfterUpdateMu.Lock()
		gooseDBVersionAfterUpdateHooks = append(gooseDBVersionAfterUpdateHooks, gooseDBVersionHook)
		gooseDBVersionAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		gooseDBVersionBeforeDeleteMu.Lock()
		gooseDBVersionBeforeDeleteHooks = append(gooseDBVersionBeforeDeleteHooks, gooseDBVersionHook)
		gooseDBVersionBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		gooseDBVersionAfterDeleteMu.Lock()
		gooseDBVersionAfterDeleteHooks = append(gooseDBVersionAfterDeleteHooks, gooseDBVersionHook)
		gooseDBVersionAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		gooseDBVersionBeforeUpsertMu.Lock()
		gooseDBVersionBeforeUpsertHooks = append(gooseDBVersionBeforeUpsertHooks, gooseDBVersionHook)
		gooseDBVersionBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		gooseDBVersionAfterUpsertMu.Lock()
		gooseDBVersionAfterUpsertHooks = append(gooseDBVersionAfterUpsertHooks, gooseDBVersionHook)
		gooseDBVersionAfterUpsertMu.Unlock()
	}
}

// One returns a single gooseDBVersion record from the query.
func (q gooseDBVersionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GooseDBVersion, error) {
	o := &GooseDBVersion{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for goose_db_version")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GooseDBVersion records from the query.
func (q gooseDBVersionQuery) All(ctx context.Context, exec boil.ContextExecutor) (GooseDBVersionSlice, error) {
	var o []*GooseDBVersion

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to GooseDBVersion slice")
	}

	if len(gooseDBVersionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GooseDBVersion records in the query.
func (q gooseDBVersionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count goose_db_version rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gooseDBVersionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if goose_db_version exists")
	}

	return count > 0, nil
}

// GooseDBVersions retrieves all the records using an executor.
func GooseDBVersions(mods ...qm.QueryMod) gooseDBVersionQuery {
	mods = append(mods, qm.From("\"goose_db_version\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"goose_db_version\".*"})
	}

	return gooseDBVersionQuery{q}
}

// FindGooseDBVersion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGooseDBVersion(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*GooseDBVersion, error) {
	gooseDBVersionObj := &GooseDBVersion{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"goose_db_version\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gooseDBVersionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from goose_db_version")
	}

	if err = gooseDBVersionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gooseDBVersionObj, err
	}

	return gooseDBVersionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GooseDBVersion) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no goose_db_version provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gooseDBVersionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gooseDBVersionInsertCacheMut.RLock()
	cache, cached := gooseDBVersionInsertCache[key]
	gooseDBVersionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gooseDBVersionAllColumns,
			gooseDBVersionColumnsWithDefault,
			gooseDBVersionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gooseDBVersionType, gooseDBVersionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gooseDBVersionType, gooseDBVersionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"goose_db_version\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"goose_db_version\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into goose_db_version")
	}

	if !cached {
		gooseDBVersionInsertCacheMut.Lock()
		gooseDBVersionInsertCache[key] = cache
		gooseDBVersionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GooseDBVersion.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GooseDBVersion) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gooseDBVersionUpdateCacheMut.RLock()
	cache, cached := gooseDBVersionUpdateCache[key]
	gooseDBVersionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gooseDBVersionAllColumns,
			gooseDBVersionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update goose_db_version, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"goose_db_version\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gooseDBVersionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gooseDBVersionType, gooseDBVersionMapping, append(wl, gooseDBVersionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update goose_db_version row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for goose_db_version")
	}

	if !cached {
		gooseDBVersionUpdateCacheMut.Lock()
		gooseDBVersionUpdateCache[key] = cache
		gooseDBVersionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gooseDBVersionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for goose_db_version")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for goose_db_version")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GooseDBVersionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gooseDBVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"goose_db_version\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gooseDBVersionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in gooseDBVersion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all gooseDBVersion")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GooseDBVersion) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no goose_db_version provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gooseDBVersionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	gooseDBVersionUpsertCacheMut.RLock()
	cache, cached := gooseDBVersionUpsertCache[key]
	gooseDBVersionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			gooseDBVersionAllColumns,
			gooseDBVersionColumnsWithDefault,
			gooseDBVersionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gooseDBVersionAllColumns,
			gooseDBVersionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert goose_db_version, could not build update column list")
		}

		ret := strmangle.SetComplement(gooseDBVersionAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(gooseDBVersionPrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert goose_db_version, could not build conflict column list")
			}

			conflict = make([]string, len(gooseDBVersionPrimaryKeyColumns))
			copy(conflict, gooseDBVersionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"goose_db_version\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(gooseDBVersionType, gooseDBVersionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gooseDBVersionType, gooseDBVersionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert goose_db_version")
	}

	if !cached {
		gooseDBVersionUpsertCacheMut.Lock()
		gooseDBVersionUpsertCache[key] = cache
		gooseDBVersionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GooseDBVersion record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GooseDBVersion) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no GooseDBVersion provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gooseDBVersionPrimaryKeyMapping)
	sql := "DELETE FROM \"goose_db_version\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from goose_db_version")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for goose_db_version")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gooseDBVersionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no gooseDBVersionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from goose_db_version")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for goose_db_version")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GooseDBVersionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gooseDBVersionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gooseDBVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"goose_db_version\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gooseDBVersionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from gooseDBVersion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for goose_db_version")
	}

	if len(gooseDBVersionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GooseDBVersion) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGooseDBVersion(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GooseDBVersionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GooseDBVersionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gooseDBVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"goose_db_version\".* FROM \"goose_db_version\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gooseDBVersionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in GooseDBVersionSlice")
	}

	*o = slice

	return nil
}

// GooseDBVersionExists checks if the GooseDBVersion row exists.
func GooseDBVersionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"goose_db_version\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if goose_db_version exists")
	}

	return exists, nil
}

// Exists checks if the GooseDBVersion row exists.
func (o *GooseDBVersion) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GooseDBVersionExists(ctx, exec, o.ID)
}
