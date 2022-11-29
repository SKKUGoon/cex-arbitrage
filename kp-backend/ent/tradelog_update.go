// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kimchi/ent/predicate"
	"kimchi/ent/tradelog"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TradeLogUpdate is the builder for updating TradeLog entities.
type TradeLogUpdate struct {
	config
	hooks    []Hook
	mutation *TradeLogMutation
}

// Where appends a list predicates to the TradeLogUpdate builder.
func (tlu *TradeLogUpdate) Where(ps ...predicate.TradeLog) *TradeLogUpdate {
	tlu.mutation.Where(ps...)
	return tlu
}

// SetDate sets the "date" field.
func (tlu *TradeLogUpdate) SetDate(t time.Time) *TradeLogUpdate {
	tlu.mutation.SetDate(t)
	return tlu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillableDate(t *time.Time) *TradeLogUpdate {
	if t != nil {
		tlu.SetDate(*t)
	}
	return tlu
}

// ClearDate clears the value of the "date" field.
func (tlu *TradeLogUpdate) ClearDate() *TradeLogUpdate {
	tlu.mutation.ClearDate()
	return tlu
}

// SetExchange sets the "exchange" field.
func (tlu *TradeLogUpdate) SetExchange(s string) *TradeLogUpdate {
	tlu.mutation.SetExchange(s)
	return tlu
}

// SetNillableExchange sets the "exchange" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillableExchange(s *string) *TradeLogUpdate {
	if s != nil {
		tlu.SetExchange(*s)
	}
	return tlu
}

// ClearExchange clears the value of the "exchange" field.
func (tlu *TradeLogUpdate) ClearExchange() *TradeLogUpdate {
	tlu.mutation.ClearExchange()
	return tlu
}

// SetTicker sets the "ticker" field.
func (tlu *TradeLogUpdate) SetTicker(s string) *TradeLogUpdate {
	tlu.mutation.SetTicker(s)
	return tlu
}

// SetNillableTicker sets the "ticker" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillableTicker(s *string) *TradeLogUpdate {
	if s != nil {
		tlu.SetTicker(*s)
	}
	return tlu
}

// ClearTicker clears the value of the "ticker" field.
func (tlu *TradeLogUpdate) ClearTicker() *TradeLogUpdate {
	tlu.mutation.ClearTicker()
	return tlu
}

// SetPosition sets the "position" field.
func (tlu *TradeLogUpdate) SetPosition(s string) *TradeLogUpdate {
	tlu.mutation.SetPosition(s)
	return tlu
}

// SetNillablePosition sets the "position" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillablePosition(s *string) *TradeLogUpdate {
	if s != nil {
		tlu.SetPosition(*s)
	}
	return tlu
}

// ClearPosition clears the value of the "position" field.
func (tlu *TradeLogUpdate) ClearPosition() *TradeLogUpdate {
	tlu.mutation.ClearPosition()
	return tlu
}

// SetStrategy sets the "strategy" field.
func (tlu *TradeLogUpdate) SetStrategy(s string) *TradeLogUpdate {
	tlu.mutation.SetStrategy(s)
	return tlu
}

// SetNillableStrategy sets the "strategy" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillableStrategy(s *string) *TradeLogUpdate {
	if s != nil {
		tlu.SetStrategy(*s)
	}
	return tlu
}

// ClearStrategy clears the value of the "strategy" field.
func (tlu *TradeLogUpdate) ClearStrategy() *TradeLogUpdate {
	tlu.mutation.ClearStrategy()
	return tlu
}

// SetPrice sets the "price" field.
func (tlu *TradeLogUpdate) SetPrice(f float64) *TradeLogUpdate {
	tlu.mutation.ResetPrice()
	tlu.mutation.SetPrice(f)
	return tlu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillablePrice(f *float64) *TradeLogUpdate {
	if f != nil {
		tlu.SetPrice(*f)
	}
	return tlu
}

// AddPrice adds f to the "price" field.
func (tlu *TradeLogUpdate) AddPrice(f float64) *TradeLogUpdate {
	tlu.mutation.AddPrice(f)
	return tlu
}

// ClearPrice clears the value of the "price" field.
func (tlu *TradeLogUpdate) ClearPrice() *TradeLogUpdate {
	tlu.mutation.ClearPrice()
	return tlu
}

// SetQuantity sets the "quantity" field.
func (tlu *TradeLogUpdate) SetQuantity(f float64) *TradeLogUpdate {
	tlu.mutation.ResetQuantity()
	tlu.mutation.SetQuantity(f)
	return tlu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillableQuantity(f *float64) *TradeLogUpdate {
	if f != nil {
		tlu.SetQuantity(*f)
	}
	return tlu
}

// AddQuantity adds f to the "quantity" field.
func (tlu *TradeLogUpdate) AddQuantity(f float64) *TradeLogUpdate {
	tlu.mutation.AddQuantity(f)
	return tlu
}

// ClearQuantity clears the value of the "quantity" field.
func (tlu *TradeLogUpdate) ClearQuantity() *TradeLogUpdate {
	tlu.mutation.ClearQuantity()
	return tlu
}

// SetLeverage sets the "leverage" field.
func (tlu *TradeLogUpdate) SetLeverage(i int) *TradeLogUpdate {
	tlu.mutation.ResetLeverage()
	tlu.mutation.SetLeverage(i)
	return tlu
}

// SetNillableLeverage sets the "leverage" field if the given value is not nil.
func (tlu *TradeLogUpdate) SetNillableLeverage(i *int) *TradeLogUpdate {
	if i != nil {
		tlu.SetLeverage(*i)
	}
	return tlu
}

// AddLeverage adds i to the "leverage" field.
func (tlu *TradeLogUpdate) AddLeverage(i int) *TradeLogUpdate {
	tlu.mutation.AddLeverage(i)
	return tlu
}

// ClearLeverage clears the value of the "leverage" field.
func (tlu *TradeLogUpdate) ClearLeverage() *TradeLogUpdate {
	tlu.mutation.ClearLeverage()
	return tlu
}

// Mutation returns the TradeLogMutation object of the builder.
func (tlu *TradeLogUpdate) Mutation() *TradeLogMutation {
	return tlu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tlu *TradeLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tlu.hooks) == 0 {
		affected, err = tlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TradeLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tlu.mutation = mutation
			affected, err = tlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tlu.hooks) - 1; i >= 0; i-- {
			if tlu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlu *TradeLogUpdate) SaveX(ctx context.Context) int {
	affected, err := tlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tlu *TradeLogUpdate) Exec(ctx context.Context) error {
	_, err := tlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlu *TradeLogUpdate) ExecX(ctx context.Context) {
	if err := tlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tlu *TradeLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tradelog.Table,
			Columns: tradelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradelog.FieldID,
			},
		},
	}
	if ps := tlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlu.mutation.Date(); ok {
		_spec.SetField(tradelog.FieldDate, field.TypeTime, value)
	}
	if tlu.mutation.DateCleared() {
		_spec.ClearField(tradelog.FieldDate, field.TypeTime)
	}
	if value, ok := tlu.mutation.Exchange(); ok {
		_spec.SetField(tradelog.FieldExchange, field.TypeString, value)
	}
	if tlu.mutation.ExchangeCleared() {
		_spec.ClearField(tradelog.FieldExchange, field.TypeString)
	}
	if value, ok := tlu.mutation.Ticker(); ok {
		_spec.SetField(tradelog.FieldTicker, field.TypeString, value)
	}
	if tlu.mutation.TickerCleared() {
		_spec.ClearField(tradelog.FieldTicker, field.TypeString)
	}
	if value, ok := tlu.mutation.Position(); ok {
		_spec.SetField(tradelog.FieldPosition, field.TypeString, value)
	}
	if tlu.mutation.PositionCleared() {
		_spec.ClearField(tradelog.FieldPosition, field.TypeString)
	}
	if value, ok := tlu.mutation.Strategy(); ok {
		_spec.SetField(tradelog.FieldStrategy, field.TypeString, value)
	}
	if tlu.mutation.StrategyCleared() {
		_spec.ClearField(tradelog.FieldStrategy, field.TypeString)
	}
	if value, ok := tlu.mutation.Price(); ok {
		_spec.SetField(tradelog.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := tlu.mutation.AddedPrice(); ok {
		_spec.AddField(tradelog.FieldPrice, field.TypeFloat64, value)
	}
	if tlu.mutation.PriceCleared() {
		_spec.ClearField(tradelog.FieldPrice, field.TypeFloat64)
	}
	if value, ok := tlu.mutation.Quantity(); ok {
		_spec.SetField(tradelog.FieldQuantity, field.TypeFloat64, value)
	}
	if value, ok := tlu.mutation.AddedQuantity(); ok {
		_spec.AddField(tradelog.FieldQuantity, field.TypeFloat64, value)
	}
	if tlu.mutation.QuantityCleared() {
		_spec.ClearField(tradelog.FieldQuantity, field.TypeFloat64)
	}
	if value, ok := tlu.mutation.Leverage(); ok {
		_spec.SetField(tradelog.FieldLeverage, field.TypeInt, value)
	}
	if value, ok := tlu.mutation.AddedLeverage(); ok {
		_spec.AddField(tradelog.FieldLeverage, field.TypeInt, value)
	}
	if tlu.mutation.LeverageCleared() {
		_spec.ClearField(tradelog.FieldLeverage, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tradelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TradeLogUpdateOne is the builder for updating a single TradeLog entity.
type TradeLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TradeLogMutation
}

// SetDate sets the "date" field.
func (tluo *TradeLogUpdateOne) SetDate(t time.Time) *TradeLogUpdateOne {
	tluo.mutation.SetDate(t)
	return tluo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillableDate(t *time.Time) *TradeLogUpdateOne {
	if t != nil {
		tluo.SetDate(*t)
	}
	return tluo
}

// ClearDate clears the value of the "date" field.
func (tluo *TradeLogUpdateOne) ClearDate() *TradeLogUpdateOne {
	tluo.mutation.ClearDate()
	return tluo
}

// SetExchange sets the "exchange" field.
func (tluo *TradeLogUpdateOne) SetExchange(s string) *TradeLogUpdateOne {
	tluo.mutation.SetExchange(s)
	return tluo
}

// SetNillableExchange sets the "exchange" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillableExchange(s *string) *TradeLogUpdateOne {
	if s != nil {
		tluo.SetExchange(*s)
	}
	return tluo
}

// ClearExchange clears the value of the "exchange" field.
func (tluo *TradeLogUpdateOne) ClearExchange() *TradeLogUpdateOne {
	tluo.mutation.ClearExchange()
	return tluo
}

// SetTicker sets the "ticker" field.
func (tluo *TradeLogUpdateOne) SetTicker(s string) *TradeLogUpdateOne {
	tluo.mutation.SetTicker(s)
	return tluo
}

// SetNillableTicker sets the "ticker" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillableTicker(s *string) *TradeLogUpdateOne {
	if s != nil {
		tluo.SetTicker(*s)
	}
	return tluo
}

// ClearTicker clears the value of the "ticker" field.
func (tluo *TradeLogUpdateOne) ClearTicker() *TradeLogUpdateOne {
	tluo.mutation.ClearTicker()
	return tluo
}

// SetPosition sets the "position" field.
func (tluo *TradeLogUpdateOne) SetPosition(s string) *TradeLogUpdateOne {
	tluo.mutation.SetPosition(s)
	return tluo
}

// SetNillablePosition sets the "position" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillablePosition(s *string) *TradeLogUpdateOne {
	if s != nil {
		tluo.SetPosition(*s)
	}
	return tluo
}

// ClearPosition clears the value of the "position" field.
func (tluo *TradeLogUpdateOne) ClearPosition() *TradeLogUpdateOne {
	tluo.mutation.ClearPosition()
	return tluo
}

// SetStrategy sets the "strategy" field.
func (tluo *TradeLogUpdateOne) SetStrategy(s string) *TradeLogUpdateOne {
	tluo.mutation.SetStrategy(s)
	return tluo
}

// SetNillableStrategy sets the "strategy" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillableStrategy(s *string) *TradeLogUpdateOne {
	if s != nil {
		tluo.SetStrategy(*s)
	}
	return tluo
}

// ClearStrategy clears the value of the "strategy" field.
func (tluo *TradeLogUpdateOne) ClearStrategy() *TradeLogUpdateOne {
	tluo.mutation.ClearStrategy()
	return tluo
}

// SetPrice sets the "price" field.
func (tluo *TradeLogUpdateOne) SetPrice(f float64) *TradeLogUpdateOne {
	tluo.mutation.ResetPrice()
	tluo.mutation.SetPrice(f)
	return tluo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillablePrice(f *float64) *TradeLogUpdateOne {
	if f != nil {
		tluo.SetPrice(*f)
	}
	return tluo
}

// AddPrice adds f to the "price" field.
func (tluo *TradeLogUpdateOne) AddPrice(f float64) *TradeLogUpdateOne {
	tluo.mutation.AddPrice(f)
	return tluo
}

// ClearPrice clears the value of the "price" field.
func (tluo *TradeLogUpdateOne) ClearPrice() *TradeLogUpdateOne {
	tluo.mutation.ClearPrice()
	return tluo
}

// SetQuantity sets the "quantity" field.
func (tluo *TradeLogUpdateOne) SetQuantity(f float64) *TradeLogUpdateOne {
	tluo.mutation.ResetQuantity()
	tluo.mutation.SetQuantity(f)
	return tluo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillableQuantity(f *float64) *TradeLogUpdateOne {
	if f != nil {
		tluo.SetQuantity(*f)
	}
	return tluo
}

// AddQuantity adds f to the "quantity" field.
func (tluo *TradeLogUpdateOne) AddQuantity(f float64) *TradeLogUpdateOne {
	tluo.mutation.AddQuantity(f)
	return tluo
}

// ClearQuantity clears the value of the "quantity" field.
func (tluo *TradeLogUpdateOne) ClearQuantity() *TradeLogUpdateOne {
	tluo.mutation.ClearQuantity()
	return tluo
}

// SetLeverage sets the "leverage" field.
func (tluo *TradeLogUpdateOne) SetLeverage(i int) *TradeLogUpdateOne {
	tluo.mutation.ResetLeverage()
	tluo.mutation.SetLeverage(i)
	return tluo
}

// SetNillableLeverage sets the "leverage" field if the given value is not nil.
func (tluo *TradeLogUpdateOne) SetNillableLeverage(i *int) *TradeLogUpdateOne {
	if i != nil {
		tluo.SetLeverage(*i)
	}
	return tluo
}

// AddLeverage adds i to the "leverage" field.
func (tluo *TradeLogUpdateOne) AddLeverage(i int) *TradeLogUpdateOne {
	tluo.mutation.AddLeverage(i)
	return tluo
}

// ClearLeverage clears the value of the "leverage" field.
func (tluo *TradeLogUpdateOne) ClearLeverage() *TradeLogUpdateOne {
	tluo.mutation.ClearLeverage()
	return tluo
}

// Mutation returns the TradeLogMutation object of the builder.
func (tluo *TradeLogUpdateOne) Mutation() *TradeLogMutation {
	return tluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tluo *TradeLogUpdateOne) Select(field string, fields ...string) *TradeLogUpdateOne {
	tluo.fields = append([]string{field}, fields...)
	return tluo
}

// Save executes the query and returns the updated TradeLog entity.
func (tluo *TradeLogUpdateOne) Save(ctx context.Context) (*TradeLog, error) {
	var (
		err  error
		node *TradeLog
	)
	if len(tluo.hooks) == 0 {
		node, err = tluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TradeLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tluo.mutation = mutation
			node, err = tluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tluo.hooks) - 1; i >= 0; i-- {
			if tluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TradeLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TradeLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tluo *TradeLogUpdateOne) SaveX(ctx context.Context) *TradeLog {
	node, err := tluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tluo *TradeLogUpdateOne) Exec(ctx context.Context) error {
	_, err := tluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tluo *TradeLogUpdateOne) ExecX(ctx context.Context) {
	if err := tluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tluo *TradeLogUpdateOne) sqlSave(ctx context.Context) (_node *TradeLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tradelog.Table,
			Columns: tradelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradelog.FieldID,
			},
		},
	}
	id, ok := tluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TradeLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tradelog.FieldID)
		for _, f := range fields {
			if !tradelog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tradelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tluo.mutation.Date(); ok {
		_spec.SetField(tradelog.FieldDate, field.TypeTime, value)
	}
	if tluo.mutation.DateCleared() {
		_spec.ClearField(tradelog.FieldDate, field.TypeTime)
	}
	if value, ok := tluo.mutation.Exchange(); ok {
		_spec.SetField(tradelog.FieldExchange, field.TypeString, value)
	}
	if tluo.mutation.ExchangeCleared() {
		_spec.ClearField(tradelog.FieldExchange, field.TypeString)
	}
	if value, ok := tluo.mutation.Ticker(); ok {
		_spec.SetField(tradelog.FieldTicker, field.TypeString, value)
	}
	if tluo.mutation.TickerCleared() {
		_spec.ClearField(tradelog.FieldTicker, field.TypeString)
	}
	if value, ok := tluo.mutation.Position(); ok {
		_spec.SetField(tradelog.FieldPosition, field.TypeString, value)
	}
	if tluo.mutation.PositionCleared() {
		_spec.ClearField(tradelog.FieldPosition, field.TypeString)
	}
	if value, ok := tluo.mutation.Strategy(); ok {
		_spec.SetField(tradelog.FieldStrategy, field.TypeString, value)
	}
	if tluo.mutation.StrategyCleared() {
		_spec.ClearField(tradelog.FieldStrategy, field.TypeString)
	}
	if value, ok := tluo.mutation.Price(); ok {
		_spec.SetField(tradelog.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := tluo.mutation.AddedPrice(); ok {
		_spec.AddField(tradelog.FieldPrice, field.TypeFloat64, value)
	}
	if tluo.mutation.PriceCleared() {
		_spec.ClearField(tradelog.FieldPrice, field.TypeFloat64)
	}
	if value, ok := tluo.mutation.Quantity(); ok {
		_spec.SetField(tradelog.FieldQuantity, field.TypeFloat64, value)
	}
	if value, ok := tluo.mutation.AddedQuantity(); ok {
		_spec.AddField(tradelog.FieldQuantity, field.TypeFloat64, value)
	}
	if tluo.mutation.QuantityCleared() {
		_spec.ClearField(tradelog.FieldQuantity, field.TypeFloat64)
	}
	if value, ok := tluo.mutation.Leverage(); ok {
		_spec.SetField(tradelog.FieldLeverage, field.TypeInt, value)
	}
	if value, ok := tluo.mutation.AddedLeverage(); ok {
		_spec.AddField(tradelog.FieldLeverage, field.TypeInt, value)
	}
	if tluo.mutation.LeverageCleared() {
		_spec.ClearField(tradelog.FieldLeverage, field.TypeInt)
	}
	_node = &TradeLog{config: tluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tradelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}