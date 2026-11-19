//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestSentenceCases — the rows of the case exception table under a rule of a sentence's name defeat its claim for a candidate whose word is one of their words (the trace adds the row), not for another word; with no row under the name nothing is put on the graph
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSentenceCases(t *testing.T) {
	old := caseExceptions
	t.Cleanup(func() { caseExceptions = old })
	names := []string{"그래프", "낱말", "바른 읽기", "세기", "근거 종류", "근거"}
	caseExceptions = map[string][]caseException{"문장:시험 반박": {{Graph: "문장:시험 반박", Word: "각", Row: data.Row{File: "c.tsv", Line: 2, Names: names, Cells: []string{"문장:시험 반박", "각", "-", "지움", "관찰", "시험"}}}}}
	g, _ := sentenceRebuttal("문장:시험 반박", rule.Listed, listed("줄기", rule.ListSelfStem, []string{"각", "새"}))
	for text, want := range map[string]bool{"각": false, "새": true} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, rule.Around{Self: rule.Near{Stem: text, Text: text}})
		v, why := evaluate(g, ctx, true)
		if (v > 0) != want {
			t.Fatal(text, v)
		}
		if !want && (len(why) != 1 || !strings.HasPrefix(why[0], whyCase) || !strings.Contains(why[0], "각")) {
			t.Fatal("the row is noted", why)
		}
	}
	plain := toulmin.NewGraph("문장:없음")
	claim := plain.Rule(rule.Listed).With(listed("줄기", rule.ListSelfStem, []string{"각"})).With(&rule.Claim{Name: "문장:없음", Concludes: rebutLabel})
	sentenceCases(plain, claim, "문장:없음")
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyAround, rule.Around{Self: rule.Near{Stem: "각", Text: "각"}})
	if v, why := evaluate(plain, ctx, true); v <= 0 || len(why) != 0 {
		t.Fatal("no row, no exception", v, why)
	}
}
