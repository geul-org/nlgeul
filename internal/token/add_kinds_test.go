//ff:func feature=token type=rule control=sequence
//ff:what TestAddKinds — a graph for each kind; a nominal takes particles and stands alone but takes no predicate tail; a predicate stem takes only a predicate tail
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddKinds(t *testing.T) {
	ix := &index{}
	ix.addKinds()
	if ix.graphs != 5 || ix.edges != 5 {
		t.Fatal(ix.graphs, ix.edges)
	}
	if ix.fits[kindNominal][rule.TailCase] != 1 || ix.fits[kindNominal][rule.TailBare] != 1 || ix.fits[kindNominal][rule.TailPred] != -1 || ix.fits[kindVerbal][rule.TailPred] != 1 || ix.fits[kindVerbal][rule.TailBare] != -1 || ix.fits[kindAdverb][rule.TailCase] != -1 {
		t.Fatal(ix.fits)
	}
}
