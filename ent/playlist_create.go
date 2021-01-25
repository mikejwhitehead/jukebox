// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/mikejwhitehead/jukebox/ent/card"
	"github.com/mikejwhitehead/jukebox/ent/cardscan"
	"github.com/mikejwhitehead/jukebox/ent/playlist"
)

// PlaylistCreate is the builder for creating a Playlist entity.
type PlaylistCreate struct {
	config
	mutation *PlaylistMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PlaylistCreate) SetName(s string) *PlaylistCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetCreated sets the "created" field.
func (pc *PlaylistCreate) SetCreated(t time.Time) *PlaylistCreate {
	pc.mutation.SetCreated(t)
	return pc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableCreated(t *time.Time) *PlaylistCreate {
	if t != nil {
		pc.SetCreated(*t)
	}
	return pc
}

// SetModified sets the "modified" field.
func (pc *PlaylistCreate) SetModified(t time.Time) *PlaylistCreate {
	pc.mutation.SetModified(t)
	return pc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableModified(t *time.Time) *PlaylistCreate {
	if t != nil {
		pc.SetModified(*t)
	}
	return pc
}

// SetMediaURL sets the "media_url" field.
func (pc *PlaylistCreate) SetMediaURL(s string) *PlaylistCreate {
	pc.mutation.SetMediaURL(s)
	return pc
}

// SetCardID sets the "card" edge to the Card entity by ID.
func (pc *PlaylistCreate) SetCardID(id int) *PlaylistCreate {
	pc.mutation.SetCardID(id)
	return pc
}

// SetNillableCardID sets the "card" edge to the Card entity by ID if the given value is not nil.
func (pc *PlaylistCreate) SetNillableCardID(id *int) *PlaylistCreate {
	if id != nil {
		pc = pc.SetCardID(*id)
	}
	return pc
}

// SetCard sets the "card" edge to the Card entity.
func (pc *PlaylistCreate) SetCard(c *Card) *PlaylistCreate {
	return pc.SetCardID(c.ID)
}

// AddCardScanIDs adds the "card_scans" edge to the CardScan entity by IDs.
func (pc *PlaylistCreate) AddCardScanIDs(ids ...int) *PlaylistCreate {
	pc.mutation.AddCardScanIDs(ids...)
	return pc
}

// AddCardScans adds the "card_scans" edges to the CardScan entity.
func (pc *PlaylistCreate) AddCardScans(c ...*CardScan) *PlaylistCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCardScanIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pc *PlaylistCreate) Mutation() *PlaylistMutation {
	return pc.mutation
}

// Save creates the Playlist in the database.
func (pc *PlaylistCreate) Save(ctx context.Context) (*Playlist, error) {
	var (
		err  error
		node *Playlist
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlaylistCreate) SaveX(ctx context.Context) *Playlist {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PlaylistCreate) defaults() {
	if _, ok := pc.mutation.Created(); !ok {
		v := playlist.DefaultCreated()
		pc.mutation.SetCreated(v)
	}
	if _, ok := pc.mutation.Modified(); !ok {
		v := playlist.DefaultModified()
		pc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlaylistCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := pc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New("ent: missing required field \"created\"")}
	}
	if _, ok := pc.mutation.Modified(); !ok {
		return &ValidationError{Name: "modified", err: errors.New("ent: missing required field \"modified\"")}
	}
	if _, ok := pc.mutation.MediaURL(); !ok {
		return &ValidationError{Name: "media_url", err: errors.New("ent: missing required field \"media_url\"")}
	}
	return nil
}

func (pc *PlaylistCreate) sqlSave(ctx context.Context) (*Playlist, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PlaylistCreate) createSpec() (*Playlist, *sqlgraph.CreateSpec) {
	var (
		_node = &Playlist{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: playlist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlist.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playlist.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := pc.mutation.Modified(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playlist.FieldModified,
		})
		_node.Modified = value
	}
	if value, ok := pc.mutation.MediaURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldMediaURL,
		})
		_node.MediaURL = value
	}
	if nodes := pc.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   playlist.CardTable,
			Columns: []string{playlist.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CardScansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   playlist.CardScansTable,
			Columns: []string{playlist.CardScansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cardscan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaylistCreateBulk is the builder for creating many Playlist entities in bulk.
type PlaylistCreateBulk struct {
	config
	builders []*PlaylistCreate
}

// Save creates the Playlist entities in the database.
func (pcb *PlaylistCreateBulk) Save(ctx context.Context) ([]*Playlist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Playlist, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaylistMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlaylistCreateBulk) SaveX(ctx context.Context) []*Playlist {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
