//ff:func feature=token type=helper control=iteration dimension=1
//ff:what checkRedo — a region a window cut off, when it has what is needed to read it again, gives the same first five ways as the region read again (the same words in the same order with the same lambdas; the numbers of birth are counted again)
package token

import "testing"

func checkRedo(t *testing.T, rg *region) {
	t.Helper()
	if rg.redo == nil {
		return
	}
	again := rg.redo.rebuild()
	for i := 0; i < 5; i++ {
		a, ok := rg.deal()
		b, ok2 := again.deal()
		if ok != ok2 || (ok && (segmentLine(a) != segmentLine(b) || a.Lambda != b.Lambda || a.Ways.Cmp(b.Ways) != 0 || a.Rank != b.Rank)) {
			t.Fatal(rg.span, i, segmentLine(a), segmentLine(b))
		}
	}
	rg.sent = 0
}
