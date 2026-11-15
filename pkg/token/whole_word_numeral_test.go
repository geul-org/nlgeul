//ff:func feature=token type=reader control=sequence
//ff:what TestWholeWordNumeral — a Hangul numeral of more syllables and the piece after it are a word together when the word table has them (여러 and 분: 여러분, 조사 and 관: 조사관), and not when it does not (여러 and 개); a numeral with a letter that is no Hangul syllable (1여러) never is
package token

import "testing"

func TestWholeWordNumeral(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-whole-numeral.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	if !ix.wholeWord("여러", "분") || !ix.wholeWord("조사", "관") {
		t.Fatal("a word of the table")
	}
	if ix.wholeWord("여러", "개") || ix.wholeWord("1여러", "분") || ix.wholeWord("", "분") {
		t.Fatal("no word of the table, or no Hangul numeral")
	}
}
