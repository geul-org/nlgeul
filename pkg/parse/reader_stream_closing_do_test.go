//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestStreamClosingDo — the auxiliary particle 도 right after the closing 는데, ㄴ데, 은데 or 던데 written onto it is read first as one word with the closing (가는데도, 어려운데도, 좋은데도, 먹던데도), not as a modifier then the dependent noun 데 and 도; 는 after it keeps its old reading, and a word of its own (그런데도) is not cut
package parse

import (
	"strings"
	"testing"
)

func TestStreamClosingDo(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["가"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["좋"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	r.ix.lex["먹"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	for in, want := range map[string]string{
		"가는데도": "가[열쇠:동사]+는데[맺음:는데]+도[보조사:도]",
		"좋은데도": "좋[열쇠:형용사]+은데[맺음:은데]+도[보조사:도]",
		"먹던데도": "먹[열쇠:동사]+던데[맺음:던데]+도[보조사:도]",
	} {
		if got := flatNodes(r, in, true); len(got) != 1 || !strings.HasSuffix(got[0], "} "+want) {
			t.Errorf("%s: want %s first\n%v", in, want, got)
		}
	}
}
