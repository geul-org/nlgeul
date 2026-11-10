//ff:func feature=token type=helper control=iteration dimension=1
//ff:what namesOf — the readings a piece's joined names stand for, each with its kind and form apart: 주격:가|보격:가 is two readings; the kinds a word table gives a stem (열쇠:체언|부사) are one reading each of the kind 열쇠; a name with no colon is a kind alone (수, 경로); no names give none
package token

import "strings"

func namesOf(joined string) []Name {
	if joined == "" {
		return nil
	}
	var out []Name
	for _, one := range strings.Split(joined, "|") {
		kind, form, ok := strings.Cut(one, ":")
		if !ok && len(out) > 0 && out[len(out)-1].Kind == keyKind {
			kind, form = keyKind, one
		}
		out = append(out, Name{kind, form})
	}
	return out
}
