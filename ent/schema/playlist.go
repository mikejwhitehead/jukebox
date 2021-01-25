package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Playlist holds the schema definition for the Playlist entity.
type Playlist struct {
	ent.Schema
}

// Fields of the Playlist.
func (Playlist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created").
			Default(time.Now).
			Immutable(),
		field.Time("modified").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("media_url"),
	}
}

// Edges of the Playlist.
func (Playlist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).
			Unique(),
		edge.From("card_scans", CardScan.Type).
			Ref("playlist"),
	}
}
