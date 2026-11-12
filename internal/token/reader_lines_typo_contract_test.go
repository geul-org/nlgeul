//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderLinesTypoContract — 껀 and 껄 right after a modifying closing, written onto it or apart, are the dependent noun 거 written with 꺼 and the particle under it (먹는껄 as 먹+는 거+(를)), named as 거 is with the letters it is written with; they are not read so after a word that is not a modifying form (학생껀), and not read with no word table
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesTypoContract(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      string
	}{
		{"가는껄", "  가[열쇠:동사]+는[꾸밈:는] 거[의존명사:거,오타정정:꺼]+(를)[목적격:를]\t"},
		{"가는껀", "  가[열쇠:동사]+는[꾸밈:는] 거[의존명사:거,오타정정:꺼]+(는)[보조사:는]\t"},
		{"가는 껄", "  거[의존명사:거,오타정정:꺼]+(를)[목적격:를]\t"},
		{"가는 껀", "  거[의존명사:거,오타정정:꺼]+(는)[보조사:는]\t"},
		{"가는걸", "  가[열쇠:동사]+는[꾸밈:는] 거[의존명사:거]+(를)[목적격:를]\t"},
	}
	for _, c := range cases {
		if got := strings.Join(r.Lines(c.sentence, false), "\n"); !strings.Contains(got, c.has) {
			t.Errorf("%s: missing %q in\n%s", c.sentence, c.has, got)
		}
	}
	if got := strings.Join(r.Lines("학생껀", false), "\n"); strings.Contains(got, "오타정정") {
		t.Errorf("after a word that is no modifying form:\n%s", got)
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(bare.Lines("가는껄", false), "\n"); strings.Contains(got, "오타정정") {
		t.Errorf("no word table:\n%s", got)
	}
}
