//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderBundlesQuotes — a sentence with two pairs of quotes after digits and letters: the second pair opens at depth 1 and closes at depth 1 (the quote after the first digit leaves the ways it stood open); feet and inches stand with no word table (5'10": the unit of feet beside the pair's candidates, the quote after the ten a bundle of its own); the old whole-unit candidate stays
package token

import "testing"

func TestReaderBundlesQuotes(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	two := r.Bundles("값은 \"6\"이고 이름은 \"가\"다")
	var quotes []Bundle
	for _, b := range two {
		if b.Text == "\"" {
			quotes = append(quotes, b)
		}
	}
	if len(quotes) != 4 {
		t.Fatal(two)
	}
	opens := kindsOf(quotes[2])
	if opens[0] != "열기:1" || opens[1] != "닫기:1" {
		t.Fatal("the second pair opens at depth 1", opens)
	}
	if shut := kindsOf(quotes[3]); shut[0] != "열기:1" || shut[1] != "닫기:1" {
		t.Fatal("and closes at depth 1", shut)
	}
	feet := r.Bundles("5'10\"")
	if len(feet) != 2 || feet[0].Text != "5'10" || feet[1].Text != "\"" {
		t.Fatal(feet)
	}
	var unit, whole bool
	for _, br := range feet[0].Branches {
		if len(br.Words) == 2 && len(br.Words[0].Parts) == 2 {
			unit = unit || br.Words[0].Parts[1].Name.Form == "피트" && br.Words[0].Parts[0].Text == "5"
		}
		whole = whole || len(br.Words) == 1 && br.Words[0].Parts[0].Text == "5'10"
	}
	if !unit || !whole {
		t.Fatal("feet with no word table, and the whole unit", feet[0])
	}
	if kinds := kindsOf(feet[1]); kinds[0] != "열기:1" || kinds[1] != "단위:인치" {
		t.Fatal(kinds)
	}
}
