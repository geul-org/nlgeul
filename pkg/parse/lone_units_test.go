//ff:func feature=parse type=reader control=sequence
//ff:what TestLoneUnits — a stem whose letters are a unit of the unit table written in letters that are no Hangul syllables has one candidate for each line of the table with those letters, named as the table names it (m: 단위:미터 and 단위:분), with the verdict of the graph loneUnitGraph; a unit written in Hangul syllables (명), a stem that is no unit and the empty stem have none
package parse

import "testing"

func TestLoneUnits(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	got := ix.loneUnits("m")
	if len(got) != 2 || got[0].name != "단위:미터" || got[1].name != "단위:분" || got[0].v != ix.loneUnit || got[1].v != ix.loneUnit || ix.loneUnit <= 0 || ix.loneUnit >= 1 {
		t.Fatal(got, ix.loneUnit)
	}
	if got = ix.loneUnits("km"); len(got) != 1 || got[0].name != "단위:킬로미터" {
		t.Fatal(got)
	}
	if ix.loneUnits("명") != nil || ix.loneUnits("xyz") != nil || ix.loneUnits("") != nil {
		t.Fatal("no candidate for a Hangul unit, a stem that is no unit and no stem")
	}
}
