//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestModifiesAt — the letters on the two sides of a place read as a modifier and a known word it modifies: 못 before the verb 와 (오+아), 두 before the noun 손, the adverb 좀 before the adverb 더, the modifying form 할 (하+ㄹ) before the dependent noun 거; not 이 (a determiner, but not a numeral one) before a noun, not 못 before a noun, not 두 before an adverb, not an adverb before a noun, not the modifying form before another noun, and not a noun before 거
package token

import "testing"

func TestModifiesAt(t *testing.T) {
	verb := word{kinds: kindVerb, count: seenEnough, by: [6]int{0, seenEnough}}
	noun := word{kinds: kindNominal, count: 500, by: [6]int{500}}
	adverb := word{kinds: kindAdverb, count: 500, by: [6]int{0, 0, 0, 500}}
	for _, c := range []struct {
		text string
		k    int
		want bool
	}{
		{"못와", 1, true},
		{"두손", 1, true},
		{"좀더", 1, true},
		{"할거", 1, true},
		{"이손", 1, false},
		{"못손", 1, false},
		{"두더", 1, false},
		{"좀손", 1, false},
		{"할손", 1, false},
		{"손거", 1, false},
	} {
		sc := wordsScan(c.text, layerCost)
		sc.ix.lex["오"], sc.ix.lex["손"], sc.ix.lex["좀"], sc.ix.lex["더"] = verb, noun, adverb, adverb
		if got := sc.modifiesAt(0, c.k, len(sc.s)); got != c.want {
			t.Errorf("%s at %d: %v, want %v", c.text, c.k, got, c.want)
		}
	}
}
