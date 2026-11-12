//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAttached — a unit that is a particle chain (로, 에서, 까지는) or a tail of the copula (인데, 이다) is read as markers alone: one word with no stem part, attaching to the unit before, printed with a leading "+", its verdict no higher than a new word's; 게 alone is 거 with 이 under it; a closing alone (다, 고) and a unit with a stem (곰이) are not read so; the scan is no longer attaching afterwards; end to end a unit that is also a word (이) keeps its word reading next to the attached one
package token

import (
	"strings"
	"testing"
)

func TestAttached(t *testing.T) {
	for in, want := range map[string]string{
		"로":   "+로[부사격:로]",
		"에서":  "+에서[주격:에서|부사격:에서]",
		"까지는": "+까지[보조사:까지]+는[보조사:는]",
		"인데":  "+이[서술격:이]+(ㄴ)데[맺음:ㄴ데]",
		"이다":  "+이[서술격:이]+다[맺음:다]",
	} {
		sc := scanOf(in)
		got := sc.attached()
		if len(got) != 1 || got[0].sig != want || !got[0].last.word.Attach || got[0].last.word.Parts[0].Stem || got[0].verdict > sc.ix.unknown || got[0].last.end != len([]rune(in)) || sc.attach {
			t.Errorf("%s: %v", in, got)
		}
	}
	ge := scanOf("게")
	if got := ge.attached(); len(got) != 1 || got[0].sig != "거[의존명사:거]+(이)[주격:이]" || got[0].last.word.Attach || ge.attach {
		t.Fatal("게 after a modifying form", got)
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
	if got := strings.Join(r.Lines("곰 이", false), "\n"); !strings.Contains(got, "이[관형사:이]") || !strings.Contains(got, "+이[주격:이|보격:이]") {
		t.Fatal("이 keeps its word reading next to the attached one\n" + got)
	}
	if got := strings.Join(r.Lines("곰 만", false), "\n"); !strings.Contains(got, "+만[보조사:만]") || !strings.Contains(got, "  만\t") {
		t.Fatal("만 keeps its reading as a word next to the attached one\n" + got)
	}
}
