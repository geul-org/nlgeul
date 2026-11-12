//ff:func feature=token type=reader control=sequence
//ff:what TestBundles — a unit is cut only where all its readings agree: 아버지가방에들어가신다, where some readings put no boundary after 에, is one stretch holding every reading, best first; 고양이가마당에서물을마신다 is cut after 에서, where every reading puts a boundary, and a stretch its readings read differently holds each way once; 다음으로써 is one contest
package token

import "testing"

func TestBundles(t *testing.T) {
	sc := scanOf("아버지가방에들어가신다")
	paths := sc.search()
	got := sc.bundles(paths)
	if len(got) != 1 || got[0].Text != "아버지가방에들어가신다" || len(paths) < 8 || len(got[0].Branches) != len(paths) || got[0].Branches[0].String() != "아버지+가[주격:가|보격:가] 방+에[부사격:에] 들어+가[주격:가|보격:가] 신+다[맺음:다]" || got[0].Branches[len(paths)-1].String() != "아버지+가[주격:가|보격:가] 방에+들[접미:들] 어가신+다[맺음:다]" {
		t.Fatalf("%+v", got)
	}
	cut := scanOf("고양이가마당에서물을마신다")
	parts := cut.bundles(cut.search())
	if len(parts) != 2 || parts[0].Text != "고양이가마당에서" || parts[1].Text != "물을마신다" || len(parts[0].Branches) != 4 || len(parts[1].Branches) != 4 || parts[0].Branches[1].String() != "고양이+가[주격:가|보격:가] 마당+에서[주격:에서|부사격:에서]" || parts[1].Branches[2].String() != "물을마신+다[맺음:다]" {
		t.Fatalf("%+v", parts)
	}
	contest := scanOf("다음으로써")
	if b := contest.bundles(contest.search()); len(b) != 1 || len(b[0].Branches) < 2 {
		t.Fatalf("%+v", b)
	}
}
