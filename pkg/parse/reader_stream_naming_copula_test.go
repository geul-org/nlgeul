//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNamingCopula — end to end with the word table of testdata, the copula's stem is read right after a naming closing of a predicate (먹음이다, 모름이라, 먹었음일, 모름인가), as a particle chain is; a word that only ends in such letters has no such reading (마당이다), nor has the naming 기 (먹기이다) or a stem of one syllable (됨이다), and nothing is read so without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNamingCopula(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, _ := New(Options{})
	for in, want := range map[string]string{
		"먹음이다": "먹[열쇠:동사]+음[이름:음]+이[서술격:이]+다[맺음:다]",
		"모름이라": "모르[열쇠:동사]+(ㅁ)[이름:ㅁ]+이[서술격:이]+라[맺음:라]",
		"먹었음일": "먹[열쇠:동사]+었[선어말:었]+음[이름:음]+이[서술격:이]+(ㄹ)[꾸밈:ㄹ]",
		"모름인가": "모르[열쇠:동사]+(ㅁ)[이름:ㅁ]+이[서술격:이]+(ㄴ)가[맺음:ㄴ가]",
	} {
		if got := streamText(r, in); !strings.Contains(got, "\n"+want+"\n") {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
		if got := streamText(bare, in); strings.Contains(got, "]+이[서술격:이]+") && strings.Contains(got, "[이름:") && strings.Contains(got, want[strings.Index(want, "+이[서술격"):]) && strings.Contains(got, want[strings.Index(want, "[이름:"):]) {
			t.Errorf("%s without a word table\n%s", in, got)
		}
	}
	for _, in := range []string{"마당이다", "먹기이다", "됨이다"} {
		if got := streamText(r, in); strings.Contains(got, "]+이[서술격:이]+다") && strings.Contains(got, "[이름:") {
			t.Errorf("%s: the copula read after a naming closing\n%s", in, got)
		}
	}
}
