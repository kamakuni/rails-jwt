package schema

import (
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
		field.String("user_id").NotEmpty(),
		field.String("scopes").NotEmpty(),
	}
}

// Edges of the AuthorizationCode.
func (AuthorizationCode) Edges() []ent.Edge {
	return nil
}
