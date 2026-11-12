//ff:func feature=token type=helper control=sequence
//ff:what nameOf — the candidate a piece's name stands for, with its kind and form apart: 주격:가 is the kind 주격 and the form 가; a name with no colon is a kind alone (수, 경로); no name gives the empty one
package token

import "strings"

func nameOf(name string) Name {
	kind, form, _ := strings.Cut(name, ":")
	return Name{kind, form}
}
