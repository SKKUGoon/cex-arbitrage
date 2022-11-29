// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kimchi/ent/predicate"
	"kimchi/ent/tradelog"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TradeLogQuery is the builder for querying TradeLog entities.
type TradeLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TradeLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TradeLogQuery builder.
func (tlq *TradeLogQuery) Where(ps ...predicate.TradeLog) *TradeLogQuery {
	tlq.predicates = append(tlq.predicates, ps...)
	return tlq
}

// Limit adds a limit step to the query.
func (tlq *TradeLogQuery) Limit(limit int) *TradeLogQuery {
	tlq.limit = &limit
	return tlq
}

// Offset adds an offset step to the query.
func (tlq *TradeLogQuery) Offset(offset int) *TradeLogQuery {
	tlq.offset = &offset
	return tlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tlq *TradeLogQuery) Unique(unique bool) *TradeLogQuery {
	tlq.unique = &unique
	return tlq
}

// Order adds an order step to the query.
func (tlq *TradeLogQuery) Order(o ...OrderFunc) *TradeLogQuery {
	tlq.order = append(tlq.order, o...)
	return tlq
}

// First returns the first TradeLog entity from the query.
// Returns a *NotFoundError when no TradeLog was found.
func (tlq *TradeLogQuery) First(ctx context.Context) (*TradeLog, error) {
	nodes, err := tlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tradelog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tlq *TradeLogQuery) FirstX(ctx context.Context) *TradeLog {
	node, err := tlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TradeLog ID from the query.
// Returns a *NotFoundError when no TradeLog ID was found.
func (tlq *TradeLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tradelog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tlq *TradeLogQuery) FirstIDX(ctx context.Context) int {
	id, err := tlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TradeLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TradeLog entity is found.
// Returns a *NotFoundError when no TradeLog entities are found.
func (tlq *TradeLogQuery) Only(ctx context.Context) (*TradeLog, error) {
	nodes, err := tlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tradelog.Label}
	default:
		return nil, &NotSingularError{tradelog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tlq *TradeLogQuery) OnlyX(ctx context.Context) *TradeLog {
	node, err := tlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TradeLog ID in the query.
// Returns a *NotSingularError when more than one TradeLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (tlq *TradeLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tradelog.Label}
	default:
		err = &NotSingularError{tradelog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tlq *TradeLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := tlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TradeLogs.
func (tlq *TradeLogQuery) All(ctx context.Context) ([]*TradeLog, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tlq *TradeLogQuery) AllX(ctx context.Context) []*TradeLog {
	nodes, err := tlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TradeLog IDs.
func (tlq *TradeLogQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tlq.Select(tradelog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tlq *TradeLogQuery) IDsX(ctx context.Context) []int {
	ids, err := tlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tlq *TradeLogQuery) Count(ctx context.Context) (int, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tlq *TradeLogQuery) CountX(ctx context.Context) int {
	count, err := tlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tlq *TradeLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tlq *TradeLogQuery) ExistX(ctx context.Context) bool {
	exist, err := tlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TradeLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tlq *TradeLogQuery) Clone() *TradeLogQuery {
	if tlq == nil {
		return nil
	}
	return &TradeLogQuery{
		config:     tlq.config,
		limit:      tlq.limit,
		offset:     tlq.offset,
		order:      append([]OrderFunc{}, tlq.order...),
		predicates: append([]predicate.TradeLog{}, tlq.predicates...),
		// clone intermediate query.
		sql:    tlq.sql.Clone(),
		path:   tlq.path,
		unique: tlq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TradeLog.Query().
//		GroupBy(tradelog.FieldDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tlq *TradeLogQuery) GroupBy(field string, fields ...string) *TradeLogGroupBy {
	grbuild := &TradeLogGroupBy{config: tlq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tlq.sqlQuery(ctx), nil
	}
	grbuild.label = tradelog.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//	}
//
//	client.TradeLog.Query().
//		Select(tradelog.FieldDate).
//		Scan(ctx, &v)
func (tlq *TradeLogQuery) Select(fields ...string) *TradeLogSelect {
	tlq.fields = append(tlq.fields, fields...)
	selbuild := &TradeLogSelect{TradeLogQuery: tlq}
	selbuild.label = tradelog.Label
	selbuild.flds, selbuild.scan = &tlq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a TradeLogSelect configured with the given aggregations.
func (tlq *TradeLogQuery) Aggregate(fns ...AggregateFunc) *TradeLogSelect {
	return tlq.Select().Aggregate(fns...)
}

func (tlq *TradeLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tlq.fields {
		if !tradelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tlq.path != nil {
		prev, err := tlq.path(ctx)
		if err != nil {
			return err
		}
		tlq.sql = prev
	}
	return nil
}

func (tlq *TradeLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TradeLog, error) {
	var (
		nodes = []*TradeLog{}
		_spec = tlq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TradeLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TradeLog{config: tlq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tlq *TradeLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tlq.querySpec()
	_spec.Node.Columns = tlq.fields
	if len(tlq.fields) > 0 {
		_spec.Unique = tlq.unique != nil && *tlq.unique
	}
	return sqlgraph.CountNodes(ctx, tlq.driver, _spec)
}

func (tlq *TradeLogQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := tlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (tlq *TradeLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tradelog.Table,
			Columns: tradelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradelog.FieldID,
			},
		},
		From:   tlq.sql,
		Unique: true,
	}
	if unique := tlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tradelog.FieldID)
		for i := range fields {
			if fields[i] != tradelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tlq *TradeLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tlq.driver.Dialect())
	t1 := builder.Table(tradelog.Table)
	columns := tlq.fields
	if len(columns) == 0 {
		columns = tradelog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tlq.sql != nil {
		selector = tlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tlq.unique != nil && *tlq.unique {
		selector.Distinct()
	}
	for _, p := range tlq.predicates {
		p(selector)
	}
	for _, p := range tlq.order {
		p(selector)
	}
	if offset := tlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TradeLogGroupBy is the group-by builder for TradeLog entities.
type TradeLogGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlgb *TradeLogGroupBy) Aggregate(fns ...AggregateFunc) *TradeLogGroupBy {
	tlgb.fns = append(tlgb.fns, fns...)
	return tlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tlgb *TradeLogGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tlgb.path(ctx)
	if err != nil {
		return err
	}
	tlgb.sql = query
	return tlgb.sqlScan(ctx, v)
}

func (tlgb *TradeLogGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tlgb.fields {
		if !tradelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tlgb *TradeLogGroupBy) sqlQuery() *sql.Selector {
	selector := tlgb.sql.Select()
	aggregation := make([]string, 0, len(tlgb.fns))
	for _, fn := range tlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tlgb.fields)+len(tlgb.fns))
		for _, f := range tlgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tlgb.fields...)...)
}

// TradeLogSelect is the builder for selecting fields of TradeLog entities.
type TradeLogSelect struct {
	*TradeLogQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tls *TradeLogSelect) Aggregate(fns ...AggregateFunc) *TradeLogSelect {
	tls.fns = append(tls.fns, fns...)
	return tls
}

// Scan applies the selector query and scans the result into the given value.
func (tls *TradeLogSelect) Scan(ctx context.Context, v any) error {
	if err := tls.prepareQuery(ctx); err != nil {
		return err
	}
	tls.sql = tls.TradeLogQuery.sqlQuery(ctx)
	return tls.sqlScan(ctx, v)
}

func (tls *TradeLogSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(tls.fns))
	for _, fn := range tls.fns {
		aggregation = append(aggregation, fn(tls.sql))
	}
	switch n := len(*tls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		tls.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		tls.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := tls.sql.Query()
	if err := tls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}