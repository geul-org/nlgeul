//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestModifiable — an adnominal may stand before a word whose first part is a nominal stem (a noun with or without particles, a number, a new word), an order or a determiner, or that ends in a modifying or a naming form; not before an adverb, a verb with another closing, a sign, a word that begins with no stem, or nothing
package parse

import "testing"

func TestModifiable(t *testing.T) {
	noun := Part{Text: "사람", Name: Name{Kind: "열쇠", Form: "체언"}, Stem: true}
	particle := Part{Text: "이", Name: Name{Kind: "주격", Form: "이"}}
	number := Part{Text: "3", Name: Name{Kind: "수"}, Stem: true}
	order := Part{Text: "제1", Name: Name{Kind: "차례"}, Stem: true}
	fresh := Part{Text: "구룡", Stem: true}
	det := Part{Text: "그", Name: Name{Kind: "관형사", Form: "그"}, Stem: true}
	tableDet := Part{Text: "새", Name: Name{Kind: "열쇠", Form: "관형사"}, Stem: true}
	adj := Part{Text: "착하", Name: Name{Kind: "열쇠", Form: "형용사"}, Stem: true}
	adn := Part{Text: "ㄴ", Name: Name{Kind: "꾸밈", Form: "ㄴ"}}
	naming := Part{Text: "기", Name: Name{Kind: "이름", Form: "기"}}
	end := Part{Text: "다", Name: Name{Kind: "맺음", Form: "다"}}
	adverb := Part{Text: "및", Name: Name{Kind: "열쇠", Form: "부사"}, Stem: true}
	comma := Part{Text: ",", Name: Name{Kind: "쉼표"}}
	for _, c := range []struct {
		name  string
		parts []Part
		want  bool
	}{
		{"noun", []Part{noun}, true},
		{"noun and particle", []Part{noun, particle}, true},
		{"number", []Part{number}, true},
		{"order", []Part{order}, true},
		{"new word", []Part{fresh}, true},
		{"determiner", []Part{det}, true},
		{"determiner of the table", []Part{tableDet}, true},
		{"modifying form", []Part{adj, adn}, true},
		{"naming form", []Part{adj, naming}, true},
		{"adverb", []Part{adverb}, false},
		{"predicate", []Part{adj, end}, false},
		{"sign", []Part{comma}, false},
		{"particle alone", []Part{particle}, false},
		{"nothing", nil, false},
	} {
		if got := modifiable(c.parts); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
