//ff:func feature=parse type=reader control=sequence
//ff:what TestBareBase — ㅁ under 하, 되, 이 only; ㄴ and ㄹ under 하, 되, 시, 키, 이 and under an auxiliary verb's stem after an 어/아 form
package parse

import "testing"

func TestBareBase(t *testing.T) {
	if !bareBase([]rune("훌륭함"), jongM) || !bareBase([]rune("됨"), jongM) || bareBase([]rune("가심"), jongM) || bareBase([]rune("사람"), jongM) {
		t.Fatal("ㅁ")
	}
	if !bareBase([]rune("분해한"), jongN) || !bareBase([]rune("가신"), jongN) || !bareBase([]rune("작성해준"), jongN) || !bareBase([]rune("해볼"), jongL) || bareBase([]rune("큰"), jongN) || bareBase([]rune("준"), jongN) {
		t.Fatal("ㄴ ㄹ")
	}
}
