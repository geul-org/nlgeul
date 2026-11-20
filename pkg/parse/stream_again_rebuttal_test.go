//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestStreamAgainRebuttal — in a stretch of words written together rules of rebuttal change the ways: 밥 rebutted twice stands with its verdict (that of a key of one syllable) multiplied twice by 2/(1+forceUngrammatical)-1 wherever it is a word, so that the way 밥 학생이 되었다 goes behind 밥학생이 되었다; Next and Again give every way the stretch has (none is taken away) in the order of the lambdas after the rebuttal, each way's Lambda the sum of its nodes' first lambdas
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStreamAgainRebuttal(t *testing.T) {
	const s = "밥학생이되었다"
	bap := sentenceRuleOf("밥 반박", 0, func(n rule.Near) bool { return n.Stem == "밥" }, nil, listed("밥", rule.ListSelfStem, []string{"밥"}), nil, "")
	bap.force = forceUngrammatical
	again := bap
	again.name = "밥 반박 둘째"
	count := func(st *Stream) []Segment {
		seg, _ := st.Next()
		ways := []Segment{seg}
		for next, ok := st.Again(seg.Region); ok; next, ok = st.Again(seg.Region) {
			ways = append(ways, next)
		}
		return ways
	}
	plain := count(sentenceStream(t, s, false))
	down := count(sentenceStream(t, s, false, bap, again))
	if len(down) != len(plain) || segmentLine(plain[0]) != "밥[0-1] 학생이[1-4] 되었다[4-7]" || segmentLine(down[0]) != "밥학생이[0-4] 되었다[4-7]" {
		t.Fatal("the way of 밥 goes behind", len(down), len(plain), segmentLine(down[0]))
	}
	for k, w := range down {
		if !almost(w.Lambda, lambdaSum(w)) || w.Rank != k+1 || (k > 0 && w.Lambda > down[k-1].Lambda) {
			t.Fatal(k, segmentLine(w), w.Lambda)
		}
		if w.Nodes[0].Contest.Text == "밥" && w.Nodes[0].Contest.Branches[0].Verdict > (2/(1+forceShort)-1)*(2/(1+forceUngrammatical)-1)*(2/(1+forceUngrammatical)-1)+1e-9 {
			t.Fatal("밥 is rebutted", w.Nodes[0].Contest.Branches[0])
		}
	}
}
