//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestUnmodifiable — no adnominal may stand before a word whose known candidates all take none (an adverb; a predicate and an adverb; with a new word's reading beside them), nor before a word whose first candidate is a full stop, a question mark or an exclamation mark; one may stand before a word with a nominal or modifying reading among its candidates, a word of a new word's reading alone, and a comma
package parse

import "testing"

func TestUnmodifiable(t *testing.T) {
	br := func(parts ...Part) Branch { return Branch{Words: []Eojeol{{Parts: parts}}} }
	node := func(brs ...Branch) Node { return Node{Contest: Contest{Branches: brs}} }
	adverb := br(Part{Text: "역시", Name: Name{Kind: "열쇠", Form: "부사"}, Stem: true})
	verb := br(Part{Text: "가", Name: Name{Kind: "열쇠", Form: "동사"}, Stem: true}, Part{Text: "다", Name: Name{Kind: "맺음", Form: "다"}})
	adn := br(Part{Text: "가", Name: Name{Kind: "열쇠", Form: "동사"}, Stem: true}, Part{Text: "ㄴ", Name: Name{Kind: "꾸밈", Form: "ㄴ"}})
	noun := br(Part{Text: "채", Name: Name{Kind: "열쇠", Form: "체언"}, Stem: true})
	fresh := br(Part{Text: "역시", Stem: true})
	stop := br(Part{Text: ".", Name: Name{Kind: "마침표"}})
	ask := br(Part{Text: "?", Name: Name{Kind: "물음표"}})
	comma := br(Part{Text: ",", Name: Name{Kind: "쉼표"}})
	for _, c := range []struct {
		name string
		n    Node
		want bool
	}{
		{"adverb", node(adverb), true},
		{"predicate and adverb", node(verb, adverb), true},
		{"with a new word beside", node(adverb, fresh), true},
		{"full stop", node(stop), true},
		{"question mark", node(ask), true},
		{"with a nominal", node(adverb, noun), false},
		{"with a modifying form", node(verb, adn), false},
		{"new word alone", node(fresh), false},
		{"comma", node(comma), false},
	} {
		if got := unmodifiable(c.n); got != c.want {
			t.Errorf("%s: %v", c.name, got)
		}
	}
}
