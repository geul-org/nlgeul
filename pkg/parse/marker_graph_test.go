//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestMarkerGraph — the exceptions a graph gets by its entry: every marker the stem check and, but for a suffix (접미:이), the check of a unit glossed with Han characters; a particle its end and chain order; a slot its order; a lone consonant its base (not with a word table); a piece that begins with 어/아 the rebuttal of a vowel that does not fit; a form with 으 read after a stem's ㄹ (loose) the rebuttal of a consonant that does not fit, which leaves it standing after ㄹ below the same form not loose, and not after another consonant; a form that ends a closed-set word that word; a closing read only with a word table what stands right after it, and so a particle read only with a word table (a shortened particle); a closing of hangulBound the exception that no Hangul syllable stands before it (counted); a graph the case exception table names one exception more, apart from the closed-set words' and with a reason of its own, which turns down the words of its rows (the trace naming the word and the row) and no other word
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraph(t *testing.T) {
	cases := []struct {
		e     entry
		whole []string
		edges int
	}{
		{entry{name: "보조사:은", form: "은", class: classParticle, after: rule.AfterBatchim}, nil, 4},
		{entry{name: "부사격:에", form: "에", class: classParticle}, nil, 4},
		{entry{name: "접미:이", form: "이", class: classParticle, after: rule.AfterBatchim, inner: true}, nil, 4},
		{entry{name: "선어말:었", form: "었", class: classSlot}, nil, 3},
		{entry{name: "꾸밈:ㄴ", class: classClosing, lead: rule.JongN}, nil, 3},
		{entry{name: "맺음:고", form: "고", class: classClosing}, []string{"그리고"}, 3},
		{entry{name: "맺음:ㄴ다", form: "다", class: classClosing, lead: rule.JongN}, nil, 2},
		{entry{name: "맺음:재(자고 해)", form: "재", class: classClosing, wide: true}, nil, 3},
		{entry{name: "부사격:선(에서는)", form: "선", class: classParticle, wide: true}, nil, 4},
		{entry{name: "맺음:어서", form: "어서", class: classClosing, head: headWhole}, nil, 3},
		{entry{name: "꾸밈:은", form: "은", class: classClosing, after: rule.AfterLiquid, loose: true}, nil, 5},
		{entry{name: "꾸밈:을", form: "을", class: classClosing, after: rule.AfterLiquid}, nil, 4},
	}
	for _, c := range cases {
		e := c.e
		if got := markerGraph(&e, c.whole, nil, false); got != c.edges || e.g == nil {
			t.Errorf("%s: %d", e.name, got)
		}
	}
	lone := entry{name: "꾸밈:ㄴ", class: classClosing, lead: rule.JongN}
	if got := markerGraph(&lone, nil, nil, true); got != 2 {
		t.Fatal("with a word table a lone consonant's base is not the graph's to ask", got)
	}
	loose := entry{name: "꾸밈:은", form: "은", class: classClosing, after: rule.AfterLiquid, loose: true}
	tight := loose
	tight.loose = false
	markerGraph(&loose, nil, nil, true)
	markerGraph(&tight, nil, nil, true)
	liquid := rule.Piece{Form: "은", HasStem: true, Prev: '줄', Next: -1}
	v, _ := verdictOf(loose.g, liquid, "줄은", false)
	w, _ := verdictOf(tight.g, liquid, "줄은", false)
	if v <= 0 || v >= w {
		t.Fatal("after a stem's ㄹ a loose form stands, rebutted below the same form read by the tables", v, w)
	}
	if v, _ = verdictOf(loose.g, rule.Piece{Form: "은", HasStem: true, Prev: '먹', Next: -1}, "먹은", false); v > 0 {
		t.Fatal("a loose form after another consonant", v)
	}
	rows := []caseException{{Graph: "맺음:소", Word: "담소", GroundKind: groundSeen, Ground: "46/46", Row: data.Row{Line: 2}}}
	plain := entry{name: "맺음:소", form: "소", class: classClosing}
	held, both := plain, plain
	if a, b, c := markerGraph(&plain, nil, nil, true), markerGraph(&held, nil, rows, true), markerGraph(&both, []string{"담소"}, rows, true); b != a+1 || c != a+3 {
		t.Fatal("a graph the case exception table names has one exception more, apart from the closed-set words' (and from what lifts theirs for a closing with a word table)", a, b, c)
	}
	p := rule.Piece{Form: "소", HasStem: true, Prev: '담', Next: -1}
	note := whyCase + "(담소) — 예외 표 2: 관찰 46/46"
	if v, why := verdictOf(held.g, p, "담소", true); v != -1 || !slices.Equal(why, []string{note}) {
		t.Fatal(v, why)
	}
	if v, why := verdictOf(both.g, p, "담소", true); v != -1 || !slices.Equal(why, []string{whyWhole, note}) {
		t.Fatal("two exceptions, two reasons", v, why)
	}
	p.Prev = '적'
	if v, why := verdictOf(held.g, p, "적소", true); v != 1 || len(why) != 0 {
		t.Fatal("a word no row has", v, why)
	}
}
