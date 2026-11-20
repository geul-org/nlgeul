//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestAdnominalAll — each candidate of a word is an adnominal: true for a determiner alone and for a determiner beside a modifying form; false when a candidate is a noun or has no word
package parse

import "testing"

func TestAdnominalAll(t *testing.T) {
	det := Branch{Words: []Eojeol{{Parts: []Part{{Text: "스무", Name: Name{Kind: "열쇠", Form: "관형사"}, Stem: true}}}}}
	adn := Branch{Words: []Eojeol{{Parts: []Part{{Text: "스미", Name: Name{Kind: "열쇠", Form: "동사"}, Stem: true}, {Text: "ㄴ", Name: Name{Kind: "꾸밈", Form: "ㄴ"}}}}}}
	noun := Branch{Words: []Eojeol{{Parts: []Part{{Text: "스무", Name: Name{Kind: "열쇠", Form: "체언"}, Stem: true}}}}}
	for _, c := range []struct {
		name string
		brs  []Branch
		want bool
	}{
		{"determiner", []Branch{det}, true},
		{"determiner and modifying form", []Branch{det, adn}, true},
		{"with a noun", []Branch{det, noun}, false},
		{"with no word", []Branch{det, {}}, false},
	} {
		if got := adnominalAll(Node{Contest: Contest{Branches: c.brs}}); got != c.want {
			t.Errorf("%s: %v", c.name, got)
		}
	}
}
