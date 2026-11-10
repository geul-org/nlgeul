//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkAfterWord — applies when letters stand left of the sign
package rule

import "testing"

func TestMarkAfterWord(t *testing.T) {
	yes, _ := MarkAfterWord(markCtx(Mark{Left: SideHangul}), nil)
	no, _ := MarkAfterWord(markCtx(Mark{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
