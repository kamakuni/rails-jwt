// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/authorizationcode"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AuthorizationCode is the model entity for the AuthorizationCode schema.
type AuthorizationCode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Scopes holds the value of the "scopes" field.
	Scopes string `json:"scopes,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthorizationCode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authorizationcode.FieldID:
			values[i] = new(sql.NullInt64)
		case authorizationcode.FieldClientID, authorizationcode.FieldUserID, authorizationcode.FieldScopes:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuthorizationCode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthorizationCode fields.
func (ac *AuthorizationCode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authorizationcode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		case authorizationcode.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				ac.ClientID = value.String
			}
		case authorizationcode.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ac.UserID = value.String
			}
		case authorizationcode.FieldScopes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value.Valid {
				ac.Scopes = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AuthorizationCode.
// Note that you need to call AuthorizationCode.Unwrap() before calling this method if this AuthorizationCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AuthorizationCode) Update() *AuthorizationCodeUpdateOne {
	return (&AuthorizationCodeClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the AuthorizationCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AuthorizationCode) Unwrap() *AuthorizationCode {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthorizationCode is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AuthorizationCode) String() string {
	var builder strings.Builder
	builder.WriteString("AuthorizationCode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("client_id=")
	builder.WriteString(ac.ClientID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(ac.UserID)
	builder.WriteString(", ")
	builder.WriteString("scopes=")
	builder.WriteString(ac.Scopes)
	builder.WriteByte(')')
	return builder.String()
}

// AuthorizationCodes is a parsable slice of AuthorizationCode.
type AuthorizationCodes []*AuthorizationCode

func (ac AuthorizationCodes) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}