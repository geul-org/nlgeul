//ff:func feature=token type=reader control=sequence
//ff:what TestJoinChain — 0x1234 is one number named 수:16진법, 0x12G4 and 0X1F stay runs; task.py becomes one atom with its readings; a.b stays two runs; in a.b.com every dot joins; 10:30 is one number read two ways, 25:70 only a ratio, 1,00 two numbers, and 2026.11.10 a version and a date
package token

import "testing"

func TestJoinChain(t *testing.T) {
	sc := scanOf("task.py")
	if a := sc.atoms[0]; a.end != 7 || atomNames(a) != "파일명,도메인" || !sc.inside[4] || !sc.inside[5] || sc.inside[7] {
		t.Fatal(a, sc.inside)
	}
	sc = scanOf("ab.cd")
	if a := sc.atoms[0]; a.end != 2 || atomNames(a) != "" || sc.atoms[3].end != 5 || sc.inside[2] || sc.inside[3] || !sc.inside[1] {
		t.Fatal(sc.atoms, sc.inside)
	}
	sc = scanOf("a.b.com")
	if a := sc.atoms[0]; a.end != 7 || atomNames(a) != "도메인" {
		t.Fatal(sc.atoms)
	}
	sc = scanOf("10:30 25:70 1,00 2026.11.10")
	if a := sc.atoms[0]; a.end != 5 || atomNames(a) != "시각,비율" || !a.num {
		t.Fatal(sc.atoms)
	}
	if a := sc.atoms[6]; atomNames(a) != "비율" {
		t.Fatal(a)
	}
	if a := sc.atoms[12]; a.end != 13 || atomNames(a) != numName || sc.atoms[14].end != 16 {
		t.Fatal(sc.atoms)
	}
	if a := sc.atoms[17]; atomNames(a) != "버전,날짜" || !a.num {
		t.Fatal(a)
	}
	sc = scanOf("0x1234 0x12G4 0X1F")
	if a := sc.atoms[0]; a.end != 6 || atomNames(a) != hexName || !a.num || !sc.numEnds[6] {
		t.Fatal("0x and hexadecimal digits is one number", sc.atoms)
	}
	if a := sc.atoms[7]; a.end != 13 || atomNames(a) != "" || a.num {
		t.Fatal("a letter outside a-f makes it a run", a)
	}
	if a := sc.atoms[14]; atomNames(a) != "" || a.num {
		t.Fatal("an upper case 0X is no prefix", a)
	}
}
