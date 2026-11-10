//ff:func feature=token type=reader control=iteration dimension=1
//ff:what Lines — read a token stream and send up its words: the sentence is cut at spaces (fixed boundaries) into units, and each unit comes back as bundles of words, one line each
package token

import "strings"

// A unit is read without looking outside it, so the same unit always
// gives the same lines.
func Lines(sentence string, trace bool) []string {
	ix := newIndex()
	var out []string
	for _, field := range strings.Fields(sentence) {
		out = append(out, readUnit(ix, field, trace)...)
	}
	return out
}
