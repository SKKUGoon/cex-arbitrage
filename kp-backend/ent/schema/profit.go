package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Profit holds the schema definition for the Profit entity.
type Profit struct {
	ent.Schema
}

// Fields of the Profit.
func (Profit) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Time("date").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}).
			Optional().
			Comment("Date when trade happened"),
		field.Floats("profit").
			Optional(),
	}
}

// Edges of the Profit.
func (Profit) Edges() []ent.Edge {
	return nil
}

func (Profit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "profit"},
	}
}
