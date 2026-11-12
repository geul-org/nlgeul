//ff:func feature=token type=reader control=sequence
//ff:what TestBundles — a unit is cut only where all its readings agree: 아버지가방에들어가신다, where some readings put no boundary after 에, is one stretch holding every reading (none dropped for a limit), best first (the best one named, the whole-word new reading among them); 고양이가마당에서물을마신다, where some reading puts no boundary after 에서, is one stretch too, and holds the right reading of it; 다음으로써 is one contest
package token

import "testing"

func TestBundles(t *testing.T) {
	sc := scanOf("아버지가방에들어가신다")
	paths := sc.search()
	got := sc.bundles(paths)
	best := "아버지+가[주격:가|보격:가] 방+에[부사격:에] 들어+가[주격:가|보격:가] 신+다[맺음:다]"
	whole := "아버지+가[주격:가|보격:가] 방에+들[접미:들] 어가신+다[맺음:다]"
	if len(got) != 1 || got[0].Text != "아버지가방에들어가신다" || len(paths) < 8 || len(got[0].Branches) != len(paths) || got[0].Branches[0].String() != best || !hasBranch(got[0].Branches, whole) || !branchesInOrder(got[0].Branches) {
		t.Fatalf("%+v", got)
	}
	cut := scanOf("고양이가마당에서물을마신다")
	parts := cut.bundles(cut.search())
	right := "고양이+가[주격:가|보격:가] 마당+에서[주격:에서|부사격:에서] 물+을[목적격:을] 마신+다[맺음:다]"
	if len(parts) != 1 || parts[0].Text != "고양이가마당에서물을마신다" || len(parts[0].Branches) < 4 || !hasBranch(parts[0].Branches, right) || !branchesInOrder(parts[0].Branches) {
		t.Fatalf("%+v", parts)
	}
	contest := scanOf("다음으로써")
	if b := contest.bundles(contest.search()); len(b) != 1 || len(b[0].Branches) < 2 {
		t.Fatalf("%+v", b)
	}
}
