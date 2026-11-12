//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestDeterminer — a reading named as a determiner is one, by the closed set (관형사:이) or by the word table's kind (열쇠:관형사); no name, a nominal, another kind of the table and a determiner named as a pronoun are not
package token

import "testing"

func TestDeterminer(t *testing.T) {
	for _, name := range []string{"관형사:이", "열쇠:관형사", "관형사:한"} {
		if !determiner(nameOf(name)) {
			t.Errorf("%s is a determiner", name)
		}
	}
	for _, name := range []string{"", "열쇠:체언", "열쇠:홀로", "대명사:그", "의존명사:것", "꾸밈:는"} {
		if determiner(nameOf(name)) {
			t.Errorf("%q is no determiner", name)
		}
	}
}
