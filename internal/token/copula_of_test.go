//ff:func feature=token type=reader control=sequence
//ff:what TestCopulaOf — the stem said as a syllable 이 or under a fused piece, the stem dropped after a vowel before a closing that allows it (whole or fused), and no reading for a tail with no closing, a closing that holds the stem, a closing that does not allow dropping, or a consonant before it; the closings of droppableKnown (네, 냐, 야, the naming 기) allow it only with a word table; a letter that is no Hangul syllable before it (abc, 3) allows it like a vowel, before 야 and before 다 alike, leaving to the graph which closing such a letter takes
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
	ji := []piece{{text: "지", names: jiName}, {text: "(않)"}, {text: "아"}}
	if at, e, ok := school.copulaOf(tail{start: 2, pieces: ji, close: &entry{name: jiName}}); !ok || at != 2 || e != school.ix.dropped {
		t.Fatal("dropped before 잖", at)
	}
	if _, _, ok := wordsScan("일본", layerCost).copulaOf(tail{start: 2, pieces: one, base: '보', close: n, gate: &entry{}}); ok {
		t.Fatal("a lone consonant that asks for a predicate's stem has no dropped copula")
	}
	if _, _, ok := school.copulaOf(tail{start: 2, pieces: one, close: da, irr: irrS}); ok {
		t.Fatal("a tail that asks for a changed stem")
	}
	if _, _, ok := scanOf("곰고").copulaOf(tail{start: 1, pieces: one, close: go_}); ok {
		t.Fatal("a consonant before")
	}
	if _, _, ok := scanOf("다").copulaOf(tail{start: 0, pieces: one, close: da}); ok {
		t.Fatal("nothing before")
	}
	ne := tail{start: 1, pieces: one, close: &entry{name: "맺음:네"}}
	if at, e, ok := wordsScan("거네", layerCost).copulaOf(ne); !ok || at != 1 || e == nil || e.name != droppedName {
		t.Fatal("dropped before 네 with a word table", at)
	}
	if _, _, ok := scanOf("거네").copulaOf(ne); ok {
		t.Fatal("not before 네 without a word table")
	}
	if at, _, ok := wordsScan("뭐냐", layerCost).copulaOf(tail{start: 1, pieces: one, close: &entry{name: "맺음:냐"}}); !ok || at != 1 {
		t.Fatal("dropped before 냐 with a word table")
	}
	if _, _, ok := wordsScan("학교서", layerCost).copulaOf(tail{start: 2, pieces: one, close: seo}); ok {
		t.Fatal("a closing of neither list, with a word table")
	}
	ya := tail{start: 3, pieces: one, close: &entry{name: "맺음:야"}}
	if at, e, ok := wordsScan("abc야", layerCost).copulaOf(ya); !ok || at != 3 || e == nil || e.name != droppedName {
		t.Fatal("dropped before 야 after a letter that is no Hangul syllable, with a word table", at)
	}
	if at, e, ok := wordsScan("3야", layerCost).copulaOf(tail{start: 1, pieces: one, close: &entry{name: "맺음:야"}}); !ok || at != 1 || e == nil || e.name != droppedName {
		t.Fatal("dropped before 야 after a digit, with a word table", at)
	}
	if _, _, ok := scanOf("abc야").copulaOf(ya); ok {
		t.Fatal("not before 야 without a word table")
	}
	if at, e, ok := scanOf("abc다").copulaOf(tail{start: 3, pieces: one, close: da}); !ok || at != 3 || e == nil || e.name != droppedName {
		t.Fatal("before 다 after a letter that is no Hangul syllable the reading is offered, and its graph decides", at)
	}
	gi := tail{start: 2, pieces: one, close: &entry{name: "이름:기"}}
	if at, e, ok := wordsScan("문제기", layerCost).copulaOf(gi); !ok || at != 2 || e == nil || e.name != droppedName {
		t.Fatal("dropped before the naming 기 with a word table (the graph asks for a particle after it)", at)
	}
	if _, _, ok := scanOf("문제기").copulaOf(gi); ok {
		t.Fatal("not before the naming 기 without a word table")
	}
}
