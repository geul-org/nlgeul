//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNetRunSeedsOf — a root becomes a seed: a state of its own with the root's last word and nobody before it, noted in the registry, in the order the roots end, and a seed stands for each root
package token

import "testing"

func TestNetRunSeedsOf(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	order := r.byEnd()
	roots := endingAt(order, 4)
	seedOf, seeds := r.seedsOf(order, roots)
	if len(seeds) != len(roots) || len(seedOf) != len(roots) || len(r.sc.registry().seeds) != len(roots) {
		t.Fatal(len(seeds), len(roots))
	}
	for root, sd := range seedOf {
		if sd.prev != nil || sd.preds != nil || sd.last.sig != root.last.sig || !r.sc.registry().seeds[sd] {
			t.Fatal("seed", sd.last.sig)
		}
	}
}
