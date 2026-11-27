//ff:func feature=parse type=reader control=sequence
//ff:what TestLoneUnits — a stem whose letters are a unit of the unit table has a candidate for each line of the table with those letters (m is 미터 and 분), named as the table names it, with the low verdict of a unit no number stands on and the kind that stands behind the words of the word table; a unit written in Hangul syllables is read so only with no piece on it (년, not 년 before 을); a stem that is no unit has none
package parse

import "testing"

func TestLoneUnits(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	on := tail{pieces: []piece{{}}}
	got := ix.loneUnits("m", tail{})
	if len(got) != 2 || got[0].name != "단위:미터" || got[1].name != "단위:분" || got[0].kind != kindBehind || got[0].v != ix.loneUnit || ix.loneUnit <= 0 || ix.loneUnit >= 0.5 {
		t.Fatal(got)
	}
	if got = ix.loneUnits("kg", on); len(got) != 1 || got[0].name != "단위:킬로그램" {
		t.Fatal("a unit of letters that are no Hangul syllables with a piece on it", got)
	}
	if got = ix.loneUnits("년", tail{}); len(got) != 1 || got[0].name != "단위:년" || got[0].kind != kindBehind || got[0].v != ix.loneUnit {
		t.Fatal(got)
	}
	if got = ix.loneUnits("년", on); got != nil {
		t.Fatal("a unit written in Hangul with a piece on it", got)
	}
	if got = ix.loneUnits("km", tail{}); len(got) != 1 || got[0].name != "단위:킬로미터" {
		t.Fatal(got)
	}
	if ix.loneUnits("xyz", tail{}) != nil || ix.loneUnits("", tail{}) != nil {
		t.Fatal("no candidate for a stem that is no unit and no stem")
	}
}
