//ff:func feature=parse type=reader control=sequence
//ff:what TestAfterClosingsNamed — a particle of closingParticles right after a closing of its row is one more piece of the predicate's tail under its own name (먹+었+다+야), not after a closing the row does not list (먹+자); a particle read only right after a closing (ended) is read so by its name alone: the 야 after a closing that ends in a consonant stands after a linking closing of the few (가+(ㄴ)다면+야), and such a particle is not read as any particle after a naming closing is, nor as the particles after a closing that asks or after one of the few closings a form is written onto are
package parse

import "testing"

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
}
