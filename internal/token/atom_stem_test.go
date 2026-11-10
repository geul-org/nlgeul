//ff:func feature=token type=reader control=sequence
//ff:what TestAtomStem — a stem that is one joined atom has its readings and verdict 1; one run the run's name and verdict; anything else nothing; a number has its name and the number's verdict
package token

import "testing"

func TestAtomStem(t *testing.T) {
	sc := scanOf("task.py에 L2")
	if name, v := sc.atomStem(0, 7); name != "파일명|도메인" || v != 1 {
		t.Fatal(name, v)
	}
	if name, v := sc.atomStem(9, 11); name != runName || v != sc.ix.run {
		t.Fatal(name, v)
	}
	if _, v := sc.atomStem(0, 8); v != 0 {
		t.Fatal("an atom and more")
	}
	if _, v := sc.atomStem(7, 8); v != 0 {
		t.Fatal("no atom")
	}
	num := scanOf("30개")
	if name, v := num.atomStem(0, 2); name != numName || v != 1 {
		t.Fatal(name, v)
	}
}
