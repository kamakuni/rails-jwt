package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OAuthClient holds the schema definition for the OAuthClient entity.
type OAuthClient struct {
	ent.Schema
}

// Fields of the OAuthClient.
func (OAuthClient) Fields() []ent.Field {
	return []ent.Field{
		field.String("client_id").NotEmpty(),
		field.String("client_type").NotEmpty(),
		field.String("client_name").NotEmpty(),
		field.String("redirect_uri").NotEmpty(),
		field.String("scope").NotEmpty(),
	}
}

// Edges of the OAuthClient.
func (OAuthClient) Edges() []ent.Edge {
	return nil
}
