//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAttached — a unit that is a particle chain (로, 에서, 까지는) or a tail of the copula (인데, 이다) is read as markers alone: one word with no stem part, attaching to the unit before, printed with a leading "+", its verdict no higher than a new word's; 게 alone is 거 with 이 under it; 테니까 and 텐데 are 터 with the copula's stem (이) at the unit's head before the rest as a predicate tail, which do not attach either; 건요 and 게요 (거 and a particle with more after it) are not read so; a closing alone (다, 고) and a unit with a stem (곰이) are not read so; the scan is no longer attaching afterwards; end to end a unit that is also a word (이) keeps its word candidate next to the attached one
package token

import (
	"strings"
	"testing"
)

func TestAttached(t *testing.T) {
	for in, wants := range map[string][]string{
		"로":   {"+로[부사격:로]"},
		"에서":  {"+에서[주격:에서]", "+에서[부사격:에서]"},
		"까지는": {"+까지[보조사:까지]+는[보조사:는]"},
		"인데":  {"+이[서술격:이]+(ㄴ)데[맺음:ㄴ데]"},
		"이다":  {"+이[서술격:이]+다[맺음:다]"},
	} {
		checkAttached(t, in, wants)
	}
	ge := scanOf("게")
	if got := ge.attached(); len(got) != 1 || got[0].sig != "거[의존명사:거]+(이)[주격:이]" || got[0].last.word.Attach || ge.attach {
		t.Fatal("게 after a modifying form", got)
	}
	for in, want := range map[string]string{
		"테니까": "터[줄임:테]+(이)[서술격:이]+니까[맺음:니까]",
		"텐데":  "터[줄임:테]+(이)[서술격:이]+(ㄴ)데[맺음:ㄴ데]",
	} {
		sc := scanOf(in)
		if got := sc.attached(); len(got) != 1 || got[0].sig != want || got[0].last.word.Attach || got[0].last.end != len([]rune(in)) || sc.attach {
			t.Errorf("%s after a modifying form: %v", in, got)
		}
	}
	for _, in := range []string{"건요", "게요"} {
		if got := scanOf(in).attached(); len(got) != 0 {
			t.Errorf("%s: 거 with a particle and more after it: %s", in, got[0].sig)
		}
	}
	if got := scanOf("다").attached(); len(got) != 0 {
		t.Fatal("a closing alone", got[0].sig)
	}
	if got := scanOf("고").attached(); len(got) != 0 {
		t.Fatal("a closing alone", got[0].sig)
	}
	if got := scanOf("곰이").attached(); len(got) != 0 {
		t.Fatal("a unit with a stem", got[0].sig)
	}
	r, _ := New(Options{Words: "testdata/words.tsv"})
	if got := strings.Join(r.Lines("곰 이", false), "\n"); !strings.Contains(got, "이[관형사:이]") || !strings.Contains(got, "+이[주격:이]") || !strings.Contains(got, "+이[보격:이]") {
		t.Fatal("이 keeps its word candidate next to the attached one\n" + got)
	}
	if got := strings.Join(r.Lines("곰 만", false), "\n"); !strings.Contains(got, "+만[보조사:만]") || !strings.Contains(got, "  만\t") {
		t.Fatal("만 keeps its candidate as a word next to the attached one\n" + got)
	}
}
