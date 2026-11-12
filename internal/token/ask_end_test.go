//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAskEnd — 지, the last letter of the closings that ask (askClosings), is one; the other letters of their names and the last letters of other closings are not
package token

import "testing"

func TestAskEnd(t *testing.T) {
	if !askEnd('지') {
		t.Fatal("지")
	}
	for _, r := range []rune{'는', '은', '을', 'ㄴ', 'ㄹ', ':', '기', '고', '다', 'x'} {
		if askEnd(r) {
			t.Errorf("%c", r)
		}
	}
}
