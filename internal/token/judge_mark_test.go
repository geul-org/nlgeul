//ff:func feature=token type=rule control=sequence
//ff:what TestJudgeMark — the readings that stand come back one by one, each with its own verdict; each is noted in the trace, a defeated one with its reason; without a trace nothing is noted
package token

import (
	"strings"
	"testing"
)

func TestJudgeMark(t *testing.T) {
	sc := scanOf("x")
	reads := sc.judgeMark(sc.ix.joins, dotFacts([]string{"task", "py"}, 0), ".")
	if len(reads) != 2 || reads[0].name != "파일명:." || reads[1].name != "도메인:." || reads[0].v != 1 || reads[1].v != 1 || len(*sc.trace) != 4 || !strings.Contains((*sc.trace)[2], "기각: 양쪽이 맞는 글자 열이 아니다") {
		t.Fatal(reads, *sc.trace)
	}
	sc.trace = nil
	if reads = sc.judgeMark(sc.ix.joins, dotFacts([]string{"a", "b"}, 0), "."); len(reads) != 0 {
		t.Fatal(reads)
	}
}
