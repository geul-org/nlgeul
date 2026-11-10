//ff:func feature=rule type=helper control=sequence
//ff:what TestFitsShape — a date fits its shape; too few or too many runs, a wrong head, a run that is not digits, a wrong length and a value out of range do not
package rule

import "testing"

func TestFitsShape(t *testing.T) {
	date := &Shape{Min: 3, Max: 3, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}, {1, 31}}}
	phone := &Shape{Min: 3, Max: 3, Head: "0"}
	if !FitsShape([]string{"2026", "11", "10"}, date) || FitsShape([]string{"2026", "11"}, date) || FitsShape([]string{"2026", "11", "10", "9"}, date) {
		t.Fatal("count")
	}
	if FitsShape([]string{"2026", "13", "10"}, date) || FitsShape([]string{"26", "11", "10"}, date) || FitsShape([]string{"2026", "x1", "10"}, date) || FitsShape([]string{"2026", "+1", "10"}, date) {
		t.Fatal("runs")
	}
	if !FitsShape([]string{"010", "1234", "5678"}, phone) || FitsShape([]string{"110", "1234", "5678"}, phone) {
		t.Fatal("head")
	}
}
