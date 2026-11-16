//ff:func feature=token type=reader control=sequence
//ff:what TestReaderStreamShutTopic — right after a closing quotation mark 이란 reads as the topic particle first, though with no rules the noun 이란 of the word table comes first; its preference is exactly bonusOverCount above what the stream with no rules gives it, the noun stays as it was and no candidate is taken away, and the trace says the rule stood; with no closing sign before it (소득이란) the particle gains nothing and the trace says 기각 with the reason that the word before is no closing sign
package token

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamShutTopic(t *testing.T) {
	const in, topic, noun = "\"소득\"이란 다음을 말한다", "+이란[보조사:이란]", "이란[열쇠:체언]"
	got, notes := shutTopicBranches(t, in, "이란", true)
	plain, _ := shutTopicBranches(t, in, "이란", false)
	if len(got) == 0 || got[0].String() != topic || len(plain) == 0 || plain[0].String() != noun {
		t.Fatalf("the particle comes first with the rule (%v), the noun without it (%v)", got, plain)
	}
	after, ok := quotedPref(got, topic)
	before, _ := quotedPref(plain, topic)
	if !ok || math.Abs(after-before-bonusOverCount) > 1e-9 {
		t.Errorf("%v -> %v, want exactly bonusOverCount more", before, after)
	}
	rest, kept := quotedPref(got, noun)
	was, _ := quotedPref(plain, noun)
	if !kept || rest != was || len(got) != len(plain) {
		t.Errorf("the noun stays (%v, %v) and nothing is taken away (%d of %d)", was, rest, len(got), len(plain))
	}
	if !strings.Contains(strings.Join(notes, "\n"), topic+"\t"+ruleShutTopic+"\t1.00\t통과") {
		t.Errorf("no 통과 line of %s in\n%s", ruleShutTopic, strings.Join(notes, "\n"))
	}
	const bare, cand = "소득이란 다음을 말한다", "소득[열쇠:체언]+이란[보조사:이란]"
	got, notes = shutTopicBranches(t, bare, "소득이란", true)
	plain, _ = shutTopicBranches(t, bare, "소득이란", false)
	a, found := quotedPref(got, cand)
	b, _ := quotedPref(plain, cand)
	if !found || a != b {
		t.Errorf("%s gains nothing with no closing sign before it (%v, %v)", cand, b, a)
	}
	if want := cand + "\t" + ruleShutTopic + "\t-1.00\t기각: " + whyNotShut; !strings.Contains(strings.Join(notes, "\n"), want) {
		t.Errorf("no line %q in\n%s", want, strings.Join(notes, "\n"))
	}
}
