//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestHeadOf — what each kind of slot or closing begins with: fused 어 and ㅆ; fused ㄴ (ㄴ다 and a shortened closing take no 으), ㄹ, ㅁ, ㅂ (ㅂ시다 takes 으); a syllable 어/았; a form with 으; a first consonant with or without 으; 네; a particle and a piece with no letters nothing
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHeadOf(t *testing.T) {
	cases := []struct {
		e    entry
		want int
	}{
		{entry{class: classParticle, form: "는"}, 0},
		{entry{class: classClosing, lead: leadVowel, form: "서"}, headFused},
		{entry{class: classSlot, lead: jongSS}, headFused},
		{entry{class: classClosing, lead: jongN, form: "다"}, headN},
		{entry{class: classClosing, lead: jongN, form: "대", wide: true}, headN},
		{entry{class: classClosing, lead: jongN, form: "데"}, headN | headEuOpen},
		{entry{class: classClosing, lead: rule.JongL, form: "까"}, headL | headEuOpen},
		{entry{class: classClosing, lead: jongM}, headEuOpen},
		{entry{class: classClosing, lead: rule.JongB, form: "시다"}, headB | headEuOpen},
		{entry{class: classClosing, lead: rule.JongB, form: "니다"}, headB},
		{entry{class: classClosing}, 0},
		{entry{class: classClosing, form: "어서"}, headWhole},
		{entry{class: classSlot, form: "았"}, headWhole},
		{entry{class: classClosing, form: "으면"}, headEuForm},
		{entry{class: classClosing, form: "은"}, headEuForm},
		{entry{class: classClosing, form: "면", eu: true}, headEuOpen},
		{entry{class: classClosing, form: "니", name: "맺음:니"}, headN | headEuOpen},
		{entry{class: classSlot, form: "시", name: honorName}, headS | headEuOpen},
		{entry{class: classClosing, form: "는"}, headN},
		{entry{class: classClosing, form: "네"}, headN | headNe},
		{entry{class: classClosing, form: "고"}, 0},
	}
	for _, c := range cases {
		e := c.e
		if got := headOf(&e); got != c.want {
			t.Errorf("%+v: %d, want %d", c.e, got, c.want)
		}
	}
}
