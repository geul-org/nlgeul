//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNamingEnd — 기, the last letter of the naming closing a particle chain is read right after (nounHeads), is one; 음 and ㅁ, naming closings that are not among them, the other letters of the closing's name and letters of other closings are not
package token

import "testing"

func TestNamingEnd(t *testing.T) {
	if !namingEnd('기') {
		t.Fatal("기")
	}
	for _, r := range []rune{'음', 'ㅁ', '름', ':', '이', '다', '지', 'x'} {
		if namingEnd(r) {
			t.Errorf("%c", r)
		}
	}
}
