//ff:func feature=parse type=reader control=sequence
//ff:what TestLeftAdn — the longest modifying reading: 아는, 전략인, 된다는; none for 먹고 or 사과
package parse

import "testing"

func TestLeftAdn(t *testing.T) {
	if m, n, ok := leftAdn([]rune("된다는")); !ok || m != "된+다+는" || n != 5 {
		t.Fatal(m, n, ok)
	}
	if m, n, ok := leftAdn([]rune("전략인")); !ok || m != "전략+인" || n != 2 {
		t.Fatal(m, n, ok)
	}
	if m, _, ok := leftAdn([]rune("아는")); !ok || m != "아+는" {
		t.Fatal(m, ok)
	}
	if _, _, ok := leftAdn([]rune("먹고")); ok {
		t.Fatal("먹고")
	}
	if _, _, ok := leftAdn([]rune("사과")); ok {
		t.Fatal("사과")
	}
}
