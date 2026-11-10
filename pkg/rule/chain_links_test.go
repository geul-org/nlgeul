//ff:func feature=rule type=rule control=sequence
//ff:what TestChainLinks — applies when the next particle's kind may not follow this one's; not when it may, not when nothing follows
package rule

import "testing"

func TestChainLinks(t *testing.T) {
	bad, _ := ChainLinks(pieceCtx(Piece{Slot: LinkWith, Next: LinkCase}), nil)
	good, _ := ChainLinks(pieceCtx(Piece{Slot: LinkAdv, Next: LinkTopic}), nil)
	last, _ := ChainLinks(pieceCtx(Piece{Slot: LinkCase, Next: -1}), nil)
	if bad != 1 || good != -1 || last != -1 {
		t.Fatal(bad, good, last)
	}
}
