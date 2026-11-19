//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosingNames — no rows give no names; a row of an 어/아 closing gives the four names of its entries (어, 아, 여 and the fused one), one that takes 으 two (면서, 으면서), the gated naming closing 이름:기; connNames, mustNames, auxClosings (the quoting 다고, ㄴ다고, 는다고, 라고, 자고, 냐고 and 느냐고, then 다가, 고서, 려고 and 으려고, then 면 and 으면, then 지만, 건만, 거니, 다면, ㄴ다면, 는다면, the four of 어다가, 려다가 and 으려다가 at its end) and nounHeads are so made, in the order of their rows, and each of their names is that of a closing entry of the index (with a word table) under the key of its form
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestClosingNames(t *testing.T) {
	if got := closingNames(); got != nil {
		t.Fatal("no rows", got)
	}
	vowel := []string{"맺음:어", "맺음:아", "맺음:여", "맺음:어(합침)"}
	if !slices.Equal(connNames, vowel) || !slices.Equal(closingNames(parse.Marker{Lead: rule.LeadVowel}), vowel) {
		t.Fatal(connNames)
	}
	if want := []string{"맺음:어야", "맺음:아야", "맺음:여야", "맺음:어야(합침)"}; !slices.Equal(mustNames, want) {
		t.Fatal(mustNames)
	}
	want := slices.Concat([]string{"맺음:지", "맺음:고", "맺음:게"}, vowel, []string{"맺음:어서", "맺음:아서", "맺음:여서", "맺음:어서(합침)", "맺음:면서", "맺음:으면서", "맺음:다고", "맺음:ㄴ다고", "맺음:는다고", "맺음:라고", "맺음:자고", "맺음:냐고", "맺음:느냐고", "맺음:다가", "맺음:고서", "맺음:려고", "맺음:으려고", "맺음:면", "맺음:으면", "맺음:지만", "맺음:건만", "맺음:거니", "맺음:다면", "맺음:ㄴ다면", "맺음:는다면", "맺음:어다가", "맺음:아다가", "맺음:여다가", "맺음:어다가(합침)", "맺음:려다가", "맺음:으려다가"})
	if !slices.Equal(auxClosings, want) {
		t.Fatal(auxClosings)
	}
	if !slices.Equal(nounHeads, []string{"이름:기"}) {
		t.Fatal(nounHeads)
	}
	rows := []parse.Marker{{Lead: rule.LeadVowel}, {Form: "야", Lead: rule.LeadVowel}, {Form: "지"}, {Form: "고"}, {Form: "게"}, {Lead: rule.LeadVowel}, {Form: "서", Lead: rule.LeadVowel}, {Form: "면서", Eu: true}, {Form: "다고"}, {Form: "다고", Lead: rule.JongN}, {Form: "는다고", After: rule.AfterBatchim}, {Form: "라고"}, {Form: "자고"}, {Form: "냐고"}, {Form: "느냐고"}, {Form: "다가"}, {Form: "고서"}, {Form: "려고", Eu: true}, {Form: "면", Eu: true}, {Form: "지만"}, {Form: "건만"}, {Form: "거니"}, {Form: "다면"}, {Form: "다면", Lead: rule.JongN}, {Form: "는다면", After: rule.AfterBatchim}, {Form: "다가", Lead: rule.LeadVowel}, {Form: "려다가", Eu: true}}
	var entries []*entry
	for _, m := range append(rows, namingRows(gatedClosings)...) {
		entries = append(entries, closingEntries(m)...)
	}
	names := slices.Concat(connNames, mustNames, auxClosings, nounHeads)
	if len(entries) != len(names) {
		t.Fatal(len(entries), names)
	}
	ix := wordsScan("x", layerCost).ix
	for i, e := range entries {
		if got := ix.named(entryKey(e), names[i]); got.g == nil || got.class != classClosing || got.form != e.form {
			t.Errorf("%s: %+v", names[i], got)
		}
	}
}
