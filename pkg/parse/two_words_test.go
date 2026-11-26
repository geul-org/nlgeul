//ff:func feature=parse type=helper control=sequence
//ff:what TestTwoWords — a stem and a suffix are two words of the table where a nominal of two syllables or more stands first and the rest with the suffix is a word (기동+전사); not where the first word is no nominal (올리+전사), where the rest with the suffix is no word (보장+사), or where only one syllable would stand first (정+당 is the whole word's business)
package parse

import "testing"

func TestTwoWords(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	if !ix.twoWords("기동전", "사") {
		t.Fatal("기동 and 전사")
	}
	if ix.twoWords("올리전", "사") {
		t.Fatal("a verb's stem before the word")
	}
	if ix.twoWords("보장", "사") || ix.twoWords("신문", "사") || ix.twoWords("정", "당") || ix.twoWords("", "사") {
		t.Fatal("no two words")
	}
}
