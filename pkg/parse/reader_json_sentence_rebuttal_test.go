//ff:func feature=parse type=reader control=sequence
//ff:what TestReaderJSONSentenceRebuttal — a sentence set down as data with a rule of rebuttal among the index's rules: the rebutted 주격 이 of 학생이 before 되었다 is written after 보격 이 with its verdict multiplied by 2/(1+forceUngrammatical)-1, its preference as it was and its lambda lowered by Lambda times the log10 of that, and every candidate is still written
package parse

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
	if !almost(now[1].Verdict, was[0].Verdict*(2/(1+forceUngrammatical)-1)) || now[1].Pref != was[0].Pref || now[1].Lambda == nil || !almost(*now[1].Lambda, *was[0].Lambda+Lambda*math.Log10(2/(1+forceUngrammatical)-1)) {
		t.Fatal(now[1], was[0])
	}
}
