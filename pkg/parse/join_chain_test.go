//ff:func feature=parse type=reader control=sequence
//ff:what TestJoinChain — 0x1234 is one number named 수:16진법, 0x12G4 and 0X1F stay runs; a run of the Roman numeral letters of Unicode alone is one number named 수:로마십진 (Ⅱ, ⅩⅢ, with the prefix 제 before it too), and ⅡA stays a run; task.py becomes one atom with its candidates; a.b stays two runs; in a.b.com every dot joins; 10:30 is one number read two ways, 25:70 only a ratio, 1,00 two numbers, 2026.11.10 a date and a version, 1.28 with a full stop and a letter right after it a date before a decimal and a version, and 1.28 with a particle right after it a decimal and a version only; a hidden name before a slash is a path with what the slash joins (.cache/db)
package parse

import "testing"

func TestJoinChain(t *testing.T) {
	sc := scanOf("task.py")
	if a := sc.atoms[0]; a.end != 7 || atomNames(a) != "파일명,도메인" || !sc.inside[4] || !sc.inside[5] || sc.inside[7] {
		t.Fatal(a, sc.inside)
	}
	sc = scanOf(".cache/db")
	if a := sc.atoms[0]; a.end != 9 || atomNames(a) != "경로" {
		t.Fatal("a hidden name before a slash", sc.atoms)
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
	if a := sc.atoms[17]; atomNames(a) != "날짜,버전" || !a.num {
		t.Fatal(a)
	}
	sc = scanOf("1.28.은 1.28에")
	if a := sc.atoms[0]; a.end != 4 || atomNames(a) != "날짜,소수점,버전,시각" || !a.num {
		t.Fatal("a full stop right after two numbers: a date first", a)
	}
	if a := sc.atoms[7]; atomNames(a) != "소수점,버전,날짜,시각" {
		t.Fatal("no full stop after them: the decimal first, the date and the time behind", a)
	}
	sc = scanOf("Ⅱ ⅩⅢ ⅡA 제Ⅱ")
	if a := sc.atoms[0]; a.end != 1 || atomNames(a) != romanNumName || !a.num || !sc.numEnds[1] {
		t.Fatal("a Roman numeral letter of Unicode is a number", sc.atoms)
	}
	if a := sc.atoms[2]; a.end != 4 || atomNames(a) != romanNumName || !a.num {
		t.Fatal("a run of Roman numeral letters is one number", sc.atoms)
	}
	if a := sc.atoms[5]; a.end != 7 || a.num || len(a.names) != 0 {
		t.Fatal("a run with another letter stays a run", sc.atoms)
	}
	if a, ok := sc.atoms[8]; !ok || a.end != 10 || !a.num || !sc.ordEnds[10] {
		t.Fatal("the prefix 제 before a Roman numeral letter", sc.atoms)
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
