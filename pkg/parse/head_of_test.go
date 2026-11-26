//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestHeadOf — what each kind of slot or closing begins with: fused 어 and ㅆ; fused ㄴ, ㄹ, ㅁ, ㅂ, which take 으 where the closings of the index have the form with 으 (ㄴ데 as 은데, ㅂ시다 as 읍시다; not ㄴ다, ㄴ답니다, a shortened closing or ㅂ니다), and the modifying ㄴ is told as that closing (headAdn); a syllable 어/았; a form with 으; a first consonant with or without 으; 네; 건대, before which 하 drops wholly (not 건만); a particle, a piece with no letters and a form read after a stem's ㄹ against the tables (loose 은, 을) nothing
package parse

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
		{entry{class: classClosing, lead: rule.LeadVowel, form: "서"}, headFused},
		{entry{class: classSlot, lead: rule.JongSS}, headFused},
		{entry{class: classClosing, lead: rule.JongN, form: "다"}, headN},
		{entry{class: classClosing, lead: rule.JongN, form: "대", wide: true}, headN},
		{entry{class: classClosing, lead: rule.JongN, form: "데"}, headN | headEuOpen},
		{entry{class: classClosing, lead: rule.JongN, name: "꾸밈:ㄴ"}, headN | headEuOpen | headAdn},
		{entry{class: classClosing, lead: rule.JongN, form: "답니다"}, headN},
		{entry{class: classClosing, lead: rule.JongL, form: "까"}, headL | headEuOpen},
		{entry{class: classClosing, lead: rule.JongM}, headEuOpen},
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
		{entry{class: classClosing, form: "건대", name: "맺음:건대"}, headHaW},
		{entry{class: classClosing, form: "건만", name: "맺음:건만"}, 0},
		{entry{class: classClosing, form: "은", after: rule.AfterLiquid, loose: true}, 0},
		{entry{class: classClosing, form: "을", after: rule.AfterLiquid, loose: true}, 0},
	}
	ix, _ := newIndex(Options{})
	for _, c := range cases {
		e := c.e
		if got := ix.headOf(&e); got != c.want {
			t.Errorf("%+v: %d, want %d", c.e, got, c.want)
		}
	}
}
