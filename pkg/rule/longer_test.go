//ff:func feature=rule type=rule control=sequence
//ff:what TestLonger — applies only when the rival side is strictly longer: chain over tail, longest tail over chain, longest tail over this tail
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestLonger(t *testing.T) {
	w := &Word{Chain: Chain{Size: 4}, Longest: 6}
	chain, _ := Longer(ctxFor(w, Tail{Size: 2}), toulmin.Specs{&Versus{Mine: SideTail, Rival: SideChain}})
	equal, _ := Longer(ctxFor(w, Tail{Size: 4}), toulmin.Specs{&Versus{Mine: SideTail, Rival: SideChain}})
	tail, _ := Longer(ctxFor(w, Tail{}), toulmin.Specs{&Versus{Mine: SideChain, Rival: SideLongest}})
	cut, _ := Longer(ctxFor(w, Tail{Size: 6}), toulmin.Specs{&Versus{Mine: SideTail, Rival: SideLongest}})
	if chain != 1 || equal != -1 || tail != 1 || cut != -1 {
		t.Fatal(chain, equal, tail, cut)
	}
}
