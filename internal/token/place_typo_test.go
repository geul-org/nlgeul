//ff:func feature=token type=reader control=sequence
//ff:what TestPlaceTypo — the facts of the place of an entry that reads a closing written wrong tell its graph that it is a misspelt form; the facts of an entry written right do not
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPlaceTypo(t *testing.T) {
	sc := scanOf("먹을께")
	wrong := &entry{name: "맺음:을게", form: "을께", class: classClosing, wrote: "을께", right: "을게"}
	right := &entry{name: "맺음:을게", form: "을게", class: classClosing}
	if _, p, ok := sc.place(wrong, 3, false); !ok || !p.Misspelt {
		t.Fatal(p, ok)
	}
	if _, p, ok := scanOf("먹을게").place(right, 3, false); !ok || p.Misspelt {
		t.Fatal(p, ok)
	}
	lone := scanOf("할께")
	jong := &entry{name: "맺음:ㄹ게", form: "께", class: classClosing, lead: rule.JongL, wrote: "ㄹ께", right: "(ㄹ)게"}
	if start, p, ok := lone.place(jong, 2, false); !ok || start != 1 || p.Base != "하" || !p.Misspelt {
		t.Fatal(start, p, ok)
	}
}
