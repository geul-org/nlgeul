//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamOldYeo — with the word table of testdata, the old spelling 야 of the closing 여 (gated-closings.tsv, after the syllable 하) is read after a predicate's stem of the table that ends in 하 (하야, 듯하야), beside the closing 야 after the copula stem, which stands as before (학교야, 학교이야); it is not read after another stem (가야 has no 가+야) nor without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamOldYeo(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"하야":   "하[열쇠:동사]+야[맺음:야]",
		"듯하야":  "듯하[열쇠:형용사]+야[맺음:야]",
		"학교야":  "학교[열쇠:체언]+(이)[서술격:이(빠짐)]+야[맺음:야]",
		"학교이야": "학교[열쇠:체언]+이[서술격:이]+야[맺음:야]",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	if got := streamText(r, "가야"); strings.Contains(got, "가[열쇠:동사]+야[맺음:야]") {
		t.Fatal("야 after a stem that does not end in 하\n" + got)
	}
	bare, _ := New(Options{})
	if got := streamText(bare, "하야"); strings.Contains(got, "하+야[맺음:야]") {
		t.Fatal("the old spelling 야 without a word table\n" + got)
	}
}
