//ff:func feature=token type=reader control=sequence
//ff:what Stream — a sentence as a stream to read from: Next sends the best way of cutting the next stretch, Again the way after the one sent last for a stretch; a stretch is read when it is asked for, and nothing is read ahead
package token

// See the README of this package for what is sent and how the places are
// counted.
func (r *Reader) Stream(sentence string) *Stream {
	return newStream(r.ix, sentence, false)
}
