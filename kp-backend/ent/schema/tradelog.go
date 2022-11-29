package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TradeLog holds the schema definition for the TradeLog entity.
type TradeLog struct {
	ent.Schema
}

// Fields of the TradeLog.
func (TradeLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Time("date").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}).
			Optional().
			Comment("Date when transaction happened"),
		field.String("exchange").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).
			Optional().
			Comment("Exchange where transaction happened"),
		field.String("ticker").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).
			Optional().
			Comment("Asset ticker ex) BTC"),
		field.String("position").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).
			Optional().
			Comment("Long position or Short"),
		field.String("strategy").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).
			Optional().
			Comment("Strategy name"),
		field.Float("price").
			Optional(),
		field.Float("quantity").
			Optional(),
		field.Int("leverage").
			Optional(),
	}
}

// Edges of the TradeLog.
func (TradeLog) Edges() []ent.Edge {
	return nil
}

func (TradeLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tradelog"},
	}
}
