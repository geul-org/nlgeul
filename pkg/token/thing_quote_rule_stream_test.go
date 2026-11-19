//ff:func feature=token type=rule control=sequence
//ff:what TestThingQuoteRuleStream — read as a stream with words-quoted.tsv: with no rule 것이라고 after 좋은 is 것 with the particle 이라고; with the rule it is 것 with the copula and -라고, the trace says 통과 for the particle, and the particle is kept behind with a third of its verdict; 책이라고 is not asked
package token

import (
	"strings"
	"testing"
)

func TestThingQuoteRuleStream(t *testing.T) {
	const quote, copula = "것[의존명사:것]+이라고[부사격:이라고]", "것[의존명사:것]+이[서술격:이]+라고[맺음:라고]"
	r := thingQuoteRule()
	if got, _ := ruledFirst(t, "words-quoted.tsv", "좋은 것이라고 했다", "것이라고"); got != quote {
		t.Fatal("with no rule", got)
	}
	got, notes := ruledFirst(t, "words-quoted.tsv", "좋은 것이라고 했다", "것이라고", r)
	if got != copula {
		t.Fatal("with the rule", got)
	}
	auxLinkedHas(t, notes, "  · "+quote+"\t"+ruleThingQuote+"\t1.00\t통과")
	branches := func(rules ...sentenceRule) []Branch {
		st := streamReader(t, "words-quoted.tsv").Stream("좋은 것이라고 했다")
		st.rules = rules
		st.Next()
		seg, _ := st.Next()
		return seg.Nodes[0].Contest.Branches
	}
	was, _ := branchNamed(Node{Contest: Contest{Branches: branches()}}, quote)
	now, k := branchNamed(Node{Contest: Contest{Branches: branches(r)}}, quote)
	if k < 1 || !almost(now.Verdict, was.Verdict/3) || now.Pref != was.Pref {
		t.Fatal("the particle is kept behind", k, now.Verdict, was.Verdict)
	}
	if _, notes := ruledFirst(t, "words-quoted.tsv", "책이라고 했다", "책이라고", r); strings.Contains(strings.Join(notes, "\n"), "\t"+ruleThingQuote+"\t") {
		t.Fatal("책이라고 is not asked", notes)
	}
}
