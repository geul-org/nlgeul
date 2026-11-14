//ff:func feature=token type=helper control=sequence
//ff:what TestCutsOf — a contest gives one way of cutting for each of its branches, in order, from where the stretch begins in the sentence
package token

import "testing"

func TestCutsOf(t *testing.T) {
	w := func(text string, end int) Eojeol { return Eojeol{Parts: []Part{{Text: text, Stem: true}}, End: end} }
	c := Contest{Text: "가나다", Branches: []Branch{
		{Words: []Eojeol{w("가나다", 3)}, Verdict: 1},
		{Words: []Eojeol{w("가", 1), w("나다", 3)}, Verdict: 1},
	}}
	got := cutsOf(c, 4)
	if len(got) != 2 || lineKey(got[0].spans) != "4-7" || lineKey(got[1].spans) != "4-5,5-7" {
		t.Fatal(got)
	}
	if len(cutsOf(Contest{Text: "가"}, 0)) != 0 {
		t.Fatal("a contest with no branch has no way")
	}
}
