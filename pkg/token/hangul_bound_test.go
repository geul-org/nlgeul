//ff:func feature=token type=rule control=iteration dimension=2
//ff:what TestHangulBound — the graph of each closing of hangulBound (은, 는, 을 of the modifying forms and 아, 어, 나, 여), with and without a word table, stands after a Hangul syllable and is turned down for whyNoHangul after a Latin letter, a digit, a backtick or a slash; the closings that may follow a run with the copula's stem dropped (다, 고, 면, 지, 네) and the particles 는, 은, 을, 이 stand after those characters as before
package token

import (
	"maps"
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHangulBound(t *testing.T) {
	bare, err := newIndex(Options{})
	if err != nil {
		t.Fatal(err)
	}
	lexed, err := newIndex(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	all := slices.Concat(slices.Concat(slices.Collect(maps.Values(bare.byKey))...), slices.Concat(slices.Collect(maps.Values(lexed.byKey))...))
	bound := slices.DeleteFunc(slices.Clone(all), func(e *entry) bool { return !slices.Contains(hangulBound, e.name) })
	prev := map[string]rune{"꾸밈:은": '먹', "꾸밈:는": '먹', "꾸밈:을": '먹', "맺음:아": '살', "맺음:어": '먹', "맺음:나": '먹', "맺음:여": '하'}
	stands := map[string]bool{}
	for _, e := range bound {
		v, why := verdictOf(e.g, rule.Piece{Form: e.form, HasStem: true, Prev: prev[e.name], Next: -1}, "", true)
		stands[e.name] = stands[e.name] || v > 0
		if slices.Contains(why, whyNoHangul) {
			t.Errorf("%s after a Hangul syllable turned down for %s", e.name, whyNoHangul)
		}
		for _, r := range []rune{'i', 'Z', '7', '`', '/'} {
			if v, why := verdictOf(e.g, rule.Piece{Form: e.form, HasStem: true, Prev: r, Next: -1}, "", true); v > 0 || !slices.Contains(why, whyNoHangul) {
				t.Errorf("%s after %q: verdict %v, reasons %v", e.name, r, v, why)
			}
		}
	}
	for _, name := range hangulBound {
		if !stands[name] {
			t.Errorf("%s stands after no Hangul syllable", name)
		}
	}
	free := []string{"맺음:다", "맺음:고", "맺음:면", "맺음:지", "맺음:네", "보조사:는", "보조사:은", "목적격:을", "주격:이"}
	others := slices.DeleteFunc(slices.Clone(all), func(e *entry) bool { return !slices.Contains(free, e.name) })
	stood := map[string]bool{}
	for _, e := range others {
		for _, r := range []rune{'i', '7', '`'} {
			v, why := verdictOf(e.g, rule.Piece{Form: e.form, HasStem: true, Prev: r, Next: -1}, "", true)
			stood[e.name] = stood[e.name] || v > 0
			if slices.Contains(why, whyNoHangul) {
				t.Errorf("%s after %q turned down for %s (verdict %v)", e.name, r, whyNoHangul, v)
			}
		}
	}
	for _, name := range free {
		if !stood[name] {
			t.Errorf("%s stands after no Latin letter, digit or backtick", name)
		}
	}
}
