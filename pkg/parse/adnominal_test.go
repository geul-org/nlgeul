//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestAdnominal — a candidate is an adnominal when its last part is a modifying closing (착하+ㄴ, a new stem with 는) or when it is a determiner's stem standing alone (closed set, word table); not a determiner's letters with a marker after them, a noun, a predicate with another closing, a particle alone, nor nothing
package parse

import "testing"

func TestAdnominal(t *testing.T) {
	adj := Part{Text: "착하", Name: Name{Kind: "열쇠", Form: "형용사"}, Stem: true}
	adn := Part{Text: "ㄴ", Name: Name{Kind: "꾸밈", Form: "ㄴ"}}
	end := Part{Text: "다", Name: Name{Kind: "맺음", Form: "다"}}
	fresh := Part{Text: "쇼", Stem: true}
	det := Part{Text: "그", Name: Name{Kind: "관형사", Form: "그"}, Stem: true}
	tableDet := Part{Text: "새", Name: Name{Kind: "열쇠", Form: "관형사"}, Stem: true}
	noun := Part{Text: "책", Name: Name{Kind: "열쇠", Form: "체언"}, Stem: true}
	particle := Part{Text: "이", Name: Name{Kind: "주격", Form: "이"}}
	for _, c := range []struct {
		name  string
		parts []Part
		want  bool
	}{
		{"modifying form", []Part{adj, adn}, true},
		{"new stem with a modifying form", []Part{fresh, adn}, true},
		{"determiner of the closed set", []Part{det}, true},
		{"determiner of the table", []Part{tableDet}, true},
		{"determiner's letters with a particle", []Part{det, particle}, false},
		{"noun", []Part{noun}, false},
		{"predicate", []Part{adj, end}, false},
		{"particle alone", []Part{particle}, false},
		{"nothing", nil, false},
	} {
		if got := adnominal(c.parts); got != c.want {
			t.Errorf("%s: %v", c.name, got)
		}
	}
}
