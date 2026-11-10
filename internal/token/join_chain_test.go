//ff:func feature=token type=reader control=sequence
//ff:what TestJoinChain — task.py becomes one atom with its readings; a.b stays two runs; in a.b.com every dot joins; 10:30 is one number read two ways, 25:70 only a ratio, 1,00 two numbers, and 2026.11.10 a version and a date
package token

import "testing"

func TestJoinChain(t *testing.T) {
	sc := scanOf("task.py")
	if a := sc.atoms[0]; a.end != 7 || a.names != "파일명|도메인" || !sc.inside[4] || !sc.inside[5] || sc.inside[7] {
		t.Fatal(a, sc.inside)
	}
	sc = scanOf("ab.cd")
	if a := sc.atoms[0]; a.end != 2 || a.names != "" || sc.atoms[3].end != 5 || sc.inside[2] || sc.inside[3] || !sc.inside[1] {
		t.Fatal(sc.atoms, sc.inside)
	}
	sc = scanOf("a.b.com")
	if a := sc.atoms[0]; a.end != 7 || a.names != "도메인" {
		t.Fatal(sc.atoms)
	}
	sc = scanOf("10:30 25:70 1,00 2026.11.10")
	if a := sc.atoms[0]; a.end != 5 || a.names != "시각|비율" || !a.num {
		t.Fatal(sc.atoms)
	}
	if a := sc.atoms[6]; a.names != "비율" {
		t.Fatal(a)
	}
	if a := sc.atoms[12]; a.end != 13 || a.names != numName || sc.atoms[14].end != 16 {
		t.Fatal(sc.atoms)
	}
	if a := sc.atoms[17]; a.names != "버전|날짜" || !a.num {
		t.Fatal(a)
	}
}
