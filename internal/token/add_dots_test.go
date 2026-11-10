//ff:func feature=token type=rule control=sequence
//ff:what TestAddDots — four joining readings, two closing readings, and an unlisted run's verdict of 0.6
package token

import (
	"math"
	"testing"
)

func TestAddDots(t *testing.T) {
	ix := &index{}
	ix.addDots()
	if len(ix.joins) != 4 || len(ix.stops) != 2 || ix.joins[0].name != "파일명:." || ix.graphs != 7 || ix.edges != 12 || math.Abs(ix.run-(2/(1+forceRun)-1)) > 1e-9 {
		t.Fatal(len(ix.joins), ix.graphs, ix.edges, ix.run)
	}
}
