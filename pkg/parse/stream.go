//ff:type feature=parse type=model
//ff:what Stream — a sentence read as the layer above takes it, one stretch at a time: the units of the sentence, the unit being read (cur: its stretches are worked out as they are asked for, k of them taken, glued when a space did not stand before it), the regions read and not yet sent, those sent, how many letters came before the next unit, whether a space stood after the unit read last, the trace lines of the units read since they were last taken, what the nets count of themselves (stats), the region read again for the stretch Again was asked of last (redone, of the stretch redoOf), the rules of a sentence (the index's), whether a region not yet judged holds a candidate that waits for the predicate after it (held), the words judged so far as their first candidates read them, the nearest first (behind), and the trace lines read that no region has taken yet (pending)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// A stream holds nothing of the layer above: the same sentence gives the
// same segments whatever was asked in between.
type Stream struct {
	ix      *index
	trace   bool
	units   []unit
	read    int
	spaced  bool
	cursor  int
	queue   []*region
	sent    []*region
	notes   []string
	cur     *unitReader
	k       int
	glued   bool
	stats   netStats
	redone  *region
	redoOf  *region
	rules   []sentenceRule
	held    bool
	behind  []rule.Near
	pending []string
}
