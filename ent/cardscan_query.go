// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/mikejwhitehead/jukebox/ent/card"
	"github.com/mikejwhitehead/jukebox/ent/cardscan"
	"github.com/mikejwhitehead/jukebox/ent/playlist"
	"github.com/mikejwhitehead/jukebox/ent/predicate"
)

// CardScanQuery is the builder for querying CardScan entities.
type CardScanQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.CardScan
	// eager-loading edges.
	withPlaylist *PlaylistQuery
	withCard     *CardQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CardScanQuery builder.
func (csq *CardScanQuery) Where(ps ...predicate.CardScan) *CardScanQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit adds a limit step to the query.
func (csq *CardScanQuery) Limit(limit int) *CardScanQuery {
	csq.limit = &limit
	return csq
}

// Offset adds an offset step to the query.
func (csq *CardScanQuery) Offset(offset int) *CardScanQuery {
	csq.offset = &offset
	return csq
}

// Order adds an order step to the query.
func (csq *CardScanQuery) Order(o ...OrderFunc) *CardScanQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// QueryPlaylist chains the current query on the "playlist" edge.
func (csq *CardScanQuery) QueryPlaylist() *PlaylistQuery {
	query := &PlaylistQuery{config: csq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cardscan.Table, cardscan.FieldID, selector),
			sqlgraph.To(playlist.Table, playlist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cardscan.PlaylistTable, cardscan.PlaylistColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCard chains the current query on the "card" edge.
func (csq *CardScanQuery) QueryCard() *CardQuery {
	query := &CardQuery{config: csq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cardscan.Table, cardscan.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cardscan.CardTable, cardscan.CardColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CardScan entity from the query.
// Returns a *NotFoundError when no CardScan was found.
func (csq *CardScanQuery) First(ctx context.Context) (*CardScan, error) {
	nodes, err := csq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cardscan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *CardScanQuery) FirstX(ctx context.Context) *CardScan {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CardScan ID from the query.
// Returns a *NotFoundError when no CardScan ID was found.
func (csq *CardScanQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cardscan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *CardScanQuery) FirstIDX(ctx context.Context) int {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CardScan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CardScan entity is not found.
// Returns a *NotFoundError when no CardScan entities are found.
func (csq *CardScanQuery) Only(ctx context.Context) (*CardScan, error) {
	nodes, err := csq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cardscan.Label}
	default:
		return nil, &NotSingularError{cardscan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *CardScanQuery) OnlyX(ctx context.Context) *CardScan {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CardScan ID in the query.
// Returns a *NotSingularError when exactly one CardScan ID is not found.
// Returns a *NotFoundError when no entities are found.
func (csq *CardScanQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = &NotSingularError{cardscan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *CardScanQuery) OnlyIDX(ctx context.Context) int {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CardScans.
func (csq *CardScanQuery) All(ctx context.Context) ([]*CardScan, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return csq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (csq *CardScanQuery) AllX(ctx context.Context) []*CardScan {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CardScan IDs.
func (csq *CardScanQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := csq.Select(cardscan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *CardScanQuery) IDsX(ctx context.Context) []int {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *CardScanQuery) Count(ctx context.Context) (int, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return csq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (csq *CardScanQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *CardScanQuery) Exist(ctx context.Context) (bool, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return csq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *CardScanQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CardScanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *CardScanQuery) Clone() *CardScanQuery {
	if csq == nil {
		return nil
	}
	return &CardScanQuery{
		config:       csq.config,
		limit:        csq.limit,
		offset:       csq.offset,
		order:        append([]OrderFunc{}, csq.order...),
		predicates:   append([]predicate.CardScan{}, csq.predicates...),
		withPlaylist: csq.withPlaylist.Clone(),
		withCard:     csq.withCard.Clone(),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

// WithPlaylist tells the query-builder to eager-load the nodes that are connected to
// the "playlist" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CardScanQuery) WithPlaylist(opts ...func(*PlaylistQuery)) *CardScanQuery {
	query := &PlaylistQuery{config: csq.config}
	for _, opt := range opts {
		opt(query)
	}
	csq.withPlaylist = query
	return csq
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CardScanQuery) WithCard(opts ...func(*CardQuery)) *CardScanQuery {
	query := &CardQuery{config: csq.config}
	for _, opt := range opts {
		opt(query)
	}
	csq.withCard = query
	return csq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ScanAt time.Time `json:"scan_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CardScan.Query().
//		GroupBy(cardscan.FieldScanAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (csq *CardScanQuery) GroupBy(field string, fields ...string) *CardScanGroupBy {
	group := &CardScanGroupBy{config: csq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return csq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ScanAt time.Time `json:"scan_at,omitempty"`
//	}
//
//	client.CardScan.Query().
//		Select(cardscan.FieldScanAt).
//		Scan(ctx, &v)
//
func (csq *CardScanQuery) Select(field string, fields ...string) *CardScanSelect {
	csq.fields = append([]string{field}, fields...)
	return &CardScanSelect{CardScanQuery: csq}
}

func (csq *CardScanQuery) prepareQuery(ctx context.Context) error {
	for _, f := range csq.fields {
		if !cardscan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *CardScanQuery) sqlAll(ctx context.Context) ([]*CardScan, error) {
	var (
		nodes       = []*CardScan{}
		withFKs     = csq.withFKs
		_spec       = csq.querySpec()
		loadedTypes = [2]bool{
			csq.withPlaylist != nil,
			csq.withCard != nil,
		}
	)
	if csq.withPlaylist != nil || csq.withCard != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cardscan.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CardScan{config: csq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := csq.withPlaylist; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CardScan)
		for i := range nodes {
			if fk := nodes[i].card_scan_playlist; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(playlist.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "card_scan_playlist" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Playlist = n
			}
		}
	}

	if query := csq.withCard; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CardScan)
		for i := range nodes {
			if fk := nodes[i].card_scan_card; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(card.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "card_scan_card" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Card = n
			}
		}
	}

	return nodes, nil
}

func (csq *CardScanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CardScanQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := csq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (csq *CardScanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardscan.Table,
			Columns: cardscan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardscan.FieldID,
			},
		},
		From:   csq.sql,
		Unique: true,
	}
	if fields := csq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cardscan.FieldID)
		for i := range fields {
			if fields[i] != cardscan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, cardscan.ValidColumn)
			}
		}
	}
	return _spec
}

func (csq *CardScanQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(cardscan.Table)
	selector := builder.Select(t1.Columns(cardscan.Columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(cardscan.Columns...)...)
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector, cardscan.ValidColumn)
	}
	if offset := csq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CardScanGroupBy is the group-by builder for CardScan entities.
type CardScanGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CardScanGroupBy) Aggregate(fns ...AggregateFunc) *CardScanGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the group-by query and scans the result into the given value.
func (csgb *CardScanGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := csgb.path(ctx)
	if err != nil {
		return err
	}
	csgb.sql = query
	return csgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (csgb *CardScanGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := csgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CardScanGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (csgb *CardScanGroupBy) StringsX(ctx context.Context) []string {
	v, err := csgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = csgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (csgb *CardScanGroupBy) StringX(ctx context.Context) string {
	v, err := csgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CardScanGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (csgb *CardScanGroupBy) IntsX(ctx context.Context) []int {
	v, err := csgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = csgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (csgb *CardScanGroupBy) IntX(ctx context.Context) int {
	v, err := csgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CardScanGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (csgb *CardScanGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := csgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = csgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (csgb *CardScanGroupBy) Float64X(ctx context.Context) float64 {
	v, err := csgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CardScanGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (csgb *CardScanGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := csgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CardScanGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = csgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (csgb *CardScanGroupBy) BoolX(ctx context.Context) bool {
	v, err := csgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (csgb *CardScanGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range csgb.fields {
		if !cardscan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := csgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (csgb *CardScanGroupBy) sqlQuery() *sql.Selector {
	selector := csgb.sql
	columns := make([]string, 0, len(csgb.fields)+len(csgb.fns))
	columns = append(columns, csgb.fields...)
	for _, fn := range csgb.fns {
		columns = append(columns, fn(selector, cardscan.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(csgb.fields...)
}

// CardScanSelect is the builder for selecting fields of CardScan entities.
type CardScanSelect struct {
	*CardScanQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (css *CardScanSelect) Scan(ctx context.Context, v interface{}) error {
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	css.sql = css.CardScanQuery.sqlQuery()
	return css.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (css *CardScanSelect) ScanX(ctx context.Context, v interface{}) {
	if err := css.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) Strings(ctx context.Context) ([]string, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CardScanSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (css *CardScanSelect) StringsX(ctx context.Context) []string {
	v, err := css.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = css.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (css *CardScanSelect) StringX(ctx context.Context) string {
	v, err := css.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) Ints(ctx context.Context) ([]int, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CardScanSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (css *CardScanSelect) IntsX(ctx context.Context) []int {
	v, err := css.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = css.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (css *CardScanSelect) IntX(ctx context.Context) int {
	v, err := css.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CardScanSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (css *CardScanSelect) Float64sX(ctx context.Context) []float64 {
	v, err := css.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = css.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (css *CardScanSelect) Float64X(ctx context.Context) float64 {
	v, err := css.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CardScanSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (css *CardScanSelect) BoolsX(ctx context.Context) []bool {
	v, err := css.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (css *CardScanSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = css.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardscan.Label}
	default:
		err = fmt.Errorf("ent: CardScanSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (css *CardScanSelect) BoolX(ctx context.Context) bool {
	v, err := css.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (css *CardScanSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := css.sqlQuery().Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (css *CardScanSelect) sqlQuery() sql.Querier {
	selector := css.sql
	selector.Select(selector.Columns(css.fields...)...)
	return selector
}
