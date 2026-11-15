//ff:func feature=token type=helper control=sequence
//ff:what aloneFree — is a candidate's stem of a part of speech a word alone in a pair of brackets may be: not a bound noun or a determiner (aloneBarred), whether the name gives the part of speech before its colon (의존명사:시, 관형사:이) or after the key's (열쇠:관형사); a candidate with no stem or a new word is free
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func aloneFree(n rule.Near) bool {
	kind, word, _ := strings.Cut(n.Kind, ":")
	if kind == keyKind {
		kind = word
	}
	return !slices.Contains(aloneBarred, kind)
}
