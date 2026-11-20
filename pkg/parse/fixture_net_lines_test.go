//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what netLines — what follows what in the net the stream reads each unit of a sentence by (scan.net): every way through it to a state that ends the unit (stateLines), unit after unit, so a line holds only words the net lets follow one another, which the lines of the stream (streamLines: one candidate of each node of a way) do not tell
package parse

func netLines(r *Reader, sentence string) []string {
	var out []string
	for _, u := range unitsOf(r.ix, sentence, false) {
		lat := unitScan(r.ix, u).net()
		if lat == nil {
			continue
		}
		for _, e := range lat.ends {
			out = append(out, stateLines(e)...)
		}
	}
	return out
}
