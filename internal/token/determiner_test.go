//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestDeterminer — a word named only as a determiner is one, by the closed set (관형사:이), by the word table's kind (열쇠:관형사) or by both; a word with no name, a nominal, a word that is a nominal too (열쇠:체언|관형사) and a determiner named as a pronoun are not
package token

import "testing"

func TestDeterminer(t *testing.T) {
	for _, joined := range []string{"관형사:이", "열쇠:관형사", "관형사:한|관형사:두"} {
		if !determiner(namesOf(joined)) {
			t.Errorf("%s is a determiner", joined)
		}
	}
	for _, joined := range []string{"", "열쇠:체언", "열쇠:체언|관형사", "열쇠:관형사|홀로", "대명사:그", "의존명사:것", "꾸밈:는"} {
		if determiner(namesOf(joined)) {
			t.Errorf("%q is no determiner", joined)
		}
	}
}
