//ff:func feature=token type=helper control=sequence
//ff:what firstName — the name of what stands right after the stem of a predicate tail, for a rule that asks what follows the stem: the tail's closing when the tail begins with it (whole, or written out as its parts: 지 of 죠 and of 잖), else the first piece's own name (a slot, the link vowel, the 어/아 form before an auxiliary)
package token

import (
	"slices"
	"strings"
)

func firstName(t tail) string {
	first := t.pieces[0].names
	if t.close == nil {
		return first
	}
	parts := splits[t.close.name]
	if slices.Contains(strings.Split(first, "|"), t.close.name) || (len(parts) > 0 && parts[0].name == first) {
		return t.close.name
	}
	return first
}
