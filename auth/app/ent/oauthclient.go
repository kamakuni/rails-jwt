// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/oauthclient"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// OAuthClient is the model entity for the OAuthClient schema.
type OAuthClient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// ClientType holds the value of the "client_type" field.
	ClientType string `json:"client_type,omitempty"`
	// ClientName holds the value of the "client_name" field.
	ClientName string `json:"client_name,omitempty"`
	// RedirectURI holds the value of the "redirect_uri" field.
	RedirectURI string `json:"redirect_uri,omitempty"`
	// Scope holds the value of the "scope" field.
	Scope string `json:"scope,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuthClient) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthclient.FieldID:
			values[i] = new(sql.NullInt64)
		case oauthclient.FieldClientID, oauthclient.FieldClientType, oauthclient.FieldClientName, oauthclient.FieldRedirectURI, oauthclient.FieldScope:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OAuthClient", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuthClient fields.
func (oc *OAuthClient) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthclient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oc.ID = int(value.Int64)
		case oauthclient.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				oc.ClientID = value.String
			}
		case oauthclient.FieldClientType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_type", values[i])
			} else if value.Valid {
				oc.ClientType = value.String
			}
		case oauthclient.FieldClientName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_name", values[i])
			} else if value.Valid {
				oc.ClientName = value.String
			}
		case oauthclient.FieldRedirectURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_uri", values[i])
			} else if value.Valid {
				oc.RedirectURI = value.String
			}
		case oauthclient.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				oc.Scope = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OAuthClient.
// Note that you need to call OAuthClient.Unwrap() before calling this method if this OAuthClient
// was returned from a transaction, and the transaction was committed or rolled back.
func (oc *OAuthClient) Update() *OAuthClientUpdateOne {
	return (&OAuthClientClient{config: oc.config}).UpdateOne(oc)
}

// Unwrap unwraps the OAuthClient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oc *OAuthClient) Unwrap() *OAuthClient {
	_tx, ok := oc.config.driver.(*txDriver)
	if !ok {
		panic("ent: OAuthClient is not a transactional entity")
	}
	oc.config.driver = _tx.drv
	return oc
}

// String implements the fmt.Stringer.
func (oc *OAuthClient) String() string {
	var builder strings.Builder
	builder.WriteString("OAuthClient(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oc.ID))
	builder.WriteString("client_id=")
	builder.WriteString(oc.ClientID)
	builder.WriteString(", ")
	builder.WriteString("client_type=")
	builder.WriteString(oc.ClientType)
	builder.WriteString(", ")
	builder.WriteString("client_name=")
	builder.WriteString(oc.ClientName)
	builder.WriteString(", ")
	builder.WriteString("redirect_uri=")
	builder.WriteString(oc.RedirectURI)
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(oc.Scope)
	builder.WriteByte(')')
	return builder.String()
}

// OAuthClients is a parsable slice of OAuthClient.
type OAuthClients []*OAuthClient

func (oc OAuthClients) config(cfg config) {
	for _i := range oc {
		oc[_i].config = cfg
	}
}
