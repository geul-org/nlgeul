//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSearch — 다음으로써 comes back with 다음+으로써 first and 다음+으로 써 second; no reading takes a marker for a new word's stem after a boundary; an empty unit has no readings
package token

import "testing"

func TestSearch(t *testing.T) {
	got := scanOf("다음으로써").search()
	if len(got) < 2 || got[0].sig != "다음+으로써[부사격:으로써]" || got[1].sig != "다음+으로[부사격:으로] 써" {
		t.Fatal(got[0].sig, got[1].sig)
	}
	for _, p := range wordsScan("곰이다", layerCost).search() {
		if p.sig == "곰 이+다[맺음:다]" {
			t.Fatal("a marker read as a new word's stem after a boundary", p.sig)
		}
	}
	if scanOf("").search() != nil {
		t.Fatal("empty")
	}
}
