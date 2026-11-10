//ff:func feature=parse type=reader control=selection
//ff:what matchEnding — read one closing at the end of s[:end] by the way it joins: whole syllables, 어/아, or a final consonant
package parse

func matchEnding(s []rune, end int, e ending, yo bool) (tailFact, bool) {
	switch e.lead {
	case leadNone:
		return matchPlain(s, end, e)
	case leadA:
		return matchA(s, end, e, yo)
	}
	return matchJamo(s, end, e)
}
