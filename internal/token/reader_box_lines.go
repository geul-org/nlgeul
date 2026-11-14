//ff:func feature=token type=reader control=sequence
//ff:what BoxLines — a sentence read as a stream and drawn as boxes: a box for each word, the ways of a stretch cut more than one way at the right in two brackets, and the best way at the end (boxLinesOf). head cuts the candidate lines of a box and the ways of a stretch (0 for all); trace puts the trace lines of each unit before its boxes
package token

func (r *Reader) BoxLines(sentence string, trace bool, head int) []string {
	st := r.Stream(sentence)
	if trace {
		st = r.StreamTrace(sentence)
	}
	return boxLinesOf(st, lettersOf(sentence), head)
}
