//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestKnownAsNew — a word the table or a closed set knows only as a nominal is no new word's stem before a predicate tail (사과+다, 거+네, 뭐+냐); alive are a predicate's stem (먹, also before a tail it does not take, and 물, a noun and a verb), a word no table has, the same nominal before particles, a tail that names its own stem, and everything without a word table; end to end 거네 and 뭐냐 keep the dropped copula and lose the new word
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestKnownAsNew(t *testing.T) {
	pred := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	if scanOf("사과다").knownAsNew(0, tail{start: 2, pieces: pred}) {
		t.Fatal("without a word table nothing is asked")
	}
	ws := wordsScan("사과다", layerCost)
	if !ws.knownAsNew(0, tail{start: 2, pieces: pred}) {
		t.Fatal("사과, a noun of the table, before a predicate tail")
	}
	if ws.knownAsNew(0, tail{start: 2, pieces: []piece{{text: "다", tail: rule.TailCase}}}) {
		t.Fatal("the same nominal before particles")
	}
	if ws.knownAsNew(0, tail{start: 2, as: numName, pieces: pred}) {
		t.Fatal("a tail that names its own stem")
	}
	if !wordsScan("거네", layerCost).knownAsNew(0, tail{start: 1, pieces: []piece{{text: "네", names: "맺음:네", verdict: 1, tail: rule.TailPred}}}) {
		t.Fatal("거, a dependent noun of a closed set, before a predicate tail")
	}
	if wordsScan("먹다", layerCost).knownAsNew(0, tail{start: 1, pieces: pred}) || wordsScan("물다", layerCost).knownAsNew(0, tail{start: 1, pieces: pred}) {
		t.Fatal("a predicate's stem")
	}
	if wordsScan("먹대", layerCost).knownAsNew(0, tail{start: 1, pieces: []piece{{text: "대", tail: rule.TailAdj}}}) {
		t.Fatal("a verb's stem before an adjective's tail is left to its kind's graph")
	}
	if wordsScan("곰다", layerCost).knownAsNew(0, tail{start: 1, pieces: pred}) {
		t.Fatal("a word no table has")
	}

	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	r.ix.lex["뭐"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	for in, want := range map[string][2]string{
		"거네": {"거[의존명사:거]+(이)[서술격:이(빠짐)]+네[맺음:네]", "거+네["},
		"뭐냐": {"뭐[열쇠:체언]+(이)[서술격:이(빠짐)]+냐[맺음:냐]", "뭐+냐["},
	} {
		got := strings.Join(r.Lines(in, false), "\n")
		if !strings.Contains(got, want[0]) || strings.Contains(got, want[1]) {
			t.Errorf("%s: want %s and no %s\n%s", in, want[0], want[1], got)
		}
	}
	if got := strings.Join(r.Lines("사과는 싫다", false), "\n"); !strings.Contains(got, "사과[열쇠:체언]+는[보조사:는]") || !strings.Contains(got, "싫+다[맺음:다]") {
		t.Fatal("a nominal before particles, a word no table has before a predicate tail\n" + got)
	}
	if got := strings.Join(Lines("거네", false), "\n"); !strings.Contains(got, "거+네[맺음:네]") {
		t.Fatal("without a word table the new word stands\n" + got)
	}
}
