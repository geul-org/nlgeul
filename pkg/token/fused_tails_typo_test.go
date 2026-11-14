//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestFusedTypo — a closing written wrong that is a final consonant under the syllable before it is a tail whose piece is the right closing in brackets (the letters kept), not the written form
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestFusedTypo(t *testing.T) {
	sc := typoScan("할께", layerCost)
	e := &entry{name: "맺음:ㄹ게", form: "께", class: classClosing, lead: rule.JongL, wrote: "ㄹ께", right: "(ㄹ)게"}
	got := sc.fusedTails(piece{text: "께", names: "맺음:ㄹ게", wrote: "ㄹ께"}, e, 1)
	if len(got) == 0 {
		t.Fatal(got)
	}
	for _, tl := range got {
		if tl.pieces[0].text != "(ㄹ)게" || tl.pieces[0].wrote != "ㄹ께" {
			t.Fatal(tl.pieces)
		}
	}
	whole := &entry{name: "맺음:려고", form: "려고", class: classClosing, lead: rule.JongL, wrote: "ㄹ려고", right: "려고"}
	got = typoScan("갈려고", layerCost).fusedTails(piece{text: "려고", names: "맺음:려고", wrote: "ㄹ려고"}, whole, 1)
	if len(got) == 0 || got[0].pieces[0].text != "려고" || got[0].base != '가' {
		t.Fatal("a closing of whole syllables under a final consonant has no brackets", got)
	}
}
