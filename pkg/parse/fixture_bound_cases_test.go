//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what boundCases — checks each case: the word's first candidate is plain in the sentence read with no rule of a sentence and ruled when it is read with the rules given
package parse

import "testing"

func boundCases(t *testing.T, cases []boundCase, rules ...sentenceRule) {
	t.Helper()
	for _, c := range cases {
		if got := boundFirst(t, c.sentence, c.word); got != c.plain {
			t.Errorf("%s: %s with no rule is %s, want %s", c.sentence, c.word, got, c.plain)
		}
		if got := boundFirst(t, c.sentence, c.word, rules...); got != c.ruled {
			t.Errorf("%s: %s with the rule is %s, want %s", c.sentence, c.word, got, c.ruled)
		}
	}
}
