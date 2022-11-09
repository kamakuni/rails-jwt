package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RefreshToken holds the schema definition for the RefreshToken entity.
type RefreshToken struct {
	ent.Schema
}

// Fields of the RefreshToken.
func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").NotEmpty(),
		field.Bool("expired").Default(false),
	}
}

// Edges of the RefreshToken.
func (RefreshToken) Edges() []ent.Edge {
	return nil
}
