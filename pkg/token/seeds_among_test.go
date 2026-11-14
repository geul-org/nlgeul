//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSeedsAmong — the seeds of the roots that are among a list's states, in the order the roots end; none for an empty list, and a root that is not in the list has no seed there
package token

import "testing"

func TestSeedsAmong(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	order := r.byEnd()
	seedOf, seeds := r.seedsOf(order, endingAt(order, 4))
	if len(seeds) < 2 {
		t.Fatal("roots", len(seeds))
	}
	var roots []*path
	for _, p := range order {
		if _, ok := seedOf[p]; ok {
			roots = append(roots, p)
		}
	}
	got := seedsAmong(order, seedOf, []*path{roots[1], roots[0]})
	if len(got) < 2 || got[0] != seedOf[roots[0]] || got[1] != seedOf[roots[1]] {
		t.Fatal("order of the roots", len(got))
	}
	if got = seedsAmong(order, seedOf, roots[1:2]); len(got) != 1 || got[0] != seedOf[roots[1]] {
		t.Fatal("one", len(got))
	}
	if len(seedsAmong(order, seedOf, nil)) != 0 || len(seedsAmong(order, seedOf, []*path{{}})) != 0 {
		t.Fatal("none")
	}
}
