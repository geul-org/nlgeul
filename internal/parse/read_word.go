//ff:func feature=parse type=reader control=sequence
//ff:what readWord — read one word's right end against the tables and report what was read; no constituent is named here
package parse

func readWord(text string, last bool, sent *sentence) *word {
	s := core(text)
	w := &word{text: string(s), chain: readChain(s), last: last, sent: sent}
	w.tails = collectTails(s, len(s), false)
	if len(s) > 1 && s[len(s)-1] == '요' {
		w.tails = append(w.tails, collectTails(s, len(s)-1, true)...)
	}
	join(w, s)
	w.longest = longest(w.tails)
	w.bound = boundHead(w)
	return w
}
