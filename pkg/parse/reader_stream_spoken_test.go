//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamSpoken — end to end with the word table of testdata, a marker written as it is spoken is read with the letters as they stand and the name of the standard marker, after a stem of the word table: the closing 구 for 고 (마시+구), the particle 두 for 도 after a nominal and after a linking closing (마당+두, 작+게+두), the fused 어도 written 두 (켜+(어)두), and 구 for 고 after a quoted closing (마시+(ㄴ)다구); not after a stem the word table does not have, and not without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamSpoken(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"마시구":  "\n마시[열쇠:동사]+구[맺음:고]\n",
		"마당두":  "\n마당[열쇠:체언]+두[보조사:도]\n",
		"작게두":  "\n작[열쇠:형용사]+게[맺음:게]+두[보조사:도]\n",
		"켜두":   "\n켜[열쇠:동사]+(어)두[맺음:어도(합침)]\n",
		"마신다구": "\n마시[열쇠:동사]+(ㄴ)다구[맺음:ㄴ다고]\n",
		"학교에두": "\n학교[열쇠:체언]+에[부사격:에]+두[보조사:도]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{"뷁구": "구[맺음:고]", "뷁두": "두[보조사:도]", "듣두": "듣[열쇠:동사]+두[보조사:도]"} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q after a stem that does not take it\n%s", in, not, got)
		}
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	if got := streamText(bare, "마시구"); strings.Contains(got, "구[맺음:고]") {
		t.Errorf("a written form without a word table\n%s", got)
	}
}
