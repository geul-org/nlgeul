//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSortByLambdaBorn — of the same lambda the candidate born first comes first (whatever order they came in), one with no number comes after those with one, the same number stays in the order it came, and a greater lambda still comes first
package token

import "testing"

func TestSortByLambdaBorn(t *testing.T) {
	mk := func(text string, p float64, born int) Branch {
		br := Branch{Words: []Eojeol{{Parts: []Part{{Text: text, Stem: true}}}}, Verdict: 1, Pref: p, scores: []wordScore{{1, p, born}}}
		if born == noBorn {
			br.scores = nil
		}
		return br
	}
	bs := []Branch{mk("a", -2, 5), mk("b", -2, noBorn), mk("c", -2, 1), mk("d", -3, 0), mk("e", -2, 1), mk("f", -1, 9)}
	sortByLambda(bs)
	got := ""
	for _, b := range bs {
		got += b.String()
	}
	if got != "fceabd" {
		t.Fatal(got)
	}
}
