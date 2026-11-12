//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestMarkerGraph — the exceptions a graph gets by its entry: every marker the stem check; a particle its end and chain order; a slot its order; a lone consonant its base (not with a word table); a piece that begins with 어/아 the rebuttal of a vowel that does not fit; a form with 으 read after a stem's ㄹ (loose) the rebuttal of a consonant that does not fit, which leaves it standing after ㄹ below the same form not loose, and not after another consonant; a form that ends a closed-set word that word; a closing read only with a word table what stands right after it
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraph(t *testing.T) {
	cases := []struct {
		e     entry
		whole []string
		edges int
	}{
		{entry{name: "보조사:은", form: "은", class: classParticle, after: afterBatchim}, nil, 3},
		{entry{name: "부사격:에", form: "에", class: classParticle}, nil, 2},
		{entry{name: "접미:이", form: "이", class: classParticle, after: afterBatchim, inner: true}, nil, 4},
		{entry{name: "선어말:었", form: "었", class: classSlot}, nil, 2},
		{entry{name: "꾸밈:ㄴ", class: classClosing, lead: jongN}, nil, 2},
		{entry{name: "맺음:고", form: "고", class: classClosing}, []string{"그리고"}, 2},
		{entry{name: "맺음:ㄴ다", form: "다", class: classClosing, lead: jongN}, nil, 1},
		{entry{name: "맺음:재(자고 해)", form: "재", class: classClosing, wide: true}, nil, 2},
		{entry{name: "부사격:선(에서는)", form: "선", class: classParticle, wide: true}, nil, 2},
		{entry{name: "맺음:어서", form: "어서", class: classClosing, head: headWhole}, nil, 2},
		{entry{name: "꾸밈:은", form: "은", class: classClosing, after: rule.AfterLiquid, loose: true}, nil, 3},
		{entry{name: "꾸밈:을", form: "을", class: classClosing, after: rule.AfterLiquid}, nil, 2},
	}
	for _, c := range cases {
		e := c.e
		if got := markerGraph(&e, c.whole, false); got != c.edges || e.g == nil {
			t.Errorf("%s: %d", e.name, got)
		}
	}
	lone := entry{name: "꾸밈:ㄴ", class: classClosing, lead: jongN}
	if got := markerGraph(&lone, nil, true); got != 1 {
		t.Fatal("with a word table a lone consonant's base is not the graph's to ask", got)
	}
	loose := entry{name: "꾸밈:은", form: "은", class: classClosing, after: rule.AfterLiquid, loose: true}
	tight := loose
	tight.loose = false
	markerGraph(&loose, nil, true)
	markerGraph(&tight, nil, true)
	liquid := rule.Piece{Form: "은", HasStem: true, Prev: '줄', Next: -1}
	v, _ := verdictOf(loose.g, liquid, "줄은")
	w, _ := verdictOf(tight.g, liquid, "줄은")
	if v <= 0 || v >= w {
		t.Fatal("after a stem's ㄹ a loose form stands, rebutted below the same form read by the tables", v, w)
	}
	if v, _ = verdictOf(loose.g, rule.Piece{Form: "은", HasStem: true, Prev: '먹', Next: -1}, "먹은"); v > 0 {
		t.Fatal("a loose form after another consonant", v)
	}
}
