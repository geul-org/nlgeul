//ff:func feature=token type=reader control=sequence
//ff:what newStream — a stream over a sentence: its units cut out (nothing read yet), with their trace lines when asked for, and the index's rules of a sentence
package token

func newStream(ix *index, sentence string, trace bool) *Stream {
	return &Stream{ix: ix, trace: trace, units: unitsOf(ix, sentence, trace), spaced: true, rules: ix.sentence}
}
