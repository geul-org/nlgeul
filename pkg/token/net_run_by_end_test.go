//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNetRunByEnd — the states kept in the order of where they end, a copy of the list, those that end at the same place in the order they were made
package token

import "testing"

func TestNetRunByEnd(t *testing.T) {
	r, _ := windowSteps(t, 2, 13)
	order := r.byEnd()
	all := r.sc.registry().all
	if len(order) != len(all) {
		t.Fatal(len(order), len(all))
	}
	for i := 1; i < len(order); i++ {
		if order[i-1].last.end > order[i].last.end {
			t.Fatal("order", i)
		}
	}
	order[0] = nil
	if all[0] == nil {
		t.Fatal("a copy")
	}
}
