//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderLinesTypoClosing — end to end with the table of testdata: each closing of the table of closings written wrong is read as the right closing with the letters it is written with kept (할께, 먹을께, 갈껄, 할찌, 할쑤록, 할려면, 먹을려고, 갈려고, 했읍니다, 있슴, 좋구만, 학생이예요, 거에요, 들을께), the right closing's stem kept apart from the stem under a final consonant (갈려고 next to the 려고 after 갈); a closing the right form cannot follow (예쁠려고 after an adjective, 책께 with no final ㄹ) is not read with a stem of the table, nor a verb stem written wrong
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesTypoClosing(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
	}{
		{"할께", []string{"  하[열쇠:동사]+(ㄹ)게[맺음:ㄹ게,오타정정:ㄹ께]\t성립 0.56"}, nil},
		{"할께요", []string{"  하[열쇠:동사]+(ㄹ)게[맺음:ㄹ게,오타정정:ㄹ께]+요[보조사:요]\t"}, nil},
		{"먹을께", []string{"  먹[열쇠:동사]+을게[맺음:을게,오타정정:을께]\t"}, nil},
		{"들을께", []string{"  듣[열쇠:동사]+을게[맺음:을게,오타정정:을께]\t"}, nil},
		{"갈껄", []string{"  가[열쇠:동사]+(ㄹ)걸[맺음:ㄹ걸,오타정정:ㄹ껄]\t"}, nil},
		{"할찌", []string{"  하[열쇠:동사]+(ㄹ)지[맺음:ㄹ지,오타정정:ㄹ찌]\t"}, nil},
		{"할쑤록", []string{"  하[열쇠:동사]+(ㄹ)수록[맺음:ㄹ수록,오타정정:ㄹ쑤록]\t"}, nil},
		{"할려면", []string{"  하[열쇠:동사]+려면[맺음:려면,오타정정:ㄹ려면]\t"}, nil},
		{"먹을려고", []string{"  먹[열쇠:동사]+으려고[맺음:으려고,오타정정:을려고]\t"}, nil},
		{"갈려고", []string{"  가[열쇠:동사]+려고[맺음:려고,오타정정:ㄹ려고]\t"}, []string{"  갈[열쇠:동사]+려고[맺음:려고,오타정정"}},
		{"했읍니다", []string{"  하[열쇠:동사]+(였)[선어말:ㅆ]+습니다[맺음:습니다,오타정정:읍니다]\t"}, nil},
		{"있슴", []string{"  있[보조 용언:있]+음[이름:음,오타정정:슴]\t", "  있슴\t성립 0.33"}, nil},
		{"좋구만", []string{"  좋[열쇠:형용사]+구먼[맺음:구먼,오타정정:구만]\t"}, nil},
		{"학생이예요", []string{"  학생[열쇠:체언]+이[서술격:이]+에요[맺음:에요,오타정정:예요]\t"}, nil},
		{"거에요", []string{"  거[의존명사:거]+예요[맺음:예요,오타정정:에요]\t"}, nil},
		{"책께", nil, []string{"오타정정"}},
		{"예쁠려고", nil, []string{"예쁘[열쇠:형용사]+려고"}},
		{"맞기", nil, []string{"오타정정"}},
	}
	for _, c := range cases {
		got := strings.Join(r.Lines(c.sentence, false), "\n")
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
