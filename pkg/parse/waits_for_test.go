//ff:func feature=parse type=helper control=sequence
//ff:what TestWaitsFor — a candidate waits only for a rule that speaks of it (about) and waits for it; a rule with no waits, one that does not speak of it or one whose waits is false for it does not make it wait, nor does no rule
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestWaitsFor(t *testing.T) {
	yes := func(rule.Near) bool { return true }
	no := func(rule.Near) bool { return false }
	self := rule.Near{Stem: "학생", Last: "보격:이"}
	if waitsFor(self, nil) || waitsFor(self, []sentenceRule{{about: yes}, {about: no, waits: yes}, {about: yes, waits: no}}) {
		t.Fatal("no rule waits for it")
	}
	if !waitsFor(self, []sentenceRule{{about: yes}, {about: yes, waits: yes}}) {
		t.Fatal("a rule waits for it")
	}
	if !waitsFor(self, []sentenceRule{waitingComplementRule()}) || waitsFor(rule.Near{Last: "주격:이"}, []sentenceRule{waitingComplementRule()}) {
		t.Fatal("waitingComplementRule waits for 보격 only")
	}
}
