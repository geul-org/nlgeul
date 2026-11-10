//ff:func feature=token type=reader control=sequence
//ff:what TestCopulaOf — the stem said as a syllable 이 or under a fused piece, the stem dropped after a vowel before a closing that allows it (whole or fused), and no reading for a tail with no closing, a closing that holds the stem, a closing that does not allow dropping, or a consonant before it
package token

import "testing"

func TestCopulaOf(t *testing.T) {
	sc := scanOf("곰이다")
	da, go_, seo := &entry{name: "맺음:다"}, &entry{name: "맺음:고"}, &entry{name: "맺음:서"}
	n := &entry{name: "꾸밈:ㄴ"}
	one := []piece{{text: "다"}}
	if _, _, ok := sc.copulaOf(tail{start: 2, pieces: one}); ok {
		t.Fatal("no closing")
	}
	if _, _, ok := sc.copulaOf(tail{start: 2, pieces: one, close: &entry{holds: true}}); ok {
		t.Fatal("holds")
	}
	if at, e, ok := sc.copulaOf(tail{start: 2, pieces: one, close: da}); !ok || at != 1 || e != sc.ix.copula {
		t.Fatal("whole 이", at)
	}
	if at, e, ok := sc.copulaOf(tail{start: 2, pieces: one, base: '이', close: n}); !ok || at != 1 || e != sc.ix.copula {
		t.Fatal("fused 이", at)
	}
	school := scanOf("학교다")
	if _, _, ok := school.copulaOf(tail{start: 2, pieces: one, close: seo}); ok {
		t.Fatal("not droppable")
	}
	if _, _, ok := school.copulaOf(tail{start: 2, pieces: []piece{{}, {}}, close: da}); ok {
		t.Fatal("slots before the closing")
	}
	if at, e, ok := school.copulaOf(tail{start: 2, pieces: one, close: da}); !ok || at != 2 || e != school.ix.dropped {
		t.Fatal("dropped", at)
	}
	if at, e, ok := scanOf("누군").copulaOf(tail{start: 2, pieces: one, base: '구', close: n}); !ok || at != 2 || e == nil || e.name != droppedName {
		t.Fatal("dropped under a fused piece", at)
	}
	if _, _, ok := scanOf("곰고").copulaOf(tail{start: 1, pieces: one, close: go_}); ok {
		t.Fatal("a consonant before")
	}
	if _, _, ok := scanOf("다").copulaOf(tail{start: 0, pieces: one, close: da}); ok {
		t.Fatal("nothing before")
	}
}
