//ff:func feature=token type=rule control=sequence
//ff:what TestAddWide — 신 stands as the honorific at a closing's end, not before a particle or something else; 심 stands before a particle; the shortened closings and particles are in the index with their gates; particles are found by their first letter
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddWide(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	graphs := ix.graphs
	ix.addWide()
	ok := func(e *entry, right int) bool {
		v, _ := verdictOf(e.g, rule.Piece{HasStem: true, Prev: '가', Next: -1, Right: right}, "x")
		return v > 0
	}
	if !ok(ix.wide['신'], rule.RightDone) || ok(ix.wide['신'], rule.RightParticle) || ok(ix.wide['신'], rule.RightOther) {
		t.Fatal("신")
	}
	if !ok(ix.wide['심'], rule.RightParticle) || ok(ix.wide['심'], rule.RightOther) || !ok(ix.wide['세'], rule.RightDone) {
		t.Fatal("심, 세")
	}
	if ix.graphs != graphs+2+4+16+len(shortSlots)+2 || !slices.Contains(ix.heads['에'], "에서") || ix.pred == nil || ix.nominal == nil {
		t.Fatal(ix.graphs - graphs)
	}
	if e := ix.named('댔', "선어말:ㄴ댔(ㄴ다고 했)"); e.gate != ix.pred || !e.wide || e.class != classSlot || e.head != headN {
		t.Fatal("a shortened quotation with the past fused in", e)
	}
	if e := ix.named('선', "부사격:선(에서는)"); e.gate != ix.nominal || !e.wide {
		t.Fatal(e)
	}
	if e := ix.named('대', "맺음:ㄴ대(ㄴ다고 해)"); e.gate != ix.pred || e.tail != rule.TailVerb || !slices.Contains(ix.barred, e.name) {
		t.Fatal(e)
	}
}
