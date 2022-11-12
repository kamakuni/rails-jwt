package schema

import "entgo.io/ent"

// AuthorizationCode holds the schema definition for the AuthorizationCode entity.
type AuthorizationCode struct {
	ent.Schema
}

// Fields of the AuthorizationCode.
func (AuthorizationCode) Fields() []ent.Field {
	return nil
}

// Edges of the AuthorizationCode.
func (AuthorizationCode) Edges() []ent.Edge {
	return nil
}
