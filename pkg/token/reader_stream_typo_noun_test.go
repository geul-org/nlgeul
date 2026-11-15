//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderStreamTypoNoun — end to end with the table of testdata that has nouns written wrong (역활 for 역할, 설겆이, 컨텐츠; 꺼 for 거 is not in the table but in the closed table of this layer, and is named as the dependent noun): a noun written wrong is read as the right form with no marker after it (alone, written together with a noun, before 하, between a slash and a particle, as 꺼 after a modifying ㄹ, a word of its own there), the letters kept, next to the candidates it had; a verb stem written wrong (맞기) is read nowhere without a marker
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamTypoNoun(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
	}{
		{"역활", []string{"\n역할[열쇠:체언,오타정정:역활]\t성립 0.51", "\n역활\t성립 0.33"}, nil},
		{"컨텐츠 역활", []string{"\n콘텐츠[열쇠:체언,오타정정:컨텐츠]\t", "\n역할[열쇠:체언,오타정정:역활]\t"}, nil},
		{"역활분담", []string{"\n역할[열쇠:체언,오타정정:역활] 분담[열쇠:체언]\t", "\n역활분담\t성립 0.33"}, nil},
		{"설겆이하고", []string{"\n설거지[열쇠:체언,오타정정:설겆이] 하[열쇠:동사]+고[맺음:고]\t", "\n설겆이하고\t성립 0.33"}, nil},
		{"입력/역활을", []string{"\n입력/역할[대비:/,오타정정:입력/역활]+을[목적격:을]\t", "\n입력/역활을\t성립 0.33"}, nil},
		{"내꺼", []string{"\n내[열쇠:체언] 거[의존명사:거,오타정정:꺼]\t"}, nil},
		{"할꺼야", []string{"\n하[열쇠:동사]+(ㄹ)[꾸밈:ㄹ]\t", "\n거[의존명사:거,오타정정:꺼]+야[보조사:야]\t"}, []string{"오타정정:꺼]+야[호격:야]"}},
		{"역할", []string{"\n역할[열쇠:체언]\t성립 0.92"}, []string{"오타정정"}},
		{"설거지하고", []string{"\n설거지[열쇠:체언] 하[열쇠:동사]+고[맺음:고]\t"}, []string{"오타정정"}},
		{"입력/역할", []string{"\n입력/역할[대비:/]\t성립 0.92"}, []string{"오타정정"}},
		{"맞기", []string{"\n맞[열쇠:동사]+기[이름:기]\t성립 1.00"}, []string{"오타정정"}},
	}
	for _, c := range cases {
		got := scoredText(r, c.sentence)
		for _, want := range c.has {
			if !strings.Contains(got, want) {
				t.Errorf("%s: missing %q in\n%s", c.sentence, want, got)
			}
		}
		for _, no := range c.lacks {
			if strings.Contains(got, no) {
				t.Errorf("%s: %q in\n%s", c.sentence, no, got)
			}
		}
	}
}
