//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestUnitWordAt — a stem that begins right after a run of digits and is a unit of the unit table and then a word written after a unit (년간, 월경, 일전, 점차, 개월간) is one; not at the unit's head, not where no number ends, not a stem that is a unit alone or a unit and some other word (세기, 년대), nor one whose head is no unit (기간)
package token

import "testing"

func TestUnitWordAt(t *testing.T) {
	sc := scanOf("5년간")
	if !sc.numEnds[1] || !sc.unitWordAt(1, "년간") {
		t.Fatal("년간 right after 5", sc.numEnds)
	}
	for _, stem := range []string{"월경", "일전", "점차", "개월간", "년후"} {
		if !sc.unitWordAt(1, stem) {
			t.Error("a unit and a word written after a unit", stem)
		}
	}
	for _, stem := range []string{"년", "세기", "년대", "기간", "년간지", "간"} {
		if sc.unitWordAt(1, stem) {
			t.Error("no unit and word written after a unit", stem)
		}
	}
	if sc.unitWordAt(0, "년간") {
		t.Fatal("at the unit's head")
	}
	if scanOf("가년간").unitWordAt(1, "년간") {
		t.Fatal("no number ends before it")
	}
}
