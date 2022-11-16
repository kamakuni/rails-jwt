// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/authorizationcode"
	"auth/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthorizationCodeUpdate is the builder for updating AuthorizationCode entities.
type AuthorizationCodeUpdate struct {
	config
	hooks    []Hook
	mutation *AuthorizationCodeMutation
}

// Where appends a list predicates to the AuthorizationCodeUpdate builder.
func (acu *AuthorizationCodeUpdate) Where(ps ...predicate.AuthorizationCode) *AuthorizationCodeUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetClientID sets the "client_id" field.
func (acu *AuthorizationCodeUpdate) SetClientID(s string) *AuthorizationCodeUpdate {
	acu.mutation.SetClientID(s)
	return acu
}

// SetUserID sets the "user_id" field.
func (acu *AuthorizationCodeUpdate) SetUserID(s string) *AuthorizationCodeUpdate {
	acu.mutation.SetUserID(s)
	return acu
}

// SetScopes sets the "scopes" field.
func (acu *AuthorizationCodeUpdate) SetScopes(s string) *AuthorizationCodeUpdate {
	acu.mutation.SetScopes(s)
	return acu
}

// Mutation returns the AuthorizationCodeMutation object of the builder.
func (acu *AuthorizationCodeUpdate) Mutation() *AuthorizationCodeMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AuthorizationCodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(acu.hooks) == 0 {
		if err = acu.check(); err != nil {
			return 0, err
		}
		affected, err = acu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorizationCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acu.check(); err != nil {
				return 0, err
			}
			acu.mutation = mutation
			affected, err = acu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acu.hooks) - 1; i >= 0; i-- {
			if acu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AuthorizationCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AuthorizationCodeUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AuthorizationCodeUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acu *AuthorizationCodeUpdate) check() error {
	if v, ok := acu.mutation.ClientID(); ok {
		if err := authorizationcode.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`ent: validator failed for field "AuthorizationCode.client_id": %w`, err)}
		}
	}
	if v, ok := acu.mutation.UserID(); ok {
		if err := authorizationcode.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "AuthorizationCode.user_id": %w`, err)}
		}
	}
	if v, ok := acu.mutation.Scopes(); ok {
		if err := authorizationcode.ScopesValidator(v); err != nil {
			return &ValidationError{Name: "scopes", err: fmt.Errorf(`ent: validator failed for field "AuthorizationCode.scopes": %w`, err)}
		}
	}
	return nil
}

func (acu *AuthorizationCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authorizationcode.Table,
			Columns: authorizationcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authorizationcode.FieldID,
			},
		},
	}
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.ClientID(); ok {
		_spec.SetField(authorizationcode.FieldClientID, field.TypeString, value)
	}
	if value, ok := acu.mutation.UserID(); ok {
		_spec.SetField(authorizationcode.FieldUserID, field.TypeString, value)
	}
	if value, ok := acu.mutation.Scopes(); ok {
		_spec.SetField(authorizationcode.FieldScopes, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authorizationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AuthorizationCodeUpdateOne is the builder for updating a single AuthorizationCode entity.
type AuthorizationCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthorizationCodeMutation
}

// SetClientID sets the "client_id" field.
func (acuo *AuthorizationCodeUpdateOne) SetClientID(s string) *AuthorizationCodeUpdateOne {
	acuo.mutation.SetClientID(s)
	return acuo
}

// SetUserID sets the "user_id" field.
func (acuo *AuthorizationCodeUpdateOne) SetUserID(s string) *AuthorizationCodeUpdateOne {
	acuo.mutation.SetUserID(s)
	return acuo
}

// SetScopes sets the "scopes" field.
func (acuo *AuthorizationCodeUpdateOne) SetScopes(s string) *AuthorizationCodeUpdateOne {
	acuo.mutation.SetScopes(s)
	return acuo
}

// Mutation returns the AuthorizationCodeMutation object of the builder.
func (acuo *AuthorizationCodeUpdateOne) Mutation() *AuthorizationCodeMutation {
	return acuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AuthorizationCodeUpdateOne) Select(field string, fields ...string) *AuthorizationCodeUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AuthorizationCode entity.
func (acuo *AuthorizationCodeUpdateOne) Save(ctx context.Context) (*AuthorizationCode, error) {
	var (
		err  error
		node *AuthorizationCode
	)
	if len(acuo.hooks) == 0 {
		if err = acuo.check(); err != nil {
			return nil, err
		}
		node, err = acuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorizationCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acuo.check(); err != nil {
				return nil, err
			}
			acuo.mutation = mutation
			node, err = acuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acuo.hooks) - 1; i >= 0; i-- {
			if acuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AuthorizationCode)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AuthorizationCodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AuthorizationCodeUpdateOne) SaveX(ctx context.Context) *AuthorizationCode {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AuthorizationCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AuthorizationCodeUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acuo *AuthorizationCodeUpdateOne) check() error {
	if v, ok := acuo.mutation.ClientID(); ok {
		if err := authorizationcode.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`ent: validator failed for field "AuthorizationCode.client_id": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.UserID(); ok {
		if err := authorizationcode.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "AuthorizationCode.user_id": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.Scopes(); ok {
		if err := authorizationcode.ScopesValidator(v); err != nil {
			return &ValidationError{Name: "scopes", err: fmt.Errorf(`ent: validator failed for field "AuthorizationCode.scopes": %w`, err)}
		}
	}
	return nil
}

func (acuo *AuthorizationCodeUpdateOne) sqlSave(ctx context.Context) (_node *AuthorizationCode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authorizationcode.Table,
			Columns: authorizationcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authorizationcode.FieldID,
			},
		},
	}
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthorizationCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authorizationcode.FieldID)
		for _, f := range fields {
			if !authorizationcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authorizationcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.ClientID(); ok {
		_spec.SetField(authorizationcode.FieldClientID, field.TypeString, value)
	}
	if value, ok := acuo.mutation.UserID(); ok {
		_spec.SetField(authorizationcode.FieldUserID, field.TypeString, value)
	}
	if value, ok := acuo.mutation.Scopes(); ok {
		_spec.SetField(authorizationcode.FieldScopes, field.TypeString, value)
	}
	_node = &AuthorizationCode{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authorizationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
