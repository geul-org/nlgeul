//ff:func feature=token type=helper control=iteration dimension=1
//ff:what netRegions — the regions a stream of the sentence queues when every unit has been read (by the net), in order
package token

func netRegions(r *Reader, sentence string) []*region {
	st := r.Stream(sentence)
	for st.read < len(st.units) {
		st.readNext()
	}
	return st.queue
}
