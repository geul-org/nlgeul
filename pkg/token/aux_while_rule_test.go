//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAuxWhileRule — the rule on its own, a candidate judged with the words around it given: it speaks only of the auxiliary 보 (not the main verb 보, not the auxiliary 있), its name and its bonus are bonusSentence and it does not wait; 보니 as the auxiliary stands (통과) after a word that ends in 다 or 다가 and falls with the reason 앞 어절이 다·다가로 끝나지 않는다 after a linking form or no word; 보아 as the auxiliary, which has none of 니, 니까, 면, gains nothing after 다
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAuxWhileRule(t *testing.T) {
	words, _ := linkBeforeWords(t, "먹다 보니", false)
	aux := auxLinkedPick(t, words[1], "보[보조 용언:보]+니[맺음:니]")
	main := auxLinkedPick(t, words[1], "보[열쇠:동사]+니[맺음:니]")
	r := auxWhileRule()
	if r.name != ruleAuxWhile || r.bonus != bonusSentence || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	if !r.about(nearOf(aux)) || r.about(nearOf(main)) || r.about(rule.Near{Stem: "있", Kind: "보조 용언:있"}) {
		t.Fatal("it speaks of the auxiliary 보 only")
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string {
		return "  · 보[보조 용언:보]+니[맺음:니]\t" + ruleAuxWhile + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		before []rule.Near
		add    float64
		note   string
	}{
		{[]rule.Near{{Stem: "먹", Last: "맺음:다"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Stem: "먹", Last: "맺음:다가"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Stem: "먹", Last: "맺음:어"}}, 0, line("-1.00", "기각: 앞 어절이 다·다가로 끝나지 않는다")},
		{nil, 0, line("-1.00", "기각: 앞 어절이 다·다가로 끝나지 않는다")},
	} {
		add, notes := st.judgeCandidate(aux, rule.Around{Self: nearOf(aux), Before: c.before})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.before, add, notes)
		}
	}
	words, _ = linkBeforeWords(t, "먹다 보아", false)
	bare := auxLinkedPick(t, words[1], "보[보조 용언:보]+아[맺음:아]")
	if add, _ := st.judgeCandidate(bare, rule.Around{Self: nearOf(bare), Before: []rule.Near{{Last: "맺음:다"}}}); add != 0 {
		t.Fatal("보아 gains after 다", add)
	}
}
