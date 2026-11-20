//ff:func feature=parse type=reader control=sequence
//ff:what newStream — a stream over a sentence: its units cut out (nothing read yet), with their trace lines when asked for, and the index's rules of a sentence
package parse

func newStream(ix *index, sentence string, trace bool) *Stream {
	units := unitsOf(ix, sentence, trace)
	return &Stream{ix: ix, trace: trace, units: units, spaced: true, rules: ix.sentence, places: placesOf(units)}
}
