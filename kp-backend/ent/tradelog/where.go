// Code generated by ent, DO NOT EDIT.

package tradelog

import (
	"kimchi/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// Exchange applies equality check predicate on the "exchange" field. It's identical to ExchangeEQ.
func Exchange(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExchange), v))
	})
}

// Ticker applies equality check predicate on the "ticker" field. It's identical to TickerEQ.
func Ticker(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTicker), v))
	})
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// Strategy applies equality check predicate on the "strategy" field. It's identical to StrategyEQ.
func Strategy(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrategy), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// Leverage applies equality check predicate on the "leverage" field. It's identical to LeverageEQ.
func Leverage(v int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeverage), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// DateIsNil applies the IsNil predicate on the "date" field.
func DateIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDate)))
	})
}

// DateNotNil applies the NotNil predicate on the "date" field.
func DateNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDate)))
	})
}

// ExchangeEQ applies the EQ predicate on the "exchange" field.
func ExchangeEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExchange), v))
	})
}

// ExchangeNEQ applies the NEQ predicate on the "exchange" field.
func ExchangeNEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExchange), v))
	})
}

// ExchangeIn applies the In predicate on the "exchange" field.
func ExchangeIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExchange), v...))
	})
}

// ExchangeNotIn applies the NotIn predicate on the "exchange" field.
func ExchangeNotIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExchange), v...))
	})
}

// ExchangeGT applies the GT predicate on the "exchange" field.
func ExchangeGT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExchange), v))
	})
}

// ExchangeGTE applies the GTE predicate on the "exchange" field.
func ExchangeGTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExchange), v))
	})
}

// ExchangeLT applies the LT predicate on the "exchange" field.
func ExchangeLT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExchange), v))
	})
}

// ExchangeLTE applies the LTE predicate on the "exchange" field.
func ExchangeLTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExchange), v))
	})
}

// ExchangeContains applies the Contains predicate on the "exchange" field.
func ExchangeContains(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExchange), v))
	})
}

// ExchangeHasPrefix applies the HasPrefix predicate on the "exchange" field.
func ExchangeHasPrefix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExchange), v))
	})
}

// ExchangeHasSuffix applies the HasSuffix predicate on the "exchange" field.
func ExchangeHasSuffix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExchange), v))
	})
}

// ExchangeIsNil applies the IsNil predicate on the "exchange" field.
func ExchangeIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExchange)))
	})
}

// ExchangeNotNil applies the NotNil predicate on the "exchange" field.
func ExchangeNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExchange)))
	})
}

// ExchangeEqualFold applies the EqualFold predicate on the "exchange" field.
func ExchangeEqualFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExchange), v))
	})
}

// ExchangeContainsFold applies the ContainsFold predicate on the "exchange" field.
func ExchangeContainsFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExchange), v))
	})
}

// TickerEQ applies the EQ predicate on the "ticker" field.
func TickerEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTicker), v))
	})
}

// TickerNEQ applies the NEQ predicate on the "ticker" field.
func TickerNEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTicker), v))
	})
}

// TickerIn applies the In predicate on the "ticker" field.
func TickerIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTicker), v...))
	})
}

// TickerNotIn applies the NotIn predicate on the "ticker" field.
func TickerNotIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTicker), v...))
	})
}

// TickerGT applies the GT predicate on the "ticker" field.
func TickerGT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTicker), v))
	})
}

// TickerGTE applies the GTE predicate on the "ticker" field.
func TickerGTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTicker), v))
	})
}

// TickerLT applies the LT predicate on the "ticker" field.
func TickerLT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTicker), v))
	})
}

// TickerLTE applies the LTE predicate on the "ticker" field.
func TickerLTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTicker), v))
	})
}

// TickerContains applies the Contains predicate on the "ticker" field.
func TickerContains(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTicker), v))
	})
}

// TickerHasPrefix applies the HasPrefix predicate on the "ticker" field.
func TickerHasPrefix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTicker), v))
	})
}

// TickerHasSuffix applies the HasSuffix predicate on the "ticker" field.
func TickerHasSuffix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTicker), v))
	})
}

// TickerIsNil applies the IsNil predicate on the "ticker" field.
func TickerIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTicker)))
	})
}

// TickerNotNil applies the NotNil predicate on the "ticker" field.
func TickerNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTicker)))
	})
}

// TickerEqualFold applies the EqualFold predicate on the "ticker" field.
func TickerEqualFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTicker), v))
	})
}

// TickerContainsFold applies the ContainsFold predicate on the "ticker" field.
func TickerContainsFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTicker), v))
	})
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPosition), v))
	})
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPosition), v...))
	})
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPosition), v...))
	})
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPosition), v))
	})
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPosition), v))
	})
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPosition), v))
	})
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPosition), v))
	})
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPosition), v))
	})
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPosition), v))
	})
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPosition), v))
	})
}

// PositionIsNil applies the IsNil predicate on the "position" field.
func PositionIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPosition)))
	})
}

// PositionNotNil applies the NotNil predicate on the "position" field.
func PositionNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPosition)))
	})
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPosition), v))
	})
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPosition), v))
	})
}

// StrategyEQ applies the EQ predicate on the "strategy" field.
func StrategyEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrategy), v))
	})
}

// StrategyNEQ applies the NEQ predicate on the "strategy" field.
func StrategyNEQ(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStrategy), v))
	})
}

// StrategyIn applies the In predicate on the "strategy" field.
func StrategyIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStrategy), v...))
	})
}

// StrategyNotIn applies the NotIn predicate on the "strategy" field.
func StrategyNotIn(vs ...string) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStrategy), v...))
	})
}

// StrategyGT applies the GT predicate on the "strategy" field.
func StrategyGT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStrategy), v))
	})
}

// StrategyGTE applies the GTE predicate on the "strategy" field.
func StrategyGTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStrategy), v))
	})
}

// StrategyLT applies the LT predicate on the "strategy" field.
func StrategyLT(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStrategy), v))
	})
}

// StrategyLTE applies the LTE predicate on the "strategy" field.
func StrategyLTE(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStrategy), v))
	})
}

// StrategyContains applies the Contains predicate on the "strategy" field.
func StrategyContains(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStrategy), v))
	})
}

// StrategyHasPrefix applies the HasPrefix predicate on the "strategy" field.
func StrategyHasPrefix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStrategy), v))
	})
}

// StrategyHasSuffix applies the HasSuffix predicate on the "strategy" field.
func StrategyHasSuffix(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStrategy), v))
	})
}

// StrategyIsNil applies the IsNil predicate on the "strategy" field.
func StrategyIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStrategy)))
	})
}

// StrategyNotNil applies the NotNil predicate on the "strategy" field.
func StrategyNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStrategy)))
	})
}

// StrategyEqualFold applies the EqualFold predicate on the "strategy" field.
func StrategyEqualFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStrategy), v))
	})
}

// StrategyContainsFold applies the ContainsFold predicate on the "strategy" field.
func StrategyContainsFold(v string) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStrategy), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// PriceIsNil applies the IsNil predicate on the "price" field.
func PriceIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrice)))
	})
}

// PriceNotNil applies the NotNil predicate on the "price" field.
func PriceNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrice)))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...float64) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldQuantity), v...))
	})
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...float64) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldQuantity), v...))
	})
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v float64) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// QuantityIsNil applies the IsNil predicate on the "quantity" field.
func QuantityIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldQuantity)))
	})
}

// QuantityNotNil applies the NotNil predicate on the "quantity" field.
func QuantityNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldQuantity)))
	})
}

// LeverageEQ applies the EQ predicate on the "leverage" field.
func LeverageEQ(v int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeverage), v))
	})
}

// LeverageNEQ applies the NEQ predicate on the "leverage" field.
func LeverageNEQ(v int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLeverage), v))
	})
}

// LeverageIn applies the In predicate on the "leverage" field.
func LeverageIn(vs ...int) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLeverage), v...))
	})
}

// LeverageNotIn applies the NotIn predicate on the "leverage" field.
func LeverageNotIn(vs ...int) predicate.TradeLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLeverage), v...))
	})
}

// LeverageGT applies the GT predicate on the "leverage" field.
func LeverageGT(v int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLeverage), v))
	})
}

// LeverageGTE applies the GTE predicate on the "leverage" field.
func LeverageGTE(v int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLeverage), v))
	})
}

// LeverageLT applies the LT predicate on the "leverage" field.
func LeverageLT(v int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLeverage), v))
	})
}

// LeverageLTE applies the LTE predicate on the "leverage" field.
func LeverageLTE(v int) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLeverage), v))
	})
}

// LeverageIsNil applies the IsNil predicate on the "leverage" field.
func LeverageIsNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLeverage)))
	})
}

// LeverageNotNil applies the NotNil predicate on the "leverage" field.
func LeverageNotNil() predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLeverage)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TradeLog) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TradeLog) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TradeLog) predicate.TradeLog {
	return predicate.TradeLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
