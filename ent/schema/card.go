package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Time("created").
			Default(time.Now).
			Immutable(),
		field.Time("modified").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playlist", Playlist.Type).
			Ref("card").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
		edge.From("scans", CardScan.Type).
			Ref("card"),
	}
}
