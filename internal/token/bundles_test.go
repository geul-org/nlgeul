//ff:func feature=token type=reader control=sequence
//ff:what TestBundles — 아버지가방에들어가신다 is cut where its readings agree, and a stretch they read differently holds each way once, the reading with a boundary after 가 first
package token

import "testing"

func TestBundles(t *testing.T) {
	sc := scanOf("아버지가방에들어가신다")
	got := sc.bundles(sc.search())
	contest := scanOf("다음으로써")
	if b := contest.bundles(contest.search()); len(b) != 1 || len(b[0].branches) < 2 {
		t.Fatalf("%+v", b)
	}
	if len(got) != 2 || got[0].text != "아버지가방에" || len(got[0].branches) < 2 || got[0].branches[0].marked != "아버지+가[주격:가|보격:가] 방+에[부사격:에]" || got[1].text != "들어가신다" {
		t.Fatalf("%+v", got)
	}
}
