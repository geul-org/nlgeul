//ff:func feature=mini type=rule control=sequence
//ff:what mark — the word with "+" between its pieces: before each particle of the chain when a particle reading stands, else as the standing tail reading divides it, else as it is
package mini

import "strings"

func mark(w *word, r *ruling) string {
	if r.chain {
		s := []rune(w.text)
		n := len(s) - len([]rune(strings.ReplaceAll(w.chain.parts, "+", "")))
		return string(s[:n]) + "+" + w.chain.parts
	}
	if r.tail != "" {
		return r.tail
	}
	return w.text
}
