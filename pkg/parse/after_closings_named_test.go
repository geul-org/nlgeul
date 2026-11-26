//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAfterClosingsNamed — a particle of closingParticles right after a closing of its row is one more piece of the predicate's tail under its own name (먹+었+다+야), not after a closing the row does not list (먹+자); a particle read only right after a closing (ended) is read so by its name alone: the 야 after a closing that ends in a consonant stands after a linking closing of the few (가+(ㄴ)다면+야), and such a particle is not read as any particle after a naming closing is, nor as the particles after a closing that asks or after one of the few closings a form is written onto are; a particle of the table is not read where the letters of the stem, the closing and the particle are a word of the word table; a case particle of the table is read after the copula's stem said before the closing (학교+이+다+를), not after it dropped (학교다를), which a particle of no case is
package parse

import (
	"slices"
	"testing"
)

func TestAfterClosingsNamed(t *testing.T) {
	ya := piece{text: "야", names: "보조사:야", verdict: 1}
	got := scanOf("먹었다야").afterClosings(3, 0, ya, &entry{name: "보조사:야"})
	if len(got) == 0 || len(got[0].pieces) < 2 || got[0].pieces[len(got[0].pieces)-2].names != "맺음:다" || got[0].pieces[len(got[0].pieces)-1] != ya {
		t.Fatal("다 and 야 after it", got)
	}
	if got = scanOf("먹자야").afterClosings(2, 0, ya, &entry{name: "보조사:야"}); len(got) != 0 {
		t.Fatal("after a closing the row does not list", got)
	}
	ended := &entry{name: "보조사:야", ended: true}
	if got = scanOf("간다면야").afterClosings(3, 0, ya, ended); len(got) == 0 || got[0].pieces[len(got[0].pieces)-1].names != "보조사:야" {
		t.Fatal("야 after a linking closing that ends in a consonant", got)
	}
	if got = wordsScan("먹음야", layerCost).afterClosings(2, 0, ya, ended); len(got) != 0 {
		t.Fatal("an ended particle after a naming closing", got)
	}
	if got = wordsScan("먹음야", layerCost).afterClosings(2, 0, ya, &entry{name: "보조사:야"}); len(got) == 0 {
		t.Fatal("any particle after a naming closing", got)
	}
	if got = scanOf("먹느냐야").afterClosings(3, 0, ya, ended); len(got) != 0 {
		t.Fatal("an ended particle after a closing that asks", got)
	}
	if got = scanOf("먹느냐야").afterClosings(3, 0, ya, &entry{name: "보조사:야"}); len(got) == 0 {
		t.Fatal("야 after a closing that asks", got)
	}
	reul := piece{text: "를", names: "목적격:를", verdict: 1}
	if got = scanOf("먹지를").afterClosings(2, 0, reul, &entry{name: "목적격:를", ended: true}); len(got) != 0 {
		t.Fatal("an ended particle after a closing its form is written onto", got)
	}
	if got = scanOf("먹지를").afterClosings(2, 0, reul, &entry{name: "목적격:를"}); len(got) == 0 {
		t.Fatal("를 after 지", got)
	}
	na := piece{text: "나", names: "보조사:나", verdict: 1}
	over := wordsScan("마셔나", layerCost)
	if got = over.afterClosings(2, 0, na, &entry{name: "보조사:나"}); len(got) == 0 || got[0].pieces[len(got[0].pieces)-1] != na {
		t.Fatal("나 after a linking closing fused into the stem", got)
	}
	over.ix.lex["마셔나"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	if got = over.afterClosings(2, 0, na, &entry{name: "보조사:나"}); len(got) != 0 {
		t.Fatal("a particle where the letters are a word of the word table", got)
	}
	obj := &entry{name: "목적격:를", kase: 1}
	said, dropped := 0, 0
	for _, c := range scanOf("학교이다를").afterClosings(4, 0, reul, obj) {
		if len(c.pieces) == 3 && c.pieces[0].names == copulaName {
			said++
		}
	}
	for _, c := range scanOf("학교다를").afterClosings(3, 0, reul, obj) {
		if droppedIn(c) {
			dropped++
		}
	}
	if said != 1 || dropped != 0 {
		t.Fatal("a case particle after the copula's stem said, not after it dropped", said, dropped)
	}
	if got = scanOf("학교다를").afterClosings(3, 0, reul, &entry{name: "목적격:를"}); !slices.ContainsFunc(got, droppedIn) {
		t.Fatal("a particle of no case after a dropped stem", got)
	}
}
