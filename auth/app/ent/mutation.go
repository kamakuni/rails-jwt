// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/kamakuni/rails-jwt/auth/app/ent/oauthclient"
	"github.com/kamakuni/rails-jwt/auth/app/ent/predicate"
	"github.com/kamakuni/rails-jwt/auth/app/ent/refreshtoken"
	"github.com/kamakuni/rails-jwt/auth/app/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOAuthClient  = "OAuthClient"
	TypeRefreshToken = "RefreshToken"
	TypeUser         = "User"
)

// OAuthClientMutation represents an operation that mutates the OAuthClient nodes in the graph.
type OAuthClientMutation struct {
	config
	op            Op
	typ           string
	id            *int
	client_secret *string
	url           *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*OAuthClient, error)
	predicates    []predicate.OAuthClient
}

var _ ent.Mutation = (*OAuthClientMutation)(nil)

// oauthclientOption allows management of the mutation configuration using functional options.
type oauthclientOption func(*OAuthClientMutation)

// newOAuthClientMutation creates new mutation for the OAuthClient entity.
func newOAuthClientMutation(c config, op Op, opts ...oauthclientOption) *OAuthClientMutation {
	m := &OAuthClientMutation{
		config:        c,
		op:            op,
		typ:           TypeOAuthClient,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOAuthClientID sets the ID field of the mutation.
func withOAuthClientID(id int) oauthclientOption {
	return func(m *OAuthClientMutation) {
		var (
			err   error
			once  sync.Once
			value *OAuthClient
		)
		m.oldValue = func(ctx context.Context) (*OAuthClient, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().OAuthClient.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOAuthClient sets the old OAuthClient of the mutation.
func withOAuthClient(node *OAuthClient) oauthclientOption {
	return func(m *OAuthClientMutation) {
		m.oldValue = func(context.Context) (*OAuthClient, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OAuthClientMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OAuthClientMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OAuthClientMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *OAuthClientMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().OAuthClient.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetClientSecret sets the "client_secret" field.
func (m *OAuthClientMutation) SetClientSecret(s string) {
	m.client_secret = &s
}

// ClientSecret returns the value of the "client_secret" field in the mutation.
func (m *OAuthClientMutation) ClientSecret() (r string, exists bool) {
	v := m.client_secret
	if v == nil {
		return
	}
	return *v, true
}

// OldClientSecret returns the old "client_secret" field's value of the OAuthClient entity.
// If the OAuthClient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OAuthClientMutation) OldClientSecret(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldClientSecret is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldClientSecret requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldClientSecret: %w", err)
	}
	return oldValue.ClientSecret, nil
}

// ResetClientSecret resets all changes to the "client_secret" field.
func (m *OAuthClientMutation) ResetClientSecret() {
	m.client_secret = nil
}

// SetURL sets the "url" field.
func (m *OAuthClientMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *OAuthClientMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the OAuthClient entity.
// If the OAuthClient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OAuthClientMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *OAuthClientMutation) ResetURL() {
	m.url = nil
}

// Where appends a list predicates to the OAuthClientMutation builder.
func (m *OAuthClientMutation) Where(ps ...predicate.OAuthClient) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *OAuthClientMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (OAuthClient).
func (m *OAuthClientMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OAuthClientMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.client_secret != nil {
		fields = append(fields, oauthclient.FieldClientSecret)
	}
	if m.url != nil {
		fields = append(fields, oauthclient.FieldURL)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OAuthClientMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case oauthclient.FieldClientSecret:
		return m.ClientSecret()
	case oauthclient.FieldURL:
		return m.URL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OAuthClientMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case oauthclient.FieldClientSecret:
		return m.OldClientSecret(ctx)
	case oauthclient.FieldURL:
		return m.OldURL(ctx)
	}
	return nil, fmt.Errorf("unknown OAuthClient field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OAuthClientMutation) SetField(name string, value ent.Value) error {
	switch name {
	case oauthclient.FieldClientSecret:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetClientSecret(v)
		return nil
	case oauthclient.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	}
	return fmt.Errorf("unknown OAuthClient field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OAuthClientMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OAuthClientMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OAuthClientMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown OAuthClient numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OAuthClientMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OAuthClientMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OAuthClientMutation) ClearField(name string) error {
	return fmt.Errorf("unknown OAuthClient nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OAuthClientMutation) ResetField(name string) error {
	switch name {
	case oauthclient.FieldClientSecret:
		m.ResetClientSecret()
		return nil
	case oauthclient.FieldURL:
		m.ResetURL()
		return nil
	}
	return fmt.Errorf("unknown OAuthClient field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OAuthClientMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OAuthClientMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OAuthClientMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OAuthClientMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OAuthClientMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OAuthClientMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OAuthClientMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown OAuthClient unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OAuthClientMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown OAuthClient edge %s", name)
}

// RefreshTokenMutation represents an operation that mutates the RefreshToken nodes in the graph.
type RefreshTokenMutation struct {
	config
	op            Op
	typ           string
	id            *int
	token         *string
	expired       *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*RefreshToken, error)
	predicates    []predicate.RefreshToken
}

var _ ent.Mutation = (*RefreshTokenMutation)(nil)

// refreshtokenOption allows management of the mutation configuration using functional options.
type refreshtokenOption func(*RefreshTokenMutation)

// newRefreshTokenMutation creates new mutation for the RefreshToken entity.
func newRefreshTokenMutation(c config, op Op, opts ...refreshtokenOption) *RefreshTokenMutation {
	m := &RefreshTokenMutation{
		config:        c,
		op:            op,
		typ:           TypeRefreshToken,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRefreshTokenID sets the ID field of the mutation.
func withRefreshTokenID(id int) refreshtokenOption {
	return func(m *RefreshTokenMutation) {
		var (
			err   error
			once  sync.Once
			value *RefreshToken
		)
		m.oldValue = func(ctx context.Context) (*RefreshToken, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().RefreshToken.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRefreshToken sets the old RefreshToken of the mutation.
func withRefreshToken(node *RefreshToken) refreshtokenOption {
	return func(m *RefreshTokenMutation) {
		m.oldValue = func(context.Context) (*RefreshToken, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RefreshTokenMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RefreshTokenMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RefreshTokenMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RefreshTokenMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().RefreshToken.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetToken sets the "token" field.
func (m *RefreshTokenMutation) SetToken(s string) {
	m.token = &s
}

// Token returns the value of the "token" field in the mutation.
func (m *RefreshTokenMutation) Token() (r string, exists bool) {
	v := m.token
	if v == nil {
		return
	}
	return *v, true
}

// OldToken returns the old "token" field's value of the RefreshToken entity.
// If the RefreshToken object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RefreshTokenMutation) OldToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldToken: %w", err)
	}
	return oldValue.Token, nil
}

// ResetToken resets all changes to the "token" field.
func (m *RefreshTokenMutation) ResetToken() {
	m.token = nil
}

// SetExpired sets the "expired" field.
func (m *RefreshTokenMutation) SetExpired(b bool) {
	m.expired = &b
}

// Expired returns the value of the "expired" field in the mutation.
func (m *RefreshTokenMutation) Expired() (r bool, exists bool) {
	v := m.expired
	if v == nil {
		return
	}
	return *v, true
}

// OldExpired returns the old "expired" field's value of the RefreshToken entity.
// If the RefreshToken object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RefreshTokenMutation) OldExpired(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExpired is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExpired requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpired: %w", err)
	}
	return oldValue.Expired, nil
}

// ResetExpired resets all changes to the "expired" field.
func (m *RefreshTokenMutation) ResetExpired() {
	m.expired = nil
}

// Where appends a list predicates to the RefreshTokenMutation builder.
func (m *RefreshTokenMutation) Where(ps ...predicate.RefreshToken) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RefreshTokenMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (RefreshToken).
func (m *RefreshTokenMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RefreshTokenMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.token != nil {
		fields = append(fields, refreshtoken.FieldToken)
	}
	if m.expired != nil {
		fields = append(fields, refreshtoken.FieldExpired)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RefreshTokenMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case refreshtoken.FieldToken:
		return m.Token()
	case refreshtoken.FieldExpired:
		return m.Expired()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RefreshTokenMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case refreshtoken.FieldToken:
		return m.OldToken(ctx)
	case refreshtoken.FieldExpired:
		return m.OldExpired(ctx)
	}
	return nil, fmt.Errorf("unknown RefreshToken field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RefreshTokenMutation) SetField(name string, value ent.Value) error {
	switch name {
	case refreshtoken.FieldToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetToken(v)
		return nil
	case refreshtoken.FieldExpired:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpired(v)
		return nil
	}
	return fmt.Errorf("unknown RefreshToken field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RefreshTokenMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RefreshTokenMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RefreshTokenMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown RefreshToken numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RefreshTokenMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RefreshTokenMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RefreshTokenMutation) ClearField(name string) error {
	return fmt.Errorf("unknown RefreshToken nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RefreshTokenMutation) ResetField(name string) error {
	switch name {
	case refreshtoken.FieldToken:
		m.ResetToken()
		return nil
	case refreshtoken.FieldExpired:
		m.ResetExpired()
		return nil
	}
	return fmt.Errorf("unknown RefreshToken field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RefreshTokenMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RefreshTokenMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RefreshTokenMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RefreshTokenMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RefreshTokenMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RefreshTokenMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RefreshTokenMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown RefreshToken unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RefreshTokenMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown RefreshToken edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	email         *string
	password      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}