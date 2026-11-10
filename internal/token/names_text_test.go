//ff:func feature=token type=helper control=sequence
//ff:what TestNamesText — readings print as kind:form joined with "|", a kind alone with no form, a form that is a colon after the kind's colon, the kinds of a word table's stem under one 열쇠, and nothing for no readings
package token

import "testing"

func TestNamesText(t *testing.T) {
	if got := namesText([]Name{{"주격", "가"}, {"보격", "가"}}); got != "주격:가|보격:가" {
		t.Fatalf("got %q", got)
	}
	if got := namesText([]Name{{"열쇠", "체언"}, {"열쇠", "부사"}}); got != "열쇠:체언|부사" {
		t.Fatalf("got %q", got)
	}
	if got := namesText([]Name{{"소수점", ""}, {"버전", ""}}); got != "소수점|버전" {
		t.Fatalf("got %q", got)
	}
	if got := namesText([]Name{{"주격", "이"}, {"열쇠", "체언"}}); got != "주격:이|열쇠:체언" {
		t.Fatalf("got %q", got)
	}
	if got := namesText([]Name{{"시각", ":"}, {"비율", ":"}}); got != "시각::|비율::" {
		t.Fatalf("got %q", got)
	}
	if got := namesText(nil); got != "" {
		t.Fatalf("no readings: %q", got)
	}
}
