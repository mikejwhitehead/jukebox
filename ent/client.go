// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/mikejwhitehead/jukebox/ent/migrate"

	"github.com/mikejwhitehead/jukebox/ent/card"
	"github.com/mikejwhitehead/jukebox/ent/cardscan"
	"github.com/mikejwhitehead/jukebox/ent/playlist"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Card is the client for interacting with the Card builders.
	Card *CardClient
	// CardScan is the client for interacting with the CardScan builders.
	CardScan *CardScanClient
	// Playlist is the client for interacting with the Playlist builders.
	Playlist *PlaylistClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Card = NewCardClient(c.config)
	c.CardScan = NewCardScanClient(c.config)
	c.Playlist = NewPlaylistClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Card:     NewCardClient(cfg),
		CardScan: NewCardScanClient(cfg),
		Playlist: NewPlaylistClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:   cfg,
		Card:     NewCardClient(cfg),
		CardScan: NewCardScanClient(cfg),
		Playlist: NewPlaylistClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Card.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Card.Use(hooks...)
	c.CardScan.Use(hooks...)
	c.Playlist.Use(hooks...)
}

// CardClient is a client for the Card schema.
type CardClient struct {
	config
}

// NewCardClient returns a client for the Card from the given config.
func NewCardClient(c config) *CardClient {
	return &CardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `card.Hooks(f(g(h())))`.
func (c *CardClient) Use(hooks ...Hook) {
	c.hooks.Card = append(c.hooks.Card, hooks...)
}

// Create returns a create builder for Card.
func (c *CardClient) Create() *CardCreate {
	mutation := newCardMutation(c.config, OpCreate)
	return &CardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Card entities.
func (c *CardClient) CreateBulk(builders ...*CardCreate) *CardCreateBulk {
	return &CardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Card.
func (c *CardClient) Update() *CardUpdate {
	mutation := newCardMutation(c.config, OpUpdate)
	return &CardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardClient) UpdateOne(ca *Card) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCard(ca))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardClient) UpdateOneID(id int) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCardID(id))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Card.
func (c *CardClient) Delete() *CardDelete {
	mutation := newCardMutation(c.config, OpDelete)
	return &CardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CardClient) DeleteOne(ca *Card) *CardDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CardClient) DeleteOneID(id int) *CardDeleteOne {
	builder := c.Delete().Where(card.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardDeleteOne{builder}
}

// Query returns a query builder for Card.
func (c *CardClient) Query() *CardQuery {
	return &CardQuery{config: c.config}
}

// Get returns a Card entity by its id.
func (c *CardClient) Get(ctx context.Context, id int) (*Card, error) {
	return c.Query().Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardClient) GetX(ctx context.Context, id int) *Card {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlaylist queries the playlist edge of a Card.
func (c *CardClient) QueryPlaylist(ca *Card) *PlaylistQuery {
	query := &PlaylistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, id),
			sqlgraph.To(playlist.Table, playlist.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, card.PlaylistTable, card.PlaylistColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryScans queries the scans edge of a Card.
func (c *CardClient) QueryScans(ca *Card) *CardScanQuery {
	query := &CardScanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, id),
			sqlgraph.To(cardscan.Table, cardscan.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, card.ScansTable, card.ScansColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CardClient) Hooks() []Hook {
	return c.hooks.Card
}

// CardScanClient is a client for the CardScan schema.
type CardScanClient struct {
	config
}

// NewCardScanClient returns a client for the CardScan from the given config.
func NewCardScanClient(c config) *CardScanClient {
	return &CardScanClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cardscan.Hooks(f(g(h())))`.
func (c *CardScanClient) Use(hooks ...Hook) {
	c.hooks.CardScan = append(c.hooks.CardScan, hooks...)
}

// Create returns a create builder for CardScan.
func (c *CardScanClient) Create() *CardScanCreate {
	mutation := newCardScanMutation(c.config, OpCreate)
	return &CardScanCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CardScan entities.
func (c *CardScanClient) CreateBulk(builders ...*CardScanCreate) *CardScanCreateBulk {
	return &CardScanCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CardScan.
func (c *CardScanClient) Update() *CardScanUpdate {
	mutation := newCardScanMutation(c.config, OpUpdate)
	return &CardScanUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardScanClient) UpdateOne(cs *CardScan) *CardScanUpdateOne {
	mutation := newCardScanMutation(c.config, OpUpdateOne, withCardScan(cs))
	return &CardScanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardScanClient) UpdateOneID(id int) *CardScanUpdateOne {
	mutation := newCardScanMutation(c.config, OpUpdateOne, withCardScanID(id))
	return &CardScanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CardScan.
func (c *CardScanClient) Delete() *CardScanDelete {
	mutation := newCardScanMutation(c.config, OpDelete)
	return &CardScanDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CardScanClient) DeleteOne(cs *CardScan) *CardScanDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CardScanClient) DeleteOneID(id int) *CardScanDeleteOne {
	builder := c.Delete().Where(cardscan.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardScanDeleteOne{builder}
}

// Query returns a query builder for CardScan.
func (c *CardScanClient) Query() *CardScanQuery {
	return &CardScanQuery{config: c.config}
}

// Get returns a CardScan entity by its id.
func (c *CardScanClient) Get(ctx context.Context, id int) (*CardScan, error) {
	return c.Query().Where(cardscan.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardScanClient) GetX(ctx context.Context, id int) *CardScan {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlaylist queries the playlist edge of a CardScan.
func (c *CardScanClient) QueryPlaylist(cs *CardScan) *PlaylistQuery {
	query := &PlaylistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cardscan.Table, cardscan.FieldID, id),
			sqlgraph.To(playlist.Table, playlist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cardscan.PlaylistTable, cardscan.PlaylistColumn),
		)
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCard queries the card edge of a CardScan.
func (c *CardScanClient) QueryCard(cs *CardScan) *CardQuery {
	query := &CardQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cardscan.Table, cardscan.FieldID, id),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cardscan.CardTable, cardscan.CardColumn),
		)
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CardScanClient) Hooks() []Hook {
	return c.hooks.CardScan
}

// PlaylistClient is a client for the Playlist schema.
type PlaylistClient struct {
	config
}

// NewPlaylistClient returns a client for the Playlist from the given config.
func NewPlaylistClient(c config) *PlaylistClient {
	return &PlaylistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `playlist.Hooks(f(g(h())))`.
func (c *PlaylistClient) Use(hooks ...Hook) {
	c.hooks.Playlist = append(c.hooks.Playlist, hooks...)
}

// Create returns a create builder for Playlist.
func (c *PlaylistClient) Create() *PlaylistCreate {
	mutation := newPlaylistMutation(c.config, OpCreate)
	return &PlaylistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Playlist entities.
func (c *PlaylistClient) CreateBulk(builders ...*PlaylistCreate) *PlaylistCreateBulk {
	return &PlaylistCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Playlist.
func (c *PlaylistClient) Update() *PlaylistUpdate {
	mutation := newPlaylistMutation(c.config, OpUpdate)
	return &PlaylistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlaylistClient) UpdateOne(pl *Playlist) *PlaylistUpdateOne {
	mutation := newPlaylistMutation(c.config, OpUpdateOne, withPlaylist(pl))
	return &PlaylistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlaylistClient) UpdateOneID(id int) *PlaylistUpdateOne {
	mutation := newPlaylistMutation(c.config, OpUpdateOne, withPlaylistID(id))
	return &PlaylistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Playlist.
func (c *PlaylistClient) Delete() *PlaylistDelete {
	mutation := newPlaylistMutation(c.config, OpDelete)
	return &PlaylistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlaylistClient) DeleteOne(pl *Playlist) *PlaylistDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlaylistClient) DeleteOneID(id int) *PlaylistDeleteOne {
	builder := c.Delete().Where(playlist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlaylistDeleteOne{builder}
}

// Query returns a query builder for Playlist.
func (c *PlaylistClient) Query() *PlaylistQuery {
	return &PlaylistQuery{config: c.config}
}

// Get returns a Playlist entity by its id.
func (c *PlaylistClient) Get(ctx context.Context, id int) (*Playlist, error) {
	return c.Query().Where(playlist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlaylistClient) GetX(ctx context.Context, id int) *Playlist {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCard queries the card edge of a Playlist.
func (c *PlaylistClient) QueryCard(pl *Playlist) *CardQuery {
	query := &CardQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(playlist.Table, playlist.FieldID, id),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, playlist.CardTable, playlist.CardColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCardScans queries the card_scans edge of a Playlist.
func (c *PlaylistClient) QueryCardScans(pl *Playlist) *CardScanQuery {
	query := &CardScanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(playlist.Table, playlist.FieldID, id),
			sqlgraph.To(cardscan.Table, cardscan.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, playlist.CardScansTable, playlist.CardScansColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PlaylistClient) Hooks() []Hook {
	return c.hooks.Playlist
}
