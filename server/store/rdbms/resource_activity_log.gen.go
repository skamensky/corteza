package rdbms

// This file is an auto-generated file
//
// Template:    pkg/codegen/assets/store_rdbms.gen.go.tpl
// Definitions: store/resource_activity_log.yaml
//
// Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated.

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/discovery/types"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/cortezaproject/corteza-server/store"
)

var _ = errors.Is

// SearchResourceActivityLogs returns all matching rows
//
// This function calls convertResourceActivityLogFilter with the given
// types.ResourceActivityFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchResourceActivityLogs(ctx context.Context, f types.ResourceActivityFilter) (types.ResourceActivitySet, types.ResourceActivityFilter, error) {
	var (
		err error
		set []*types.ResourceActivity
		q   squirrel.SelectBuilder
	)

	return set, f, func() error {
		q, err = s.convertResourceActivityLogFilter(f)
		if err != nil {
			return err
		}

		set, err = s.QueryResourceActivityLogs(ctx, q, nil)
		return err
	}()
}

// QueryResourceActivityLogs queries the database, converts and checks each row and
// returns collected set
//
// Fn also returns total number of fetched items and last fetched item so that the caller can construct cursor
// for next page of results
func (s Store) QueryResourceActivityLogs(
	ctx context.Context,
	q squirrel.Sqlizer,
	check func(*types.ResourceActivity) (bool, error),
) ([]*types.ResourceActivity, error) {
	var (
		tmp = make([]*types.ResourceActivity, 0, DefaultSliceCapacity)
		set = make([]*types.ResourceActivity, 0, DefaultSliceCapacity)
		res *types.ResourceActivity

		// Query rows with
		rows, err = s.Query(ctx, q)
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Err(); err == nil {
			res, err = s.internalResourceActivityLogRowScanner(rows)
		}

		if err != nil {
			return nil, err
		}

		tmp = append(tmp, res)
	}

	for _, res = range tmp {

		set = append(set, res)
	}

	return set, nil
}

// CreateResourceActivityLog creates one or more rows in resource_activity_log table
func (s Store) CreateResourceActivityLog(ctx context.Context, rr ...*types.ResourceActivity) (err error) {
	for _, res := range rr {
		err = s.checkResourceActivityLogConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execCreateResourceActivityLogs(ctx, s.internalResourceActivityLogEncoder(res))
		if err != nil {
			return err
		}
	}

	return
}

// TruncateResourceActivityLogs Deletes all rows from the resource_activity_log table
func (s Store) TruncateResourceActivityLogs(ctx context.Context) error {
	return s.Truncate(ctx, s.resourceActivityLogTable())
}

// execLookupResourceActivityLog prepares ResourceActivityLog query and executes it,
// returning types.ResourceActivity (or error)
func (s Store) execLookupResourceActivityLog(ctx context.Context, cnd squirrel.Sqlizer) (res *types.ResourceActivity, err error) {
	var (
		row rowScanner
	)

	row, err = s.QueryRow(ctx, s.resourceActivityLogsSelectBuilder().Where(cnd))
	if err != nil {
		return
	}

	res, err = s.internalResourceActivityLogRowScanner(row)
	if err != nil {
		return
	}

	return res, nil
}

// execCreateResourceActivityLogs updates all matched (by cnd) rows in resource_activity_log with given data
func (s Store) execCreateResourceActivityLogs(ctx context.Context, payload store.Payload) error {
	return s.Exec(ctx, s.InsertBuilder(s.resourceActivityLogTable()).SetMap(payload))
}

func (s Store) internalResourceActivityLogRowScanner(row rowScanner) (res *types.ResourceActivity, err error) {
	res = &types.ResourceActivity{}

	err = row.Scan(
		&res.ID,
		&res.ResourceID,
		&res.ResourceType,
		&res.ResourceAction,
		&res.Timestamp,
		&res.Meta,
	)

	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound.Stack(1)
	}

	if err != nil {
		return nil, errors.Store("could not scan resourceActivityLog db row: %s", err).Wrap(err)
	} else {
		return res, nil
	}
}

// QueryResourceActivityLogs returns squirrel.SelectBuilder with set table and all columns
func (s Store) resourceActivityLogsSelectBuilder() squirrel.SelectBuilder {
	return s.SelectBuilder(s.resourceActivityLogTable("ral"), s.resourceActivityLogColumns("ral")...)
}

// resourceActivityLogTable name of the db table
func (Store) resourceActivityLogTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "resource_activity_log" + alias
}

// ResourceActivityLogColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) resourceActivityLogColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "rel_resource",
		alias + "resource_type",
		alias + "resource_action",
		alias + "ts",
		alias + "meta",
	}
}

// {true true false false false false}

// internalResourceActivityLogEncoder encodes fields from types.ResourceActivity to store.Payload (map)
//
// Encoding is done by using generic approach or by calling encodeResourceActivityLog
// func when rdbms.customEncoder=true
func (s Store) internalResourceActivityLogEncoder(res *types.ResourceActivity) store.Payload {
	return store.Payload{
		"id":              res.ID,
		"rel_resource":    res.ResourceID,
		"resource_type":   res.ResourceType,
		"resource_action": res.ResourceAction,
		"ts":              res.Timestamp,
		"meta":            res.Meta,
	}
}

// checkResourceActivityLogConstraints performs lookups (on valid) resource to check if any of the values on unique fields
// already exists in the store
//
// Using built-in constraint checking would be more performant but unfortunately we cannot rely
// on the full support (MySQL does not support conditional indexes)
func (s *Store) checkResourceActivityLogConstraints(ctx context.Context, res *types.ResourceActivity) error {
	// Consider resource valid when all fields in unique constraint check lookups
	// have valid (non-empty) value
	//
	// Only string and uint64 are supported for now
	// feel free to add additional types if needed
	var valid = true

	if !valid {
		return nil
	}

	var checks = make([]func() error, 0)

	for _, check := range checks {
		if err := check(); err != nil {
			return err
		}
	}

	return nil
}
