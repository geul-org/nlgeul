//ff:func feature=parse type=reader control=sequence
//ff:what TestReaderStreamQuotedAdn — right after the closing quotation mark of a quoted sentence that ends in a closing, 는 reads as the modifying 는 first, though with no rules the particle comes first; its preference is exactly bonusSameLetters above what the stream with no rules gives it, the particle stays as it was and no candidate is taken away, and the trace says the rule stood; right after a quoted noun the particle stays first and the trace says 기각 with the reason
package parse

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamQuotedAdn(t *testing.T) {
	const adn, particle = "+는[꾸밈:는]", "+는[보조사:는]"
	got, notes := shutTopicBranches(t, "\"다음을 말한다\"는 뜻이다", "는", true)
	plain, _ := shutTopicBranches(t, "\"다음을 말한다\"는 뜻이다", "는", false)
	if len(got) == 0 || got[0].String() != adn || len(plain) == 0 || plain[0].String() != particle {
		t.Fatalf("the modifying form comes first with the rule (%v), the particle without it (%v)", got, plain)
	}
	after, ok := quotedPref(got, adn)
	before, _ := quotedPref(plain, adn)
	if !ok || math.Abs(after-before-bonusSameLetters) > 1e-9 {
		t.Errorf("%v -> %v, want exactly bonusSameLetters more", before, after)
	}
	rest, kept := quotedPref(got, particle)
	was, _ := quotedPref(plain, particle)
	if !kept || rest != was || len(got) != len(plain) {
		t.Errorf("the particle stays (%v, %v) and nothing is taken away (%d of %d)", was, rest, len(got), len(plain))
	}
	if !strings.Contains(strings.Join(notes, "\n"), adn+"\t"+ruleQuotedAdn+"\t1.00\t통과") {
		t.Errorf("no 통과 line of %s in\n%s", ruleQuotedAdn, strings.Join(notes, "\n"))
	}
	got, notes = shutTopicBranches(t, "\"소득\"는 뜻이다", "는", true)
	if len(got) == 0 || got[0].String() != particle {
		t.Fatalf("the particle stays first after a quoted noun: %v", got)
	}
	if want := adn + "\t" + ruleQuotedAdn + "\t-1.00\t기각: " + whyNotQuoteEnd; !strings.Contains(strings.Join(notes, "\n"), want) {
		t.Errorf("no line %q in\n%s", want, strings.Join(notes, "\n"))
	}
}
