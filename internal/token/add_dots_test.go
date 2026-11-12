//ff:func feature=token type=rule control=sequence
//ff:what TestAddDots — four joining candidates, two closing candidates named with no sign after them, and an unlisted run's verdict of 0.6
package token

import (
	"math"
	"testing"
)

func TestAddDots(t *testing.T) {
	ix := &index{}
	ix.addDots()
	if ix.stops[0].name != "마침표" || ix.stops[1].name != "줄임표" {
		t.Fatal("the closing candidates are named with no sign", ix.stops[0].name, ix.stops[1].name)
	}
	if len(ix.joins) != 4 || len(ix.stops) != 2 || ix.joins[0].name != "파일명:." || ix.graphs != 7+len(linkSigns) || ix.edges != 12+len(linkSigns) || ix.links['/'][0].name != "대비:/" || math.Abs(ix.run-(2/(1+forceRun)-1)) > 1e-9 {
		t.Fatal(len(ix.joins), ix.graphs, ix.edges, ix.run)
	}
}
