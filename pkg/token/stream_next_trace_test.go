//ff:func feature=token type=reader control=sequence
//ff:what TestStreamNextTrace — with a trace each rule of a sentence asked of a candidate writes one line, and Notes gives it right after the Next of that word's stretch, after the lines of reading it: the lines are those with no rule plus the rule's line (통과 for 보격 이 before 되었다, 기각 with the reason before 밥을), and the stretches read together with it keep their lines (and the rule lines of their own words) for their own Next; nothing is left for the last Next
package token

import (
	"slices"
	"testing"
)

func TestStreamNextTrace(t *testing.T) {
	plain := sentenceStream(t, "학생이 되었다", true)
	st := sentenceStream(t, "학생이 되었다", true, complementRule(), afterComplementRule())
	plain.Next()
	st.Next()
	want := append(plain.Notes(), "  · 학생[열쇠:체언]+이[보격:이]\t보격 뒤 되\t1.00\t통과")
	if got := st.Notes(); !slices.Equal(got, want) {
		t.Fatalf("%q\n%q", got, want)
	}
	plain.Next()
	st.Next()
	want = append(plain.Notes(), "  · 되[열쇠:동사]+었[선어말:었]+다[맺음:다]\t되 앞 보격\t1.00\t통과")
	if got := st.Notes(); !slices.Equal(got, want) {
		t.Fatalf("the lines of 되었다 wait for its own Next: %q\n%q", got, want)
	}
	if _, ok := st.Next(); ok || st.Notes() != nil {
		t.Fatal("nothing left")
	}

	fell := sentenceStream(t, "학생이 밥을 먹었다", true, complementRule())
	unruled := sentenceStream(t, "학생이 밥을 먹었다", true)
	fell.Next()
	unruled.Next()
	want = append(unruled.Notes(), "  · 학생[열쇠:체언]+이[보격:이]\t보격 뒤 되\t-1.00\t기각: 뒤 어절이 되가 아니다")
	if got := fell.Notes(); !slices.Equal(got, want) {
		t.Fatalf("%q\n%q", got, want)
	}
	fell.Next()
	unruled.Next()
	if got, want := fell.Notes(), unruled.Notes(); len(got) == 0 || !slices.Equal(got, want) {
		t.Fatalf("밥을 keeps its own lines: %q\n%q", got, want)
	}
}
