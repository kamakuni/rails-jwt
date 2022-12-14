// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/refreshtoken"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// RefreshToken is the model entity for the RefreshToken schema.
type RefreshToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Expired holds the value of the "expired" field.
	Expired bool `json:"expired,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RefreshToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldExpired:
			values[i] = new(sql.NullBool)
		case refreshtoken.FieldID:
			values[i] = new(sql.NullInt64)
		case refreshtoken.FieldToken:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RefreshToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RefreshToken fields.
func (rt *RefreshToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rt.ID = int(value.Int64)
		case refreshtoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				rt.Token = value.String
			}
		case refreshtoken.FieldExpired:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field expired", values[i])
			} else if value.Valid {
				rt.Expired = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RefreshToken.
// Note that you need to call RefreshToken.Unwrap() before calling this method if this RefreshToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RefreshToken) Update() *RefreshTokenUpdateOne {
	return (&RefreshTokenClient{config: rt.config}).UpdateOne(rt)
}

// Unwrap unwraps the RefreshToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rt *RefreshToken) Unwrap() *RefreshToken {
	_tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: RefreshToken is not a transactional entity")
	}
	rt.config.driver = _tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RefreshToken) String() string {
	var builder strings.Builder
	builder.WriteString("RefreshToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rt.ID))
	builder.WriteString("token=")
	builder.WriteString(rt.Token)
	builder.WriteString(", ")
	builder.WriteString("expired=")
	builder.WriteString(fmt.Sprintf("%v", rt.Expired))
	builder.WriteByte(')')
	return builder.String()
}

// RefreshTokens is a parsable slice of RefreshToken.
type RefreshTokens []*RefreshToken

func (rt RefreshTokens) config(cfg config) {
	for _i := range rt {
		rt[_i].config = cfg
	}
}
