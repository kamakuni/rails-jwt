// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kamakuni/rails-jwt/auth/app/ent/oauthclient"
	"github.com/kamakuni/rails-jwt/auth/app/ent/predicate"
)

// OAuthClientQuery is the builder for querying OAuthClient entities.
type OAuthClientQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OAuthClient
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OAuthClientQuery builder.
func (ocq *OAuthClientQuery) Where(ps ...predicate.OAuthClient) *OAuthClientQuery {
	ocq.predicates = append(ocq.predicates, ps...)
	return ocq
}

// Limit adds a limit step to the query.
func (ocq *OAuthClientQuery) Limit(limit int) *OAuthClientQuery {
	ocq.limit = &limit
	return ocq
}

// Offset adds an offset step to the query.
func (ocq *OAuthClientQuery) Offset(offset int) *OAuthClientQuery {
	ocq.offset = &offset
	return ocq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ocq *OAuthClientQuery) Unique(unique bool) *OAuthClientQuery {
	ocq.unique = &unique
	return ocq
}

// Order adds an order step to the query.
func (ocq *OAuthClientQuery) Order(o ...OrderFunc) *OAuthClientQuery {
	ocq.order = append(ocq.order, o...)
	return ocq
}

// First returns the first OAuthClient entity from the query.
// Returns a *NotFoundError when no OAuthClient was found.
func (ocq *OAuthClientQuery) First(ctx context.Context) (*OAuthClient, error) {
	nodes, err := ocq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oauthclient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ocq *OAuthClientQuery) FirstX(ctx context.Context) *OAuthClient {
	node, err := ocq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OAuthClient ID from the query.
// Returns a *NotFoundError when no OAuthClient ID was found.
func (ocq *OAuthClientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauthclient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ocq *OAuthClientQuery) FirstIDX(ctx context.Context) int {
	id, err := ocq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OAuthClient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OAuthClient entity is found.
// Returns a *NotFoundError when no OAuthClient entities are found.
func (ocq *OAuthClientQuery) Only(ctx context.Context) (*OAuthClient, error) {
	nodes, err := ocq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oauthclient.Label}
	default:
		return nil, &NotSingularError{oauthclient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ocq *OAuthClientQuery) OnlyX(ctx context.Context) *OAuthClient {
	node, err := ocq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OAuthClient ID in the query.
// Returns a *NotSingularError when more than one OAuthClient ID is found.
// Returns a *NotFoundError when no entities are found.
func (ocq *OAuthClientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauthclient.Label}
	default:
		err = &NotSingularError{oauthclient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ocq *OAuthClientQuery) OnlyIDX(ctx context.Context) int {
	id, err := ocq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OAuthClients.
func (ocq *OAuthClientQuery) All(ctx context.Context) ([]*OAuthClient, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ocq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ocq *OAuthClientQuery) AllX(ctx context.Context) []*OAuthClient {
	nodes, err := ocq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OAuthClient IDs.
func (ocq *OAuthClientQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ocq.Select(oauthclient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ocq *OAuthClientQuery) IDsX(ctx context.Context) []int {
	ids, err := ocq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ocq *OAuthClientQuery) Count(ctx context.Context) (int, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ocq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ocq *OAuthClientQuery) CountX(ctx context.Context) int {
	count, err := ocq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ocq *OAuthClientQuery) Exist(ctx context.Context) (bool, error) {
	if err := ocq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ocq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ocq *OAuthClientQuery) ExistX(ctx context.Context) bool {
	exist, err := ocq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OAuthClientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ocq *OAuthClientQuery) Clone() *OAuthClientQuery {
	if ocq == nil {
		return nil
	}
	return &OAuthClientQuery{
		config:     ocq.config,
		limit:      ocq.limit,
		offset:     ocq.offset,
		order:      append([]OrderFunc{}, ocq.order...),
		predicates: append([]predicate.OAuthClient{}, ocq.predicates...),
		// clone intermediate query.
		sql:    ocq.sql.Clone(),
		path:   ocq.path,
		unique: ocq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClientSecret string `json:"client_secret,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OAuthClient.Query().
//		GroupBy(oauthclient.FieldClientSecret).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ocq *OAuthClientQuery) GroupBy(field string, fields ...string) *OAuthClientGroupBy {
	grbuild := &OAuthClientGroupBy{config: ocq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ocq.sqlQuery(ctx), nil
	}
	grbuild.label = oauthclient.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClientSecret string `json:"client_secret,omitempty"`
//	}
//
//	client.OAuthClient.Query().
//		Select(oauthclient.FieldClientSecret).
//		Scan(ctx, &v)
func (ocq *OAuthClientQuery) Select(fields ...string) *OAuthClientSelect {
	ocq.fields = append(ocq.fields, fields...)
	selbuild := &OAuthClientSelect{OAuthClientQuery: ocq}
	selbuild.label = oauthclient.Label
	selbuild.flds, selbuild.scan = &ocq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a OAuthClientSelect configured with the given aggregations.
func (ocq *OAuthClientQuery) Aggregate(fns ...AggregateFunc) *OAuthClientSelect {
	return ocq.Select().Aggregate(fns...)
}

func (ocq *OAuthClientQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ocq.fields {
		if !oauthclient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ocq.path != nil {
		prev, err := ocq.path(ctx)
		if err != nil {
			return err
		}
		ocq.sql = prev
	}
	return nil
}

func (ocq *OAuthClientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OAuthClient, error) {
	var (
		nodes = []*OAuthClient{}
		_spec = ocq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OAuthClient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OAuthClient{config: ocq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ocq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ocq *OAuthClientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ocq.querySpec()
	_spec.Node.Columns = ocq.fields
	if len(ocq.fields) > 0 {
		_spec.Unique = ocq.unique != nil && *ocq.unique
	}
	return sqlgraph.CountNodes(ctx, ocq.driver, _spec)
}

func (ocq *OAuthClientQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ocq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ocq *OAuthClientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oauthclient.Table,
			Columns: oauthclient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauthclient.FieldID,
			},
		},
		From:   ocq.sql,
		Unique: true,
	}
	if unique := ocq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ocq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthclient.FieldID)
		for i := range fields {
			if fields[i] != oauthclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ocq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ocq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ocq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ocq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ocq *OAuthClientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ocq.driver.Dialect())
	t1 := builder.Table(oauthclient.Table)
	columns := ocq.fields
	if len(columns) == 0 {
		columns = oauthclient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ocq.sql != nil {
		selector = ocq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ocq.unique != nil && *ocq.unique {
		selector.Distinct()
	}
	for _, p := range ocq.predicates {
		p(selector)
	}
	for _, p := range ocq.order {
		p(selector)
	}
	if offset := ocq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ocq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OAuthClientGroupBy is the group-by builder for OAuthClient entities.
type OAuthClientGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ocgb *OAuthClientGroupBy) Aggregate(fns ...AggregateFunc) *OAuthClientGroupBy {
	ocgb.fns = append(ocgb.fns, fns...)
	return ocgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ocgb *OAuthClientGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ocgb.path(ctx)
	if err != nil {
		return err
	}
	ocgb.sql = query
	return ocgb.sqlScan(ctx, v)
}

func (ocgb *OAuthClientGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ocgb.fields {
		if !oauthclient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ocgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ocgb *OAuthClientGroupBy) sqlQuery() *sql.Selector {
	selector := ocgb.sql.Select()
	aggregation := make([]string, 0, len(ocgb.fns))
	for _, fn := range ocgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ocgb.fields)+len(ocgb.fns))
		for _, f := range ocgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ocgb.fields...)...)
}

// OAuthClientSelect is the builder for selecting fields of OAuthClient entities.
type OAuthClientSelect struct {
	*OAuthClientQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ocs *OAuthClientSelect) Aggregate(fns ...AggregateFunc) *OAuthClientSelect {
	ocs.fns = append(ocs.fns, fns...)
	return ocs
}

// Scan applies the selector query and scans the result into the given value.
func (ocs *OAuthClientSelect) Scan(ctx context.Context, v any) error {
	if err := ocs.prepareQuery(ctx); err != nil {
		return err
	}
	ocs.sql = ocs.OAuthClientQuery.sqlQuery(ctx)
	return ocs.sqlScan(ctx, v)
}

func (ocs *OAuthClientSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ocs.fns))
	for _, fn := range ocs.fns {
		aggregation = append(aggregation, fn(ocs.sql))
	}
	switch n := len(*ocs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ocs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ocs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ocs.sql.Query()
	if err := ocs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
