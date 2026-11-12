//ff:func feature=token type=reader control=sequence
//ff:what TestStepOf — the word from b to end that tail t makes as the stem is read: the stem and the pieces as parts with their names, hidden letters and written forms; the least verdict (a stem that does not stand has the new word's); the preference of the stem and the pieces, or the tail's weight; whether it ends in a modifying closing, is a determiner, or ends in the linking closing; the boundary before an auxiliary verb's stem given back only for a predicate tail after a boundary
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStepOf(t *testing.T) {
	sc := wordsScan("열어두는어떤", layerCost)
	verb := stemCandidate{kindCandidate{kindVerb, "열쇠:동사", 1}, "열", ""}
	to := tail{start: 3, pieces: []piece{{"는", "꾸밈:는", 0.8, -1, rule.TailVerb, false, ""}}}
	e := sc.stepOf(2, 4, to, verb)
	if e.sig != "열[열쇠:동사]+는[꾸밈:는]" || e.end != 4 || e.kind != kindVerb || e.verdict != 0.8 || !e.adn || e.det || e.conn || e.word.End != 4 || len(e.word.Parts) != 2 || !e.word.Parts[0].Stem {
		t.Fatal(e)
	}
	if e.back != 0 {
		t.Fatal("a stem that is no auxiliary makes no boundary before a stem that is no auxiliary", e.back)
	}
	aux := stemCandidate{kindCandidate{kindVerb, "보조 용언:두", 1}, "두", ""}
	if e = sc.stepOf(2, 4, to, aux); e.back != costBoundary || e.sig != "두[보조 용언:두]+는[꾸밈:는]" {
		t.Fatal("an auxiliary verb's stem after a boundary, before a predicate tail", e)
	}
	if e = sc.stepOf(0, 2, to, aux); e.back != 0 {
		t.Fatal("as the unit's first word", e.back)
	}
	case_ := tail{start: 3, pieces: []piece{{"를", "목적격:를", 1, -1, rule.TailCase, false, ""}}}
	if e = sc.stepOf(2, 4, case_, aux); e.back != 0 {
		t.Fatal("before a particle", e.back)
	}
	other := stemCandidate{kindCandidate{kindVerb, "열쇠:동사", 1}, "열", ""}
	if e = sc.stepOf(2, 4, to, other); e.back != 0 {
		t.Fatal("a stem that is no auxiliary", e.back)
	}
	unseen := stemCandidate{kindCandidate{0, "", 0}, "두", ""}
	if e = sc.stepOf(2, 4, to, unseen); e.back != 0 {
		t.Fatal("an auxiliary verb's letters that stand as no word", e.back)
	}
	free := stemCandidate{kindCandidate{0, "", 0}, "곰", ""}
	if e = sc.stepOf(2, 4, to, free); e.back != 0 || e.verdict != sc.unlisted(2, 4, "곰", to) || e.verdict == 0.8 || e.word.Parts[0].Name != (Name{}) || e.word.Parts[0].Text != "곰" {
		t.Fatal("a stem that does not stand has no name and the new word's verdict", e)
	}
	wrote := stemCandidate{kindCandidate{kindNominal, "열쇠:체언", 1}, "역할", "역활"}
	if e = sc.stepOf(0, 2, tail{start: 2}, wrote); e.word.Parts[0].Written != "역활" || e.word.Parts[0].Name != (Name{"열쇠", "체언"}) || e.adn || e.det {
		t.Fatal(e)
	}
	hidden := tail{start: 2, pieces: []piece{{"(아)", "맺음:어(합침)", 1, 0, rule.TailPred, true, "어"}}}
	if e = sc.stepOf(0, 2, hidden, other); e.word.Parts[1].Hidden != "아" || e.word.Parts[1].Text != "" || !e.word.Parts[1].Unseen || e.word.Parts[1].Written != "어" || !e.conn {
		t.Fatal("a piece fused into the stem, a closing that links", e.word.Parts)
	}
	det := stemCandidate{kindCandidate{kindDeterminer, "관형사:어떤", 1}, "어떤", ""}
	if e = sc.stepOf(4, 6, tail{start: 6}, det); !e.det || e.adn || e.conn {
		t.Fatal("a determiner", e)
	}
	if e = sc.stepOf(4, 6, case_, det); e.det {
		t.Fatal("a determiner with a piece after it is no determiner", e)
	}
	as := tail{start: 3, as: numName, weight: -1.5, pieces: []piece{{"명", "단위:명", 1, -1, 0, false, ""}}}
	if e = sc.stepOf(0, 4, as, other); e.pref != -1.5-costPiece+(-1) {
		t.Fatal("a tail with its own weight", e.pref)
	}
}
