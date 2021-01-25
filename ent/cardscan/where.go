// Code generated by entc, DO NOT EDIT.

package cardscan

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/mikejwhitehead/jukebox/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ScanAt applies equality check predicate on the "scan_at" field. It's identical to ScanAtEQ.
func ScanAt(v time.Time) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScanAt), v))
	})
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// ScanAtEQ applies the EQ predicate on the "scan_at" field.
func ScanAtEQ(v time.Time) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScanAt), v))
	})
}

// ScanAtNEQ applies the NEQ predicate on the "scan_at" field.
func ScanAtNEQ(v time.Time) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScanAt), v))
	})
}

// ScanAtIn applies the In predicate on the "scan_at" field.
func ScanAtIn(vs ...time.Time) predicate.CardScan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScanAt), v...))
	})
}

// ScanAtNotIn applies the NotIn predicate on the "scan_at" field.
func ScanAtNotIn(vs ...time.Time) predicate.CardScan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScanAt), v...))
	})
}

// ScanAtGT applies the GT predicate on the "scan_at" field.
func ScanAtGT(v time.Time) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScanAt), v))
	})
}

// ScanAtGTE applies the GTE predicate on the "scan_at" field.
func ScanAtGTE(v time.Time) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScanAt), v))
	})
}

// ScanAtLT applies the LT predicate on the "scan_at" field.
func ScanAtLT(v time.Time) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScanAt), v))
	})
}

// ScanAtLTE applies the LTE predicate on the "scan_at" field.
func ScanAtLTE(v time.Time) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScanAt), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.CardScan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.CardScan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResult), v))
	})
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...string) predicate.CardScan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResult), v...))
	})
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...string) predicate.CardScan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardScan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResult), v...))
	})
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResult), v))
	})
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResult), v))
	})
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResult), v))
	})
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResult), v))
	})
}

// ResultContains applies the Contains predicate on the "result" field.
func ResultContains(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResult), v))
	})
}

// ResultHasPrefix applies the HasPrefix predicate on the "result" field.
func ResultHasPrefix(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResult), v))
	})
}

// ResultHasSuffix applies the HasSuffix predicate on the "result" field.
func ResultHasSuffix(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResult), v))
	})
}

// ResultIsNil applies the IsNil predicate on the "result" field.
func ResultIsNil() predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldResult)))
	})
}

// ResultNotNil applies the NotNil predicate on the "result" field.
func ResultNotNil() predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldResult)))
	})
}

// ResultEqualFold applies the EqualFold predicate on the "result" field.
func ResultEqualFold(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResult), v))
	})
}

// ResultContainsFold applies the ContainsFold predicate on the "result" field.
func ResultContainsFold(v string) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResult), v))
	})
}

// HasPlaylist applies the HasEdge predicate on the "playlist" edge.
func HasPlaylist() predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaylistTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PlaylistTable, PlaylistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaylistWith applies the HasEdge predicate on the "playlist" edge with a given conditions (other predicates).
func HasPlaylistWith(preds ...predicate.Playlist) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaylistInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PlaylistTable, PlaylistColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCard applies the HasEdge predicate on the "card" edge.
func HasCard() predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CardTable, CardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCardWith applies the HasEdge predicate on the "card" edge with a given conditions (other predicates).
func HasCardWith(preds ...predicate.Card) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CardTable, CardColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CardScan) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CardScan) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CardScan) predicate.CardScan {
	return predicate.CardScan(func(s *sql.Selector) {
		p(s.Not())
	})
}