//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestLinkBeforeAuxRule — the rule on its own, a candidate judged with the words around it given: it speaks of the candidates that end in 어/아/여 or their fused form (not in 고, not a noun) and waits for the words after them, its name and its bonus are bonusLinkBeforeAux, it stands (통과) when the stem of the word right after is one of the auxiliaries given (있, 주), and it falls with the reason 바로 뒤 어절의 줄기가 보조 용언이 아니다 when that stem is another or no word is after
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLinkBeforeAuxRule(t *testing.T) {
	words, _ := linkBeforeWords(t, "남아 먹다", false)
	link := auxLinkedPick(t, words[0], "남[열쇠:동사]+아[맺음:아]")
	noun := auxLinkedPick(t, words[0], "남아[열쇠:체언]")
	r := linkBeforeAuxRule([]string{"있", "주"})
	if r.name != ruleLinkBeforeAux || r.bonus != bonusLinkBeforeAux || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for near, want := range map[rule.Near]bool{
		nearOf(link):              true,
		{Last: "맺음:어(합침)"}:       true,
		{Last: "맺음:여"}:           true,
		{Last: "맺음:고"}:           false,
		nearOf(noun):              false,
		{Stem: "집", Last: "부사격:에"}: false,
	} {
		if r.about(near) != want || r.waits(near) != want {
			t.Fatalf("%+v: about %v", near, !want)
		}
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string {
		return "  · 남[열쇠:동사]+아[맺음:아]\t" + ruleLinkBeforeAux + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		after []rule.Near
		add   float64
		note  string
	}{
		{[]rule.Near{{Stem: "있", Kind: "열쇠:형용사"}}, bonusLinkBeforeAux, line("1.00", "통과")},
		{[]rule.Near{{Stem: "주", Kind: "열쇠:동사"}, {Stem: "먹"}}, bonusLinkBeforeAux, line("1.00", "통과")},
		{[]rule.Near{{Stem: "먹", Kind: "열쇠:동사"}, {Stem: "있"}}, 0, line("-1.00", "기각: 바로 뒤 어절의 줄기가 보조 용언이 아니다")},
		{nil, 0, line("-1.00", "기각: 바로 뒤 어절의 줄기가 보조 용언이 아니다")},
	} {
		add, notes := st.judgeCandidate(link, rule.Around{Self: nearOf(link), After: c.after})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.after, add, notes)
		}
	}
}
