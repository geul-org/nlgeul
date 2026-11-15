//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamAgainLifted — in a stretch of words written together the lifts of the rules of a sentence change the ways: Next's Lambda is the sum of the first lambdas of its nodes after the lifts (밥 raised by 5), and Again gives the next ways in the order of the raised lambdas (밥 학생이되었다, which holds the raised 밥, now comes before 밥학생이 되었다), each way's Lambda the sum of its nodes' first lambdas
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStreamAgainLifted(t *testing.T) {
	isBap := func(n rule.Near) bool { return n.Stem == "밥" }
	bap := sentenceRuleOf("밥", 5, isBap, nil, listed("밥", rule.ListSelfStem, []string{"밥"}), nil, "")
	const s = "밥학생이되었다"
	plain := sentenceStream(t, s, false)
	st := sentenceStream(t, s, false, bap)
	was, _ := plain.Next()
	seg, _ := st.Next()
	if segmentLine(seg) != segmentLine(was) || !almost(seg.Lambda, was.Lambda+5) {
		t.Fatal(seg.Lambda, was.Lambda)
	}
	want := []string{"밥[0-1] 학생이되었다[1-7]", "밥학생이[0-4] 되었다[4-7]", "밥학생이되었다[0-7]"}
	if second, _ := plain.Again(was.Region); segmentLine(second) != want[1] {
		t.Fatal("unraised, 밥학생이 되었다 is second", segmentLine(second))
	}
	ways := []Segment{seg}
	for next, ok := st.Again(seg.Region); ok; next, ok = st.Again(seg.Region) {
		ways = append(ways, next)
	}
	if len(ways) != len(want)+1 {
		t.Fatal("ways", len(ways))
	}
	for k, w := range ways {
		if !almost(w.Lambda, lambdaSum(w)) || (k > 0 && (segmentLine(w) != want[k-1] || w.Rank != k+1)) {
			t.Fatal(k, segmentLine(w), w.Lambda, lambdaSum(w))
		}
	}
}
