//ff:func feature=token type=rule control=sequence
//ff:what TestAddFarKeys — the keys of the tables longer than a new word's stem (newStemMax) are noted with their last letter taken off and each of their beginnings, and the longest key of all; the keys of the closed-set words, the word table, the changed forms of its stems and its wrong ways to write a word count; without a table only the closed-set words do
package token

import "testing"

func TestAddFarKeys(t *testing.T) {
	r := streamReader(t, "words-far.tsv")
	const long = "대한민국임시정부수립기념일"
	if !r.ix.farKeys[long] || !r.ix.farBases["대한민국임시정부수립기념"] || !r.ix.keyPrefixes["대한민국임시정부수립기념일"] || !r.ix.keyPrefixes["대한민국임시정부수립"] || r.ix.keyPrefixes["고양이"] {
		t.Fatal("a known word longer than a new word's stem", len(r.ix.farKeys))
	}
	if !r.ix.farKeys["대한민국임시정부수립기렴일"] || !r.ix.farKeys["대한민국임시정부를잘들"] || r.ix.farKeys["대한민국임시정부"] || r.ix.farKeys["고양이"] {
		t.Fatal("a wrong way to write a word and a changed form count, a key of ten letters or fewer does not")
	}
	if r.ix.keyMax != len([]rune(long)) {
		t.Fatal("the longest key", r.ix.keyMax)
	}
	none := streamReader(t, "")
	if len(none.ix.farKeys) != 0 || none.ix.keyMax == 0 {
		t.Fatal("no table: the closed-set words are the keys", len(none.ix.farKeys), none.ix.keyMax)
	}
}
