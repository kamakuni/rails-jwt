package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthorizationCode holds the schema definition for the AuthorizationCode entity.
type AuthorizationCode struct {
	ent.Schema
}

// Fields of the AuthorizationCode.
func (AuthorizationCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("client_id").NotEmpty(),
		field.String("code").NotEmpty(),
		field.Time("issued").Immutable().Default(time.Now()),
	}
}

// Edges of the AuthorizationCode.
func (AuthorizationCode) Edges() []ent.Edge {
	return nil
}
