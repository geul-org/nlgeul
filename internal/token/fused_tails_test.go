//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestFusedTails — a tail for each stem the fused syllable may come from, the piece in brackets before the rest of the form; a fused 어/아 that changes no letter is marked as leaving no trace and, with a word table, asks for a predicate's stem before it; so does a fused 어/아 or ㅆ that shows (봐, 줬), but 해 and 했 ask for nothing; 셨 is read as the honorific too; end to end 쏴 is no 쏘+아 until the table has 쏘, and 곰해 keeps the new word 곰하+여
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestFusedTails(t *testing.T) {
	sc := scanOf("갔다")
	got := sc.fusedTails(piece{names: "선어말:ㅆ"}, &entry{lead: rule.JongSS}, 1)
	if len(got) != 2 || got[0].base != '가' || got[0].pieces[0].text != "(았)" || got[1].base != '그' || got[0].start != 1 {
		t.Fatal(got)
	}
	if unseen := scanOf("가").fusedTails(piece{}, &entry{lead: rule.LeadVowel}, 1); len(unseen) != 2 || !unseen[0].pieces[0].unseen {
		t.Fatal(unseen)
	}
	if seen := scanOf("봐").fusedTails(piece{}, &entry{lead: rule.LeadVowel}, 1); len(seen) != 1 || seen[0].pieces[0].unseen {
		t.Fatal(seen)
	}
	ws := wordsScan("가서", layerCost)
	if gate := ws.fusedTails(piece{}, &entry{lead: rule.LeadVowel, form: "서"}, 1); len(gate) != 2 || gate[0].gate != ws.ix.pred || gate[0].gate == nil {
		t.Fatal("an 어/아 that leaves no trace asks for a predicate's stem of the table", gate)
	}
	if gate := wordsScan("봐", layerCost).fusedTails(piece{}, &entry{lead: rule.LeadVowel}, 1); gate[0].gate == nil {
		t.Fatal("a fused 어/아 that shows asks for a predicate's stem too", gate)
	}
	if gate := wordsScan("해", layerCost).fusedTails(piece{}, &entry{lead: rule.LeadVowel}, 1); gate[0].gate != nil {
		t.Fatal("해 asks for nothing", gate)
	}
	if gate := wordsScan("줬다", layerCost).fusedTails(piece{}, &entry{lead: rule.JongSS}, 1); len(gate) != 1 || gate[0].gate == nil || gate[0].base != '주' {
		t.Fatal("a fused ㅆ that shows asks for a predicate's stem", gate)
	}
	if gate := wordsScan("했다", layerCost).fusedTails(piece{}, &entry{lead: rule.JongSS}, 1); len(gate) != 1 || gate[0].gate != nil {
		t.Fatal("했 asks for nothing", gate)
	}
	sc = scanOf("마신다")
	got = sc.fusedTails(piece{names: "맺음:ㄴ다"}, &entry{form: "다", lead: rule.JongN, noYo: true}, 2)
	if len(got) != 1 || got[0].base != '시' || got[0].pieces[0].text != "(ㄴ)다" || !got[0].noYo {
		t.Fatal(got)
	}
	var honor bool
	for _, c := range scanOf("가셨다").fusedTails(piece{}, &entry{lead: rule.JongSS}, 2) {
		honor = honor || (c.start == 1 && c.pieces[0].names == honorName)
	}
	if !honor {
		t.Fatal("no honorific under 셨")
	}
	r, _ := New(Options{Words: "testdata/words.tsv"})
	if got := strings.Join(r.Lines("쏴", false), "\n"); strings.Contains(got, "쏘+(아)") || strings.Contains(got, "쏘[") {
		t.Fatal("쏘 is no predicate's stem of the table\n" + got)
	}
	if got := strings.Join(r.Lines("곰해", false), "\n"); !strings.Contains(got, "곰하+(여)[맺음:어(합침)]") {
		t.Fatal("해 asks for nothing\n" + got)
	}
	r.ix.lex["쏘"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	if got := strings.Join(r.Lines("쏴", false), "\n"); !strings.Contains(got, "쏘[열쇠:동사]+(아)[맺음:어(합침)]") {
		t.Fatal("쏘 of the table\n" + got)
	}
}
