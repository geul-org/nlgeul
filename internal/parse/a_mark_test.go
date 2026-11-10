//ff:func feature=parse type=reader control=sequence
//ff:what TestAMark — a whole 어 is a piece (되+어); a fused one stays with the stem, 하 set off (작성+해)
package parse

import "testing"

func TestAMark(t *testing.T) {
	if aMark([]rune("되어")) != "되+어" || aMark([]rune("작성해")) != "작성+해" || aMark([]rune("해")) != "해" || aMark([]rune("어")) != "어" {
		t.Fatal(aMark([]rune("되어")), aMark([]rune("작성해")), aMark([]rune("어")))
	}
}
