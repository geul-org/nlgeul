//ff:func feature=token type=reader control=sequence
//ff:what TestEntryForm — two candidates of one particle form share a key; a closing of the same letters, and a fused piece, have their own
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEntryForm(t *testing.T) {
	a := entryForm(&entry{name: "주격:이", form: "이", class: classParticle})
	b := entryForm(&entry{name: "보격:이", form: "이", class: classParticle})
	c := entryForm(&entry{name: "맺음:다", form: "다", class: classClosing})
	d := entryForm(&entry{name: "맺음:ㄴ다", form: "다", lead: rule.JongN, class: classClosing})
	if a != b || c == d || a == c {
		t.Fatal(a, b, c, d)
	}
}
