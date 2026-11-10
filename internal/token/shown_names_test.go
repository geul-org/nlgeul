//ff:func feature=token type=helper control=sequence
//ff:what TestShownNames — readings are shown without their sign, a thousands separator as a number, and none twice
package token

import "testing"

func TestShownNames(t *testing.T) {
	if got := shownNames("시각::|비율::"); got != "시각|비율" {
		t.Fatal(got)
	}
	if got := shownNames("자릿수:,"); got != numName {
		t.Fatal(got)
	}
	if got := shownNames("버전:.|IP:."); got != "버전|IP" {
		t.Fatal(got)
	}
}
