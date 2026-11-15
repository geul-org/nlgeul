//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestStreamNextUnruled — rules that never raise a candidate leave Next as it is with no rule: for the sentences the stream tests read and those of words-sentence.tsv, with each table and none, every stretch, way, node, candidate order, lambda and trace line (each after its own Next) is the same as with no rule, but for the trace lines of the rules themselves; a rule that waits for every candidate and never stands (it reads on to each predicate) and a rule that speaks of no candidate are both asked
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStreamNextUnruled(t *testing.T) {
	every := func(rule.Near) bool { return true }
	never := func(rule.Near) bool { return false }
	waiting := sentenceRuleOf("기다림", 7, every, every, listed("없음", rule.ListSelfStem, nil), nil, "")
	silent := sentenceRuleOf("말 없음", 7, never, nil, listed("모두", rule.ListSelfStem, nil), listed("모두", rule.ListSelfStem, nil), "늘")
	sentences := append(streamSentences(), "학생이 밥을 먹었다", "학생이 되었다 학생이", "학생이되었다")
	for _, words := range []string{"", "words-stream.tsv", "words.tsv", "words-sentence.tsv"} {
		r := streamReader(t, words)
		for _, s := range sentences {
			plain := streamDump(r.StreamTrace(s))
			ruled := r.StreamTrace(s)
			ruled.rules = []sentenceRule{waiting, silent}
			got := streamDump(ruled)
			if !slices.Equal(withoutLines(got, "기다림", "말 없음"), plain) {
				t.Fatalf("%q %q:\n%q\n%q", words, s, got, plain)
			}
			if len(got) <= len(plain) {
				t.Fatalf("the waiting rule was not asked: %q %q", words, s)
			}
		}
	}
}
