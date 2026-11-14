//ff:type feature=token type=model
//ff:what Stream — a sentence read as the layer above takes it, one stretch at a time: the units of the sentence (read one by one when a stretch is asked for), the regions read and not yet sent, those sent, how many letters came before the next unit, whether a space stood after the unit read last, and the trace lines of the units read since they were last taken
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
}
