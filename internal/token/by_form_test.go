//ff:func feature=token type=reader control=sequence
//ff:what TestByForm — the emphatic 가 is kept apart from the case particle 가; the readings of one form become one piece with the kinds of each; a suffix of the same form stays apart
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestByForm(t *testing.T) {
	sc := scanOf("곰이요")
	pieces, from, _ := sc.readGroup(classParticle, 2, rule.LinkYo, false)
	forms, merged, kinds := byForm(pieces, from)
	if len(forms) != 2 || forms[0] != "이·접미" || forms[1] != "이" || merged["이"].names != "주격:이|보격:이" || len(kinds["이"]) != 1 || merged["이·접미"].names != "접미:이" {
		t.Fatal(forms, merged)
	}
	two := scanOf("집에서")
	pieces, from, _ = two.readGroup(classParticle, 3, -1, false)
	_, merged, kinds = byForm(pieces, from)
	if merged["에서"].names != "주격:에서|부사격:에서" || len(kinds["에서"]) != 2 {
		t.Fatal(merged, kinds)
	}
	emph := scanOf("집에서가")
	pieces, from, _ = emph.readGroup(classParticle, 4, -1, false)
	if forms, merged, _ = byForm(pieces, from); len(forms) != 2 || merged["가·강조"].names != emphName || merged["가"].names != "주격:가|보격:가" {
		t.Fatal("the emphatic 가 apart from the case particle", forms, merged)
	}
}
