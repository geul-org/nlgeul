//ff:func feature=parse type=reader control=sequence
//ff:what matchPlain — read a whole-syllable closing at the end of s[:end]
package parse

func matchPlain(s []rune, end int, e ending) (tailFact, bool) {
	start, ok := plainStart(s, end, e)
	if !ok {
		return tailFact{}, false
	}
	return finish(s, start, e, 2*(end-start), closingMark(e, string(s[start:end])))
}
