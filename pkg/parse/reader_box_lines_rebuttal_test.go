//ff:func feature=parse type=reader control=sequence
//ff:what TestReaderBoxLinesRebuttal — boxes drawn with a rule of rebuttal among the index's rules: the box of 학생이 before 되었다 shows 보격 이 first and the rebutted 주격 이 after it with 성립 0.41 (2/(1+forceUngrammatical)-1), every candidate still shown; with trace the line of the rebuttal comes before the boxes in the form of every rule
package parse

import (
	"strings"
	"testing"
)

func TestReaderBoxLinesRebuttal(t *testing.T) {
	r := streamReader(t, "words-sentence.tsv")
	r.ix.sentence = nil
	plain := r.BoxLines("학생이 되었다", false, 0)
	r.ix.sentence = []sentenceRule{subjectRebuttalRule()}
	got := r.BoxLines("학생이 되었다", false, 0)
	if len(got) != len(plain) || !strings.HasPrefix(plain[1], "│ 학생[열쇠:체언]+이[주격:이]  성립 1.00") || !strings.HasPrefix(got[1], "│ 학생[열쇠:체언]+이[보격:이]  성립 1.00") || !strings.HasPrefix(got[2], "│ 학생[열쇠:체언]+이[주격:이]  성립 0.41") {
		t.Fatal(strings.Join(got, "\n"))
	}
	traced := strings.Join(r.BoxLines("학생이 되었다", true, 0), "\n")
	if !strings.Contains(traced, "  · 학생[열쇠:체언]+이[주격:이]\t되 앞 주격\t1.00\t통과") || strings.Index(traced, "되 앞 주격") > strings.Index(traced, "┌ [0-3] 학생이") {
		t.Fatal(traced)
	}
}
