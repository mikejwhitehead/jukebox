package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// CardScan holds the schema definition for the CardScan entity.
type CardScan struct {
	ent.Schema
}

// Fields of the CardScan.
func (CardScan) Fields() []ent.Field {
	return []ent.Field{
		field.Time("scan_at").
			Default(time.Now).
			Immutable(),
		field.Enum("state").
			Values(
				"success",
				"error",
			).
			Default("success"),
		field.String("result").
			Nillable().
			Optional(),
	}
}

// Edges of the CardScan.
func (CardScan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("playlist", Playlist.Type).
			Required().Unique(),
		edge.To("card", Card.Type).
			Required().Unique(),
	}
}
