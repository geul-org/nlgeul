//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNominalEnd — a word ends in a nominal when its last part is a noun of the table, a dependent noun, a number, an order or a new word, or a unit after a number; not when it ends in a particle or a closing, when its last stem is a verb or a determiner, or when it has no part
package token

import "testing"

func TestNominalEnd(t *testing.T) {
	noun := Part{Text: "패소", Name: Name{Kind: "열쇠", Form: "체언"}, Stem: true}
	bound := Part{Text: "수", Name: Name{Kind: "의존명사", Form: "수"}, Stem: true}
	number := Part{Text: "3", Name: Name{Kind: "수"}, Stem: true}
	order := Part{Text: "제1", Name: Name{Kind: "차례"}, Stem: true}
	fresh := Part{Text: "미드필더", Stem: true}
	unit := Part{Text: "차", Name: Name{Kind: "단위", Form: "차"}}
	particle := Part{Text: "의", Name: Name{Kind: "관형격", Form: "의"}}
	verb := Part{Text: "먹", Name: Name{Kind: "열쇠", Form: "동사"}, Stem: true}
	closing := Part{Text: "다", Name: Name{Kind: "맺음", Form: "다"}}
	det := Part{Text: "그", Name: Name{Kind: "관형사", Form: "그"}, Stem: true}
	for _, c := range []struct {
		name  string
		parts []Part
		want  bool
	}{
		{"noun", []Part{noun}, true},
		{"dependent noun", []Part{bound}, true},
		{"number", []Part{number}, true},
		{"order", []Part{order}, true},
		{"new word", []Part{fresh}, true},
		{"unit after a number", []Part{number, unit}, true},
		{"noun and particle", []Part{noun, particle}, false},
		{"verb and closing", []Part{verb, closing}, false},
		{"verb stem", []Part{verb}, false},
		{"determiner", []Part{det}, false},
		{"nothing", nil, false},
	} {
		if got := nominalEnd(c.parts); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
