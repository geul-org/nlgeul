//ff:func feature=token type=rule control=iteration dimension=1
//ff:what known — is a stem a known word with what follows it: its name, and true when a table has it and (from the second layer on) one of its kinds takes that tail
package token

import "strings"

// Without a word table only the closed-set words are known, whatever
// follows them, under their own names.
func (ix *index) known(stem string, tail int) (string, bool) {
	c, closed := ix.words[stem]
	if ix.layers == layerNone {
		if closed {
			return c.name, true
		}
		return "", false
	}
	kinds := ix.closed[stem] | ix.lex[stem].kinds
	var names []string
	for _, k := range kindOrder {
		if kinds&k != 0 && (ix.layers == layerKnown || ix.fits[k][tail] > 0) {
			names = append(names, kindLabel[k])
		}
	}
	if names == nil {
		return "", false
	}
	if closed {
		return c.name, true
	}
	return "열쇠:" + strings.Join(names, "|"), true
}
