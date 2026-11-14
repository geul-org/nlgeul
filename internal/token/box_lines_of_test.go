//ff:func feature=token type=reader control=sequence
//ff:what TestBoxLinesOf — a stream drawn: the stretch of a spaced sentence in order, the best way of each in "1등 길:" with the sum of the lambdas; a stretch nothing is known of is its letters, a stretch in the path as it is; the trace lines of a unit come before its boxes and those of units with no stretch after them all; head cuts the ways asked of the stream; no stretch gives no lines
package token

import (
	"strings"
	"testing"
)

func TestBoxLinesOf(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가 방에 들어가신다"
	got := boxLinesOf(r.Stream(s), lettersOf(s), 2)
	last := got[len(got)-1]
	if !strings.HasPrefix(last, "1등 길: 아버지가 │ 방에 │ 들어가신다  (λ 합 ") || !strings.HasSuffix(last, ")") {
		t.Fatal(last)
	}
	if n := strings.Count(strings.Join(got, "\n"), "어절경합"); n != 2 {
		t.Fatal(n)
	}
	st := r.Stream(s)
	boxLinesOf(st, lettersOf(s), 1)
	if len(st.sent) != 3 || len(st.sent[0].out) != 1 || len(st.sent[1].out) != 1 || len(st.sent[2].out) != 1 {
		t.Fatal("ways made beyond head")
	}

	empty := &Stream{queue: []*region{newRegion(Contest{Text: "가나"}, 1, nil)}}
	got = boxLinesOf(empty, []rune("다가나라"), 0)
	if len(got) != 2 || got[0] != "가나" || got[1] != "1등 길: 가나  (λ 합 0.00)" {
		t.Fatal(got)
	}
	if got := boxLinesOf(&Stream{}, nil, 0); len(got) != 0 {
		t.Fatal(got)
	}

	noted := &Stream{notes: []string{"처음"}, queue: []*region{newRegion(Contest{Text: "가나"}, 0, nil)}}
	got = boxLinesOf(noted, []rune("가나"), 0)
	if len(got) != 3 || got[0] != "처음" || got[1] != "가나" {
		t.Fatal(got)
	}
}
