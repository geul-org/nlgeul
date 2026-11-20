//ff:func feature=parse type=helper control=sequence
//ff:what TestRegionLift — no lift does nothing (what was sent stays sent); a lift raises the candidate it names in the region's node, the best lambdas are worked out again and the ways begin anew (none sent), so the first way is given again with the raised candidate first and its lambda raised by the lift; the lifts are kept with the region, those of later lifts added after
package parse

import "testing"

func TestRegionLift(t *testing.T) {
	rg := sentenceRegions(t, "학생이되었다")[0]
	first, _ := rg.deal()
	rg.lift(nil)
	if rg.sent != 1 || rg.lifts != nil {
		t.Fatal("no lift changes nothing", rg.sent)
	}
	const comp = "학생[열쇠:체언]+이[보격:이]"
	add := lift{span: Span{0, 3}, word: comp, add: 2}
	rg.lift([]lift{add})
	if rg.sent != 0 || len(rg.lifts) != 1 || rg.lifts[0] != add {
		t.Fatal(rg.sent, rg.lifts)
	}
	seg, ok := rg.deal()
	if !ok || seg.Rank != 1 || segmentLine(seg) != segmentLine(first) || seg.Nodes[0].Contest.Branches[0].String() != comp || !almost(seg.Lambda, first.Lambda+2) {
		t.Fatal(segmentLine(seg), seg.Lambda, first.Lambda)
	}
	if len(seg.Nodes[0].Contest.Branches) != len(first.Nodes[0].Contest.Branches) {
		t.Fatal("no candidate is taken away")
	}
	more := lift{span: Span{3, 6}, word: "되[열쇠:동사]+었[선어말:었]+다[맺음:다]", add: 1}
	rg.lift([]lift{more})
	if len(rg.lifts) != 2 || rg.lifts[1] != more {
		t.Fatal(rg.lifts)
	}
	if again, _ := rg.deal(); !almost(again.Lambda, first.Lambda+3) {
		t.Fatal(again.Lambda, first.Lambda)
	}
}
