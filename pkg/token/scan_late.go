//ff:func feature=token type=helper control=sequence
//ff:what late — whether the candidates kept at place q are not worked out yet although a sign there asks for them (a sign between letters read as a word of its own, or a cue that shows a word beginning there): the states that end at q with a word of no markers are made when they are first asked for, so a stretch cut off at q must have them made first
package token

func (sc *scan) late(q int) bool {
	if _, done := sc.kept[q]; done {
		return false
	}
	_, mid := sc.mid[q]
	return mid || len(sc.shows[q]) > 0
}
