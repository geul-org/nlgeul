//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestMarkerGraph — the exceptions a graph gets by its entry: every marker the stem check; a particle its end and chain order; a slot its order; a lone consonant its base; a form that ends a closed-set word that word
package token

import "testing"

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
	}
	for _, c := range cases {
		e := c.e
		if got := markerGraph(&e, c.whole); got != c.edges || e.g == nil {
			t.Errorf("%s: %d", e.name, got)
		}
	}
}
