//ff:func feature=token type=reader control=sequence
//ff:what StreamTrace — a Stream that also keeps the trace lines of each unit it reads (what --trace prints), to be taken with Notes
package token

func (r *Reader) StreamTrace(sentence string) *Stream {
	return newStream(r.ix, sentence, true)
}
