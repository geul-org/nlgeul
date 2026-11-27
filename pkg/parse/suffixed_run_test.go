//ff:func feature=parse type=helper control=sequence
//ff:what TestSuffixedRun — a step is a run where its stem is named as a run of Latin letters and digits no table lists is (열:영숫자: OTS, VAN); not with a stem of another name
package parse

import "testing"

func TestSuffixedRun(t *testing.T) {
	named := func(kind, text string) step {
		return step{word: Eojeol{Parts: []Part{{Text: text, Stem: true, Name: Name{Kind: kind}}}}}
	}
	if !suffixedRun(named("열", "OTS")) {
		t.Fatal("a run of Latin letters")
	}
	if suffixedRun(named("수", "3")) || suffixedRun(named("열쇠", "사업")) {
		t.Fatal("a stem of another name")
	}
}
