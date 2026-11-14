//ff:type feature=token type=model
//ff:what Stream — a sentence read as the layer above takes it, one stretch at a time: the units of the sentence, the unit being read (cur: its stretches are worked out as they are asked for, k of them taken, glued when a space did not stand before it), the regions read and not yet sent, those sent, how many letters came before the next unit, whether a space stood after the unit read last, the trace lines of the units read since they were last taken, what the nets count of themselves (stats), and the region read again for the stretch Again was asked of last (redone, of the stretch redoOf)
package token

// A stream holds nothing of the layer above: the same sentence gives the
// same segments whatever was asked in between.
type Stream struct {
	ix     *index
	trace  bool
	units  []unit
	read   int
	spaced bool
	cursor int
	queue  []*region
	sent   []*region
	notes  []string
	cur    *unitReader
	k      int
	glued  bool
	stats  netStats
	redone *region
	redoOf *region
}
