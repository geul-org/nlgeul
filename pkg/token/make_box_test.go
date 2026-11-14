//ff:func feature=token type=helper control=sequence
//ff:what TestMakeBox — the width inside is the widest row, or the title with two dashes when the title is wider; the count of those left out is a row
package token

import "testing"

func TestMakeBox(t *testing.T) {
	b := makeBox(boxNode(0, 3, "아버지", "아버지", "아버"), 0)
	if b.title != "[0-3] 아버지 ─ 후보경합 2" || len(b.rows) != 2 || b.width != 28 {
		t.Fatal(b)
	}
	short := makeBox(boxNode(0, 12, "아버지가방에들어가신다", "가"), 0)
	if short.width != cellWidth(short.title)+2 || short.width <= cellWidth(short.rows[0]) {
		t.Fatal(short)
	}
	cut := makeBox(boxNode(0, 1, "가", "가", "나", "다"), 1)
	if len(cut.rows) != 2 || cut.rows[1] != "(2개 생략됨)" {
		t.Fatal(cut)
	}
}
