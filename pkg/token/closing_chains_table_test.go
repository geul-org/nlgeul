//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosingChains — every name of closingChains (the ending and the closings it is read after) is the name of a closing of the index; with a word table the ending after a final closing of a quoted sentence is read as one more piece of its tail and comes first (좋+다+거나, 가+(ㄴ)다+라는, 가+리라+는, 학교+(이)+라+기, 가+(ㅂ)시다+고)
package token

import (
	"strings"
	"testing"
)

func TestClosingChains(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["좋"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	names := map[string]bool{}
	for _, es := range r.ix.byKey {
		for _, e := range es {
			names[e.name] = e.class == classClosing
		}
	}
	var listed []string
	for name, befores := range closingChains {
		listed = append(append(listed, name), strings.Split(befores, "|")...)
	}
	for _, n := range listed {
		if !names[n] {
			t.Errorf("%s is no closing of the index", n)
		}
	}
	for in, want := range map[string]string{
		"좋다거나": "좋[열쇠:형용사]+다[맺음:다]+거나[맺음:거나]",
		"간다라는": "가[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]+라는[꾸밈:라는]",
		"가리라는": "가[열쇠:동사]+리라[맺음:리라]+는[꾸밈:는]",
		"학교라기": "학교[열쇠:체언]+(이)[서술격:이(빠짐)]+라[맺음:라]+기[이름:기]",
		"갑시다고": "가[열쇠:동사]+(ㅂ)시다[맺음:ㅂ시다]+고[맺음:고]",
	} {
		if got := flatNodes(r, in, true); len(got) != 1 || !strings.HasSuffix(got[0], "} "+want) {
			t.Errorf("%s: want %s first\n%v", in, want, got)
		}
	}
}
