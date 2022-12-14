// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/refreshtoken"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RefreshTokenCreate is the builder for creating a RefreshToken entity.
type RefreshTokenCreate struct {
	config
	mutation *RefreshTokenMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (rtc *RefreshTokenCreate) SetToken(s string) *RefreshTokenCreate {
	rtc.mutation.SetToken(s)
	return rtc
}

// SetExpired sets the "expired" field.
func (rtc *RefreshTokenCreate) SetExpired(b bool) *RefreshTokenCreate {
	rtc.mutation.SetExpired(b)
	return rtc
}

// SetNillableExpired sets the "expired" field if the given value is not nil.
func (rtc *RefreshTokenCreate) SetNillableExpired(b *bool) *RefreshTokenCreate {
	if b != nil {
		rtc.SetExpired(*b)
	}
	return rtc
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtc *RefreshTokenCreate) Mutation() *RefreshTokenMutation {
	return rtc.mutation
}

// Save creates the RefreshToken in the database.
func (rtc *RefreshTokenCreate) Save(ctx context.Context) (*RefreshToken, error) {
	var (
		err  error
		node *RefreshToken
	)
	rtc.defaults()
	if len(rtc.hooks) == 0 {
		if err = rtc.check(); err != nil {
			return nil, err
		}
		node, err = rtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefreshTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtc.check(); err != nil {
				return nil, err
			}
			rtc.mutation = mutation
			if node, err = rtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rtc.hooks) - 1; i >= 0; i-- {
			if rtc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rtc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RefreshToken)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RefreshTokenMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rtc *RefreshTokenCreate) SaveX(ctx context.Context) *RefreshToken {
	v, err := rtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rtc *RefreshTokenCreate) Exec(ctx context.Context) error {
	_, err := rtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtc *RefreshTokenCreate) ExecX(ctx context.Context) {
	if err := rtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtc *RefreshTokenCreate) defaults() {
	if _, ok := rtc.mutation.Expired(); !ok {
		v := refreshtoken.DefaultExpired
		rtc.mutation.SetExpired(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtc *RefreshTokenCreate) check() error {
	if _, ok := rtc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "RefreshToken.token"`)}
	}
	if v, ok := rtc.mutation.Token(); ok {
		if err := refreshtoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "RefreshToken.token": %w`, err)}
		}
	}
	if _, ok := rtc.mutation.Expired(); !ok {
		return &ValidationError{Name: "expired", err: errors.New(`ent: missing required field "RefreshToken.expired"`)}
	}
	return nil
}

func (rtc *RefreshTokenCreate) sqlSave(ctx context.Context) (*RefreshToken, error) {
	_node, _spec := rtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rtc *RefreshTokenCreate) createSpec() (*RefreshToken, *sqlgraph.CreateSpec) {
	var (
		_node = &RefreshToken{config: rtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: refreshtoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: refreshtoken.FieldID,
			},
		}
	)
	if value, ok := rtc.mutation.Token(); ok {
		_spec.SetField(refreshtoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := rtc.mutation.Expired(); ok {
		_spec.SetField(refreshtoken.FieldExpired, field.TypeBool, value)
		_node.Expired = value
	}
	return _node, _spec
}

// RefreshTokenCreateBulk is the builder for creating many RefreshToken entities in bulk.
type RefreshTokenCreateBulk struct {
	config
	builders []*RefreshTokenCreate
}

// Save creates the RefreshToken entities in the database.
func (rtcb *RefreshTokenCreateBulk) Save(ctx context.Context) ([]*RefreshToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rtcb.builders))
	nodes := make([]*RefreshToken, len(rtcb.builders))
	mutators := make([]Mutator, len(rtcb.builders))
	for i := range rtcb.builders {
		func(i int, root context.Context) {
			builder := rtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RefreshTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rtcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rtcb *RefreshTokenCreateBulk) SaveX(ctx context.Context) []*RefreshToken {
	v, err := rtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rtcb *RefreshTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := rtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtcb *RefreshTokenCreateBulk) ExecX(ctx context.Context) {
	if err := rtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
