//ff:func feature=rule type=helper control=sequence
//ff:what TestMeasure — each side names its length
package rule

import "testing"

func TestMeasure(t *testing.T) {
	ctx := ctxFor(&Word{Chain: Chain{Size: 2}, Longest: 6}, Tail{Size: 4})
	if measure(ctx, SideChain) != 2 || measure(ctx, SideTail) != 4 || measure(ctx, SideLongest) != 6 {
		t.Fatal("measure")
	}
}
