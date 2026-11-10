//ff:func feature=parse type=reader control=sequence
//ff:what plainStart — where a whole-syllable ending starts at the end of s[:end], with 으 before it when the ending takes one
package parse

func plainStart(s []rune, end int, e ending) (int, bool) {
	after := e.after
	if e.flags&fEu != 0 {
		after = afterOpen
		start, ok := suffixAt(s, end, "으"+e.form)
		if ok && start > 0 && fits(s[start-1], afterBatchim) {
			return start, true
		}
	}
	start, ok := suffixAt(s, end, e.form)
	return start, ok && start > 0 && fits(s[start-1], after)
}
