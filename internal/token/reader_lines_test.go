//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLines — with the word table of testdata: 마당에서 and 마신다 put the known-word reading first, the stem looked up as it was before fusing, a reading that leans on a new word stays with its low verdict, a noun is not read before a closing, and 물을 puts the noun with its object particle first, the modifying 을 written onto the verb stem 물 (which ends in ㄹ) after it with a lower verdict
package token

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestReaderLines(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	all := r.Lines("고양이가 마당에서 물을 마신다.", false)
	got := strings.Join(all, "\n")
	for _, want := range []string{
		"  고양이[열쇠:체언]+가[주격:가|보격:가]\t성립 ",
		"  마당[열쇠:체언]+에서[주격:에서|부사격:에서]\t성립 ",
		"  물[열쇠:체언]+을[목적격:을]\t성립 ",
		"  마시[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]\t성립 ",
		"  마신다\t성립 0.33",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in\n%s", want, got)
		}
	}
	noun, verb := "  물[열쇠:체언]+을[목적격:을]\t성립 ", "  물[열쇠:동사]+을[꾸밈:을]\t성립 "
	at := func(prefix string) int {
		return slices.IndexFunc(all, func(l string) bool { return strings.HasPrefix(l, prefix) })
	}
	verdict := func(i int, prefix string) float64 {
		v, _ := strconv.ParseFloat(strings.Fields(strings.TrimPrefix(all[i], prefix))[0], 64)
		return v
	}
	unit, n, v := at("물을\t"), at(noun), at(verb)
	if unit < 0 || n != unit+1 || v <= n || verdict(v, verb) >= verdict(n, noun) {
		t.Errorf("을 after a stem that ends in ㄹ stands after the noun reading of 물을, with a lower verdict, in\n%s", got)
	}
	if strings.Contains(got, "마신[열쇠:체언]+다") || strings.Contains(got, "마[") {
		t.Fatal(got)
	}
	lines := r.Lines("마신다", false)
	if !strings.HasPrefix(lines[1], "  마시[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]\t") {
		t.Fatal(lines)
	}
}
