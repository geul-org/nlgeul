//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNewStemMax — the stem of a new word is a candidate only up to newStemMax tokens (the particle and closing after it are not counted): ten letters of no table and a particle are read as a new word and eleven are not, whatever follows; the letters of a unit that are one word of no table stay one new word, bare or with a tail, however long; with no table and with a table
package token

import "testing"

func TestNewStemMax(t *testing.T) {
	if newStemMax != 10 {
		t.Fatal("the limit of a new word's stem", newStemMax)
	}
	for _, words := range []string{"", "words-far.tsv"} {
		r := streamReader(t, words)
		if !unitHas(r, "가나다라마바사아자차를마신다", "가나다라마바사아자차+를[목적격:를]") {
			t.Fatal(words, "ten letters and a particle are a new word")
		}
		if unitHas(r, "가나다라마바사아자차카를마신다", "가나다라마바사아자차카+를[목적격:를]") {
			t.Fatal(words, "eleven letters are not a new word's stem when another word follows")
		}
		if !unitHas(r, "가나다라마바사아자차카를마신다", "라마바사아자차카+를[목적격:를]") {
			t.Fatal(words, "the last letters of the same unit still are")
		}
		if !hasUnitWord(r, "가나다라마바사아자차카타파하거너더러", "가나다라마바사아자차카타파하거너더러") || !unitHas(r, "가나다라마바사아자차카타파하거너더러를", "가나다라마바사아자차카타파하거너더러+를[목적격:를]") {
			t.Fatal(words, "the whole unit as one new word, bare and with a particle")
		}
		if unitHas(r, "가나다라마바사아자차카타파하거너더러를마신다", "가나다라마바사아자차카타파하거너더러+를[목적격:를]") {
			t.Fatal(words, "the same letters followed by another word are not the whole unit")
		}
	}
}
