//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSortByLambda — candidates come greatest lambda first, those of the same lambda as they came, those with a verdict of 0 last in the order they came
package token

import "testing"

func TestSortByLambda(t *testing.T) {
	mk := func(text string, v, p float64) Branch {
		return Branch{Words: []Eojeol{{Parts: []Part{{Text: text, Stem: true}}}}, Verdict: v, Pref: p}
	}
	bs := []Branch{mk("a", 0, 0), mk("b", 1, -5), mk("c", 0, -1), mk("d", 1, -2), mk("e", 1, -5), mk("f", 0.5, 0)}
	sortByLambda(bs)
	got := ""
	for _, b := range bs {
		got += b.String()
	}
	if got != "dfbeac" {
		t.Fatal(got)
	}
}
