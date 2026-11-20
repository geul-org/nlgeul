//ff:func feature=parse type=reader control=sequence
//ff:what TestMainBeforeAux — a stem that begins its unit, whose word ends in the linking 어/아 (whole or fused) right before the letters of an auxiliary verb's stem (봐 = 보+아, 둬, 줄), is a main verb's with the auxiliary written onto it; not away from the unit's head, not at the unit's end, not with no tail, not with the copula's tail (a nominal before 이+어), not with another closing (고), and not before letters that begin no auxiliary's stem; 가 begins one only with a word table
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMainBeforeAux(t *testing.T) {
	fused := tail{start: 1, base: '가', pieces: []piece{{text: "(아)", names: "맺음:어(합침)", verdict: 1, tail: rule.TailPred}}}
	whole := tail{start: 1, pieces: []piece{{text: "어", names: "맺음:어", verdict: 1, tail: rule.TailPred}}}
	other := tail{start: 1, pieces: []piece{{text: "고", names: "맺음:고", verdict: 1, tail: rule.TailPred}}}
	if !scanOf("가봐라").mainBeforeAux(0, 1, fused) {
		t.Fatal("가 before 봐 (보+아)")
	}
	if !scanOf("먹어둬").mainBeforeAux(0, 2, whole) || !scanOf("써줄").mainBeforeAux(0, 1, fused) {
		t.Fatal("before 둬 (두+어) and 줄 (주+ㄹ)")
	}
	if scanOf("또가봐라").mainBeforeAux(1, 2, fused) {
		t.Fatal("away from the unit's head")
	}
	if scanOf("가").mainBeforeAux(0, 1, fused) {
		t.Fatal("at the unit's end")
	}
	if scanOf("가봐라").mainBeforeAux(0, 1, tail{start: 1}) {
		t.Fatal("with no tail")
	}
	copula := tail{start: 1, pieces: []piece{{text: "이", names: "서술격:이", verdict: 1, tail: rule.TailCopula}, {text: "(어)", names: "맺음:어(합침)", verdict: 1, tail: rule.TailPred}}}
	if scanOf("내여보다").mainBeforeAux(0, 2, copula) {
		t.Fatal("with the copula's tail")
	}
	if scanOf("먹고봐").mainBeforeAux(0, 2, other) {
		t.Fatal("with the closing 고")
	}
	if scanOf("가자라").mainBeforeAux(0, 1, fused) {
		t.Fatal("before letters that begin no auxiliary's stem")
	}
	if scanOf("먹어가").mainBeforeAux(0, 2, whole) || !wordsScan("먹어가", layerCost).mainBeforeAux(0, 2, whole) {
		t.Fatal("가 is an auxiliary's stem only with a word table")
	}
}
