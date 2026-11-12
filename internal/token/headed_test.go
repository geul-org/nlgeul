//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestHeaded — a reading begins with the sign of a pair at its unit's head when its first part has the one name 묶음; a part with another name, with two names, with none, a reading with no word and one with no part do not
package token

import "testing"

func TestHeaded(t *testing.T) {
	branch := func(names ...Name) Branch {
		return Branch{Words: []Eojeol{{Parts: []Part{{Text: "(", Names: names}}}}}
	}
	cases := []struct {
		br   Branch
		want bool
	}{
		{branch(Name{Kind: pairName}), true},
		{branch(Name{Kind: openName, Form: "1"}), false},
		{branch(Name{Kind: pairName}, Name{Kind: "열쇠"}), false},
		{branch(), false},
		{Branch{Words: []Eojeol{{}}}, false},
		{Branch{}, false},
	}
	for i, c := range cases {
		if got := headed(c.br); got != c.want {
			t.Errorf("case %d: %v, want %v", i, got, c.want)
		}
	}
}
