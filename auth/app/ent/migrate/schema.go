// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthorizationCodesColumns holds the columns for the "authorization_codes" table.
	AuthorizationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "client_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "scopes", Type: field.TypeString},
	}
	// AuthorizationCodesTable holds the schema information for the "authorization_codes" table.
	AuthorizationCodesTable = &schema.Table{
		Name:       "authorization_codes",
		Columns:    AuthorizationCodesColumns,
		PrimaryKey: []*schema.Column{AuthorizationCodesColumns[0]},
	}
	// OauthClientsColumns holds the columns for the "oauth_clients" table.
	OauthClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "client_id", Type: field.TypeString},
		{Name: "client_type", Type: field.TypeString},
		{Name: "client_name", Type: field.TypeString},
		{Name: "redirect_uri", Type: field.TypeString},
		{Name: "scope", Type: field.TypeString},
	}
	// OauthClientsTable holds the schema information for the "oauth_clients" table.
	OauthClientsTable = &schema.Table{
		Name:       "oauth_clients",
		Columns:    OauthClientsColumns,
		PrimaryKey: []*schema.Column{OauthClientsColumns[0]},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString},
		{Name: "expired", Type: field.TypeBool, Default: false},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthorizationCodesTable,
		OauthClientsTable,
		RefreshTokensTable,
		UsersTable,
	}
)

func init() {
}
