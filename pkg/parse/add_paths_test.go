//ff:func feature=parse type=rule control=sequence
//ff:what TestAddPaths — the graphs of the two readings of a path that the letters alone do not settle are made (two graphs, two edges) and their verdicts worked out: a path's Hangul name stands fully, lower when the Hangul reads as markers on the path (forcePathMarked), and a path with no head and one slash at the verdict its rebuttal leaves (forcePathWeak)
package parse

import (
	"math"
	"testing"
)

func TestAddPaths(t *testing.T) {
	ix := &index{}
	ix.addPaths()
	if ix.graphs != 2 || ix.edges != 2 {
		t.Fatal(ix.graphs, ix.edges)
	}
	if ix.pathNamed != 1 || math.Abs(ix.pathMarked-(2/(1+forcePathMarked)-1)) > 1e-9 || math.Abs(ix.pathWeak-(2/(1+forcePathWeak)-1)) > 1e-9 {
		t.Fatal(ix.pathNamed, ix.pathMarked, ix.pathWeak)
	}
	if ix.pathMarked >= ix.pathWeak || ix.pathWeak >= ix.pathNamed {
		t.Fatal("the order of the verdicts", ix.pathMarked, ix.pathWeak)
	}
}
