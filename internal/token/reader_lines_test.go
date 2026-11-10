//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLines — with the word table of testdata: 마당에서 and 마신다 put the known-word reading first, the stem looked up as it was before fusing, and readings that lean on a new word are left out; 고양이가 keeps 고양이+가 first
package token

import (
	"strings"
	"testing"
)

func TestReaderLines(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(r.Lines("고양이가 마당에서 물을 마신다.", false), "\n")
	for _, want := range []string{
		"  고양이[열쇠:체언]+가[주격:가|보격:가]\t성립 1.00",
		"마당[열쇠:체언]+에서[주격:에서|부사격:에서]",
		"  물[열쇠:체언]+을[목적격:을]\t성립 1.00",
		"  물[열쇠:용언]+을[꾸밈:을]\t성립 1.00",
		"마시[열쇠:용언]+(ㄴ)다[맺음:ㄴ다]",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in\n%s", want, got)
		}
	}
	if strings.Contains(got, "성립 0.33") || strings.Contains(got, "마신[열쇠:체언]+다") {
		t.Fatal(got)
	}
}
