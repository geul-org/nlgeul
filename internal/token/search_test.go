//ff:func feature=token type=reader control=sequence
//ff:what TestSearch — 다음으로써 comes back with 다음+으로써 first and 다음+으로 써+(어) second; an empty unit has no readings
package token

import "testing"

func TestSearch(t *testing.T) {
	got := scanOf("다음으로써").search()
	if len(got) < 2 || got[0].sig != "다음+으로써[부사격:으로써]" || got[1].sig != "다음+으로[부사격:으로] 써+(어)[맺음:어(합침)]" {
		t.Fatal(got[0].sig, got[1].sig)
	}
	if scanOf("").search() != nil {
		t.Fatal("empty")
	}
}
