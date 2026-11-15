//ff:func feature=token type=reader control=sequence
//ff:what TestGatedJoined — every gate is told whether the stem and the piece right after it together are a word of the table: the gate of a shortened particle passes a stem they make no word with and turns down one they do, the nominal gate passes both; a tail with a gate and no piece is asked with nothing joined
package token

import "testing"

func TestGatedJoined(t *testing.T) {
	sc := wordsScan("학교서 요구서", layerCost)
	stem := func(text string) step {
		return step{word: Eojeol{Parts: []Part{{Text: text, Stem: true}}}}
	}
	joined := joinedGateOf(joinedGate, 0)
	sc.ix.lex["요구"] = word{kinds: kindNominal, count: 100}
	sc.ix.lex["요구서"] = word{kinds: kindNominal, count: 10}
	seo := []piece{{text: "서", names: "부사격:서"}}
	if !sc.gated(0, tail{start: 2, gate: joined, pieces: seo}, stem("학교")) {
		t.Fatal("학교서 is no word: 서 stands on 학교")
	}
	if sc.gated(4, tail{start: 6, gate: joined, pieces: seo}, stem("요구")) {
		t.Fatal("요구서 is a word: 서 does not stand on 요구")
	}
	if !sc.gated(4, tail{start: 6, gate: sc.ix.nominal, pieces: seo}, stem("요구")) {
		t.Fatal("the nominal gate does not ask whether the two make a word")
	}
	if !sc.gated(4, tail{start: 6, gate: joined}, stem("요구")) {
		t.Fatal("a tail with no piece is asked with nothing joined")
	}
}
