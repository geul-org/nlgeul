//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestReaderStreamTypoContract — 껀 and 껄 right after a modifying closing, written onto it or apart, are the dependent noun 거 written with 꺼 and the particle under it (먹는껄 as 먹+는 거+(를), with the new word 가는껄 a candidate behind it), named as 거 is with the letters it is written with; they are not read so after a word that is not a modifying form (학생껀, where 꺼 is still read with the copula's stem dropped before ㄴ, as it is in 학생꺼다), and not read with no word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamTypoContract(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      string
	}{
		{"가는껄", "\n가[열쇠:동사]+는[꾸밈:는] 거[의존명사:거,오타정정:꺼]+(를)[목적격:를]\n"},
		{"가는껄", "\n가는껄\n"},
		{"가는껀", "\n가[열쇠:동사]+는[꾸밈:는] 거[의존명사:거,오타정정:꺼]+(는)[보조사:는]\n"},
		{"가는 껄", "\n거[의존명사:거,오타정정:꺼]+(를)[목적격:를]\n"},
		{"가는 껀", "\n거[의존명사:거,오타정정:꺼]+(는)[보조사:는]\n"},
		{"가는걸", "\n가[열쇠:동사]+는[꾸밈:는] 거[의존명사:거]+(를)[목적격:를]\n"},
	}
	for _, c := range cases {
		if got := streamText(r, c.sentence); !strings.Contains(got, c.has) {
			t.Errorf("%s: missing %q in\n%s", c.sentence, c.has, got)
		}
	}
	if got := streamText(r, "학생껀"); strings.Contains(got, "오타정정:꺼]+(는)") {
		t.Errorf("after a word that is no modifying form:\n%s", got)
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	if got := streamText(bare, "가는껄"); strings.Contains(got, "오타정정") {
		t.Errorf("no word table:\n%s", got)
	}
}
