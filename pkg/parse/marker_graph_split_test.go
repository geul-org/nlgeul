//ff:func feature=parse type=rule control=sequence
//ff:what TestMarkerGraphSplit — with a word table a closing of whole syllables that ends a closed-set word has one exception more, which lifts the word's exception where a predicate's stem of the table stands before the closing (하+지만) and not where none does; a particle, a fused closing and a graph built with no word table get none
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraphSplit(t *testing.T) {
	whole := []string{"하지만"}
	plain := entry{name: "맺음:지만", form: "지만", class: classClosing}
	lexed := plain
	if a, b := markerGraph(&plain, whole, nil, false), markerGraph(&lexed, whole, nil, true); b != a+1 || plain.split || !lexed.split {
		t.Fatal("one exception more with a word table", a, b, plain.split, lexed.split)
	}
	at := func(e *entry, kinds int) (float64, []string) {
		return verdictOf(e.g, rule.Piece{Form: "지만", HasStem: true, Prev: '하', Next: -1, Stem: kinds}, "하지만", true)
	}
	if v, _ := at(&lexed, kindVerb); v <= 0 {
		t.Fatal("after a verb's stem of the table", v)
	}
	if v, _ := at(&lexed, kindAdj); v <= 0 {
		t.Fatal("after an adjective's stem of the table", v)
	}
	if v, why := at(&lexed, kindNominal); v > 0 || !slices.Contains(why, whyWhole) {
		t.Fatal("after a nominal the closed-set word holds", v, why)
	}
	if v, why := at(&plain, kindVerb); v > 0 || !slices.Contains(why, whyWhole) {
		t.Fatal("with no word table the closed-set word holds", v, why)
	}
	particle := entry{name: "보조사:만", form: "만", class: classParticle}
	fused := entry{name: "맺음:ㄴ데", form: "데", class: classClosing, lead: rule.JongN}
	bare := plain
	markerGraph(&particle, whole, nil, true)
	markerGraph(&fused, []string{"그런데"}, nil, true)
	markerGraph(&bare, nil, nil, true)
	if particle.split || fused.split || bare.split {
		t.Fatal("a particle, a fused closing and a closing that ends no closed-set word are not split", particle.split, fused.split, bare.split)
	}
}
