//ff:func feature=token type=reader control=sequence
//ff:what TestHelped — a tail that attaches to 지 is read once more after the 어/아 form before 지, with 지 as a piece between and the main verb asked to be a predicate of the table; so too when the piece is fused into 지 and when the 어/아 leaves no trace; not without a word table, after another syllable, where no 어/아 form ends, or for a tail that names its stem
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHelped(t *testing.T) {
	ws := wordsScan("먹어지는", layerCost)
	neun := []piece{{text: "는", names: "꾸밈:는", verdict: 1, tail: rule.TailVerb}}
	got := ws.helped([]tail{{start: 3, pieces: neun, close: &entry{name: "꾸밈:는"}}})
	if len(got) < 1 || got[0].start != 1 || len(got[0].pieces) != 3 || got[0].pieces[1].names != "보조:지" || got[0].pieces[2].tail != rule.TailPred || got[0].gate != ws.ix.pred || got[0].close == nil {
		t.Fatal(got)
	}
	fused := wordsScan("지워진", layerCost)
	got = fused.helped([]tail{{start: 3, base: '지', pieces: []piece{{text: "(ㄴ)", names: "꾸밈:ㄴ", verdict: 1, tail: rule.TailPred}}, gate: fused.ix.pred}})
	if len(got) < 1 || got[0].base != '우' || got[0].pieces[0].text != "(어)" || got[0].pieces[1].text != "지" {
		t.Fatal("a piece fused into 지 after a fused 어", got)
	}
	if got = scanOf("먹어지는").helped([]tail{{start: 3, pieces: neun}}); len(got) != 0 {
		t.Fatal("no word table", got)
	}
	if got = ws.helped([]tail{{start: 2, pieces: neun}, {start: 4, pieces: neun}, {start: 3, pieces: neun, only: "아니"}, {start: 3, pieces: neun, gate: ws.ix.nominal}}); len(got) != 0 {
		t.Fatal("after another syllable, or a tail with its own stem", got)
	}
	if got = wordsScan("먹고지는", layerCost).helped([]tail{{start: 3, pieces: neun}}); len(got) != 0 {
		t.Fatal("no 어/아 form before 지", got)
	}
}
