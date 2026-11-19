//ff:func feature=token type=reader control=sequence
//ff:what TestReaderJSONSentenceRebuttal — a sentence set down as data with a rule of rebuttal among the index's rules: the rebutted 주격 이 of 학생이 before 되었다 is written after 보격 이 with a third of its verdict, its preference as it was and its lambda lowered by about 7.16, and every candidate is still written
package token

import (
	"math"
	"testing"
)

func TestReaderJSONSentenceRebuttal(t *testing.T) {
	r := streamReader(t, "words-sentence.tsv")
	r.ix.sentence = nil
	plain := r.JSONSentence("학생이 되었다", 0)
	r.ix.sentence = []sentenceRule{subjectRebuttalRule()}
	d := r.JSONSentence("학생이 되었다", 0)
	was, now := plain.Segments[0].Words[0].Candidates, d.Segments[0].Words[0].Candidates
	if len(now) != len(was) || len(was) < 2 || now[0].Parts[1].Name == nil || *now[0].Parts[1].Name != "보격:이" || *now[1].Parts[1].Name != "주격:이" || *was[0].Parts[1].Name != "주격:이" {
		t.Fatal(was, now)
	}
	if !almost(now[1].Verdict, was[0].Verdict/3) || now[1].Pref != was[0].Pref || now[1].Lambda == nil || !almost(*now[1].Lambda, *was[0].Lambda+Lambda*math.Log10(1.0/3)) {
		t.Fatal(now[1], was[0])
	}
}
