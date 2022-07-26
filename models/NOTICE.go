// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// NOTICE is an object representing the database table.
type NOTICE struct { // お知らせID
	NOTICE_ID int `boil:"NOTICE_ID" json:"NOTICE_ID" toml:"NOTICE_ID" yaml:"NOTICE_ID"`
	// OEMID
	OEM_ID int `boil:"OEM_ID" json:"OEM_ID" toml:"OEM_ID" yaml:"OEM_ID"`
	// タイトル
	TITLE string `boil:"TITLE" json:"TITLE" toml:"TITLE" yaml:"TITLE"`
	// 本文
	BODY string `boil:"BODY" json:"BODY" toml:"BODY" yaml:"BODY"`
	// 提示日
	PRE_DTE time.Time `boil:"PRE_DTE" json:"PRE_DTE" toml:"PRE_DTE" yaml:"PRE_DTE"`
	// 削除フラグ
	DELE_FLG string `boil:"DELE_FLG" json:"DELE_FLG" toml:"DELE_FLG" yaml:"DELE_FLG"`
	// 更新日時
	UPDA_DTE time.Time `boil:"UPDA_DTE" json:"UPDA_DTE" toml:"UPDA_DTE" yaml:"UPDA_DTE"`
	// 更新者ID
	UPDA_USER_ID int `boil:"UPDA_USER_ID" json:"UPDA_USER_ID" toml:"UPDA_USER_ID" yaml:"UPDA_USER_ID"`
	// 作成日時
	CREA_DTE time.Time `boil:"CREA_DTE" json:"CREA_DTE" toml:"CREA_DTE" yaml:"CREA_DTE"`
	// 作成者ID
	CREA_USER_ID int `boil:"CREA_USER_ID" json:"CREA_USER_ID" toml:"CREA_USER_ID" yaml:"CREA_USER_ID"`

	R *nOTICER `boil:"-" json:"-" toml:"-" yaml:"-"`
	L nOTICEL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NOTICEColumns = struct {
	NOTICE_ID    string
	OEM_ID       string
	TITLE        string
	BODY         string
	PRE_DTE      string
	DELE_FLG     string
	UPDA_DTE     string
	UPDA_USER_ID string
	CREA_DTE     string
	CREA_USER_ID string
}{
	NOTICE_ID:    "NOTICE_ID",
	OEM_ID:       "OEM_ID",
	TITLE:        "TITLE",
	BODY:         "BODY",
	PRE_DTE:      "PRE_DTE",
	DELE_FLG:     "DELE_FLG",
	UPDA_DTE:     "UPDA_DTE",
	UPDA_USER_ID: "UPDA_USER_ID",
	CREA_DTE:     "CREA_DTE",
	CREA_USER_ID: "CREA_USER_ID",
}

var NOTICETableColumns = struct {
	NOTICE_ID    string
	OEM_ID       string
	TITLE        string
	BODY         string
	PRE_DTE      string
	DELE_FLG     string
	UPDA_DTE     string
	UPDA_USER_ID string
	CREA_DTE     string
	CREA_USER_ID string
}{
	NOTICE_ID:    "NOTICE.NOTICE_ID",
	OEM_ID:       "NOTICE.OEM_ID",
	TITLE:        "NOTICE.TITLE",
	BODY:         "NOTICE.BODY",
	PRE_DTE:      "NOTICE.PRE_DTE",
	DELE_FLG:     "NOTICE.DELE_FLG",
	UPDA_DTE:     "NOTICE.UPDA_DTE",
	UPDA_USER_ID: "NOTICE.UPDA_USER_ID",
	CREA_DTE:     "NOTICE.CREA_DTE",
	CREA_USER_ID: "NOTICE.CREA_USER_ID",
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

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

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

var NOTICEWhere = struct {
	NOTICE_ID    whereHelperint
	OEM_ID       whereHelperint
	TITLE        whereHelperstring
	BODY         whereHelperstring
	PRE_DTE      whereHelpertime_Time
	DELE_FLG     whereHelperstring
	UPDA_DTE     whereHelpertime_Time
	UPDA_USER_ID whereHelperint
	CREA_DTE     whereHelpertime_Time
	CREA_USER_ID whereHelperint
}{
	NOTICE_ID:    whereHelperint{field: "`NOTICE`.`NOTICE_ID`"},
	OEM_ID:       whereHelperint{field: "`NOTICE`.`OEM_ID`"},
	TITLE:        whereHelperstring{field: "`NOTICE`.`TITLE`"},
	BODY:         whereHelperstring{field: "`NOTICE`.`BODY`"},
	PRE_DTE:      whereHelpertime_Time{field: "`NOTICE`.`PRE_DTE`"},
	DELE_FLG:     whereHelperstring{field: "`NOTICE`.`DELE_FLG`"},
	UPDA_DTE:     whereHelpertime_Time{field: "`NOTICE`.`UPDA_DTE`"},
	UPDA_USER_ID: whereHelperint{field: "`NOTICE`.`UPDA_USER_ID`"},
	CREA_DTE:     whereHelpertime_Time{field: "`NOTICE`.`CREA_DTE`"},
	CREA_USER_ID: whereHelperint{field: "`NOTICE`.`CREA_USER_ID`"},
}

// NOTICERels is where relationship names are stored.
var NOTICERels = struct {
}{}

// nOTICER is where relationships are stored.
type nOTICER struct {
}

// NewStruct creates a new relationship struct
func (*nOTICER) NewStruct() *nOTICER {
	return &nOTICER{}
}

// nOTICEL is where Load methods for each relationship are stored.
type nOTICEL struct{}

var (
	nOTICEAllColumns            = []string{"NOTICE_ID", "OEM_ID", "TITLE", "BODY", "PRE_DTE", "DELE_FLG", "UPDA_DTE", "UPDA_USER_ID", "CREA_DTE", "CREA_USER_ID"}
	nOTICEColumnsWithoutDefault = []string{"OEM_ID", "TITLE", "BODY"}
	nOTICEColumnsWithDefault    = []string{"NOTICE_ID", "PRE_DTE", "DELE_FLG", "UPDA_DTE", "UPDA_USER_ID", "CREA_DTE", "CREA_USER_ID"}
	nOTICEPrimaryKeyColumns     = []string{"NOTICE_ID"}
	nOTICEGeneratedColumns      = []string{}
)

type (
	// NOTICESlice is an alias for a slice of pointers to NOTICE.
	// This should almost always be used instead of []NOTICE.
	NOTICESlice []*NOTICE
	// NOTICEHook is the signature for custom NOTICE hook methods
	NOTICEHook func(context.Context, boil.ContextExecutor, *NOTICE) error

	nOTICEQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	nOTICEType                 = reflect.TypeOf(&NOTICE{})
	nOTICEMapping              = queries.MakeStructMapping(nOTICEType)
	nOTICEPrimaryKeyMapping, _ = queries.BindMapping(nOTICEType, nOTICEMapping, nOTICEPrimaryKeyColumns)
	nOTICEInsertCacheMut       sync.RWMutex
	nOTICEInsertCache          = make(map[string]insertCache)
	nOTICEUpdateCacheMut       sync.RWMutex
	nOTICEUpdateCache          = make(map[string]updateCache)
	nOTICEUpsertCacheMut       sync.RWMutex
	nOTICEUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var nOTICEAfterSelectHooks []NOTICEHook

var nOTICEBeforeInsertHooks []NOTICEHook
var nOTICEAfterInsertHooks []NOTICEHook

var nOTICEBeforeUpdateHooks []NOTICEHook
var nOTICEAfterUpdateHooks []NOTICEHook

var nOTICEBeforeDeleteHooks []NOTICEHook
var nOTICEAfterDeleteHooks []NOTICEHook

var nOTICEBeforeUpsertHooks []NOTICEHook
var nOTICEAfterUpsertHooks []NOTICEHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NOTICE) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NOTICE) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NOTICE) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NOTICE) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NOTICE) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NOTICE) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NOTICE) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NOTICE) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NOTICE) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nOTICEAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNOTICEHook registers your hook function for all future operations.
func AddNOTICEHook(hookPoint boil.HookPoint, nOTICEHook NOTICEHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		nOTICEAfterSelectHooks = append(nOTICEAfterSelectHooks, nOTICEHook)
	case boil.BeforeInsertHook:
		nOTICEBeforeInsertHooks = append(nOTICEBeforeInsertHooks, nOTICEHook)
	case boil.AfterInsertHook:
		nOTICEAfterInsertHooks = append(nOTICEAfterInsertHooks, nOTICEHook)
	case boil.BeforeUpdateHook:
		nOTICEBeforeUpdateHooks = append(nOTICEBeforeUpdateHooks, nOTICEHook)
	case boil.AfterUpdateHook:
		nOTICEAfterUpdateHooks = append(nOTICEAfterUpdateHooks, nOTICEHook)
	case boil.BeforeDeleteHook:
		nOTICEBeforeDeleteHooks = append(nOTICEBeforeDeleteHooks, nOTICEHook)
	case boil.AfterDeleteHook:
		nOTICEAfterDeleteHooks = append(nOTICEAfterDeleteHooks, nOTICEHook)
	case boil.BeforeUpsertHook:
		nOTICEBeforeUpsertHooks = append(nOTICEBeforeUpsertHooks, nOTICEHook)
	case boil.AfterUpsertHook:
		nOTICEAfterUpsertHooks = append(nOTICEAfterUpsertHooks, nOTICEHook)
	}
}

// One returns a single nOTICE record from the query.
func (q nOTICEQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NOTICE, error) {
	o := &NOTICE{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for NOTICE")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NOTICE records from the query.
func (q nOTICEQuery) All(ctx context.Context, exec boil.ContextExecutor) (NOTICESlice, error) {
	var o []*NOTICE

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NOTICE slice")
	}

	if len(nOTICEAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NOTICE records in the query.
func (q nOTICEQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count NOTICE rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q nOTICEQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if NOTICE exists")
	}

	return count > 0, nil
}

// NOTICES retrieves all the records using an executor.
func NOTICES(mods ...qm.QueryMod) nOTICEQuery {
	mods = append(mods, qm.From("`NOTICE`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`NOTICE`.*"})
	}

	return nOTICEQuery{q}
}

// FindNOTICE retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNOTICE(ctx context.Context, exec boil.ContextExecutor, nOTICEID int, selectCols ...string) (*NOTICE, error) {
	nOTICEObj := &NOTICE{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `NOTICE` where `NOTICE_ID`=?", sel,
	)

	q := queries.Raw(query, nOTICEID)

	err := q.Bind(ctx, exec, nOTICEObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from NOTICE")
	}

	if err = nOTICEObj.doAfterSelectHooks(ctx, exec); err != nil {
		return nOTICEObj, err
	}

	return nOTICEObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NOTICE) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no NOTICE provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(nOTICEColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	nOTICEInsertCacheMut.RLock()
	cache, cached := nOTICEInsertCache[key]
	nOTICEInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			nOTICEAllColumns,
			nOTICEColumnsWithDefault,
			nOTICEColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(nOTICEType, nOTICEMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(nOTICEType, nOTICEMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `NOTICE` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `NOTICE` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `NOTICE` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, nOTICEPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into NOTICE")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.NOTICE_ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == nOTICEMapping["NOTICE_ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.NOTICE_ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for NOTICE")
	}

CacheNoHooks:
	if !cached {
		nOTICEInsertCacheMut.Lock()
		nOTICEInsertCache[key] = cache
		nOTICEInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NOTICE.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NOTICE) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	nOTICEUpdateCacheMut.RLock()
	cache, cached := nOTICEUpdateCache[key]
	nOTICEUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			nOTICEAllColumns,
			nOTICEPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update NOTICE, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `NOTICE` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, nOTICEPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(nOTICEType, nOTICEMapping, append(wl, nOTICEPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update NOTICE row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for NOTICE")
	}

	if !cached {
		nOTICEUpdateCacheMut.Lock()
		nOTICEUpdateCache[key] = cache
		nOTICEUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q nOTICEQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for NOTICE")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for NOTICE")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NOTICESlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nOTICEPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `NOTICE` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nOTICEPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in nOTICE slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all nOTICE")
	}
	return rowsAff, nil
}

var mySQLNOTICEUniqueColumns = []string{
	"NOTICE_ID",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NOTICE) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no NOTICE provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(nOTICEColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLNOTICEUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	nOTICEUpsertCacheMut.RLock()
	cache, cached := nOTICEUpsertCache[key]
	nOTICEUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			nOTICEAllColumns,
			nOTICEColumnsWithDefault,
			nOTICEColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			nOTICEAllColumns,
			nOTICEPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert NOTICE, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`NOTICE`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `NOTICE` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(nOTICEType, nOTICEMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(nOTICEType, nOTICEMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for NOTICE")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.NOTICE_ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == nOTICEMapping["NOTICE_ID"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(nOTICEType, nOTICEMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for NOTICE")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for NOTICE")
	}

CacheNoHooks:
	if !cached {
		nOTICEUpsertCacheMut.Lock()
		nOTICEUpsertCache[key] = cache
		nOTICEUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single NOTICE record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NOTICE) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NOTICE provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), nOTICEPrimaryKeyMapping)
	sql := "DELETE FROM `NOTICE` WHERE `NOTICE_ID`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from NOTICE")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for NOTICE")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q nOTICEQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no nOTICEQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from NOTICE")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for NOTICE")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NOTICESlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(nOTICEBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nOTICEPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `NOTICE` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nOTICEPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from nOTICE slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for NOTICE")
	}

	if len(nOTICEAfterDeleteHooks) != 0 {
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
func (o *NOTICE) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNOTICE(ctx, exec, o.NOTICE_ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NOTICESlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NOTICESlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nOTICEPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `NOTICE`.* FROM `NOTICE` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nOTICEPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NOTICESlice")
	}

	*o = slice

	return nil
}

// NOTICEExists checks if the NOTICE row exists.
func NOTICEExists(ctx context.Context, exec boil.ContextExecutor, nOTICEID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `NOTICE` where `NOTICE_ID`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, nOTICEID)
	}
	row := exec.QueryRowContext(ctx, sql, nOTICEID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if NOTICE exists")
	}

	return exists, nil
}
