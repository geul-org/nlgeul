//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestTableDeterminerRuleModified — read as a stream with words-alone.tsv (the determiner 현 counted more than the nominal) right after a modifying form and before a noun: with no rule the determiner comes first, the rule "a noun stands right after a modifier" alone puts the nominal first, and with this rule beside it the determiner comes first again (both add the same step, so the table's counts decide)
package token

import "testing"

func TestTableDeterminerRuleModified(t *testing.T) {
	ix := streamReader(t, "words-alone.tsv").ix
	noun := nounModifiedRule(modifiedSkip)
	det := tableDeterminerRule(ix.lex, ix.closed)
	for _, c := range []struct {
		name  string
		rules []sentenceRule
		want  string
	}{
		{"no rule", nil, "현[열쇠:관형사]"},
		{"noun after a modifier", []sentenceRule{noun}, "현[열쇠:체언]"},
		{"both", []sentenceRule{noun, det}, "현[열쇠:관형사]"},
	} {
		if got, _ := aloneFirst(t, "좋은 현 비용이 좋다", "현", c.rules...); got != c.want {
			t.Errorf("%s: 현 is %s, want %s", c.name, got, c.want)
		}
	}
}
