//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderStreamTypoKept — the candidates a form written right had are all still there: 할게, 먹을게 (but for the new word 할 or 먹을 before the closing 게, which is the modifying form and 게 written together), 가려고, 했습니다, 있음, 좋구먼, 학생이에요, 거예요 and the rest (a candidate of a closing written wrong may be added, never taken away), and the candidates of the letters of a closing written wrong as a word of their own (구만, 있슴, 좋구만) stand as they did; 갈려고 keeps the 려고 after 갈
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamTypoKept(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      []string
	}{
		{"할게", []string{"\n하[열쇠:동사]+(ㄹ)게[맺음:ㄹ게]\t성립 1.00", "\n하[열쇠:동사]+(ㄹ)[꾸밈:ㄹ] 거[의존명사:거]+(이)[주격:이]\t"}},
		{"먹을게", []string{"\n먹[열쇠:동사]+을게[맺음:을게]\t성립 1.00", "\n먹[열쇠:동사]+을[꾸밈:을] 거[의존명사:거]+(이)[주격:이]\t"}},
		{"가려고", []string{"\n가[열쇠:동사]+려고[맺음:려고]\t성립 1.00", "\n가려+고[맺음:고]\t"}},
		{"갈려고", []string{"\n갈+려고[맺음:려고]\t", "\n갈려+고[맺음:고]\t", "\n갈려고\t성립 0.33"}},
		{"했습니다", []string{"\n하[열쇠:동사]+(였)[선어말:ㅆ]+습니다[맺음:습니다]\t성립 1.00"}},
		{"있음", []string{"\n있[보조 용언:있]+음[이름:음]\t성립 1.00", "\n있음\t성립 0.33"}},
		{"좋구먼", []string{"\n좋[열쇠:형용사]+구먼[맺음:구먼]\t성립 1.00", "\n좋구먼\t성립 0.33"}},
		{"학생이에요", []string{"\n학생[열쇠:체언]+이[서술격:이]+에요[맺음:에요]\t성립 1.00", "\n학생[열쇠:체언] 이[열쇠:체언]+에[부사격:에]+요[보조사:요]\t", "\n학생[열쇠:체언]+이[접미:이]+에[부사격:에]+요[보조사:요]\t"}},
		{"거예요", []string{"\n거[의존명사:거]+예요[맺음:예요]\t성립 1.00", "\n거예요\t성립 0.33"}},
		{"구만", []string{"\n구+만[보조사:만]\t", "\n구만\t성립 0.33"}},
		{"있슴", []string{"\n있슴\t성립 0.33"}},
		{"좋구만", []string{"\n좋구만\t성립 0.33"}},
		{"역할", []string{"\n역할[열쇠:체언]\t성립 1.00"}},
	}
	for _, c := range cases {
		got := scoredText(r, c.sentence)
		for _, want := range c.has {
			if !strings.Contains(got, want) {
				t.Errorf("%s: missing %q in\n%s", c.sentence, want, got)
			}
		}
	}
}
