//ff:func feature=token type=reader control=sequence
//ff:what TestUnion — names of both lists once each, in order
package token

import "testing"

func TestUnion(t *testing.T) {
	if union("", "파일명|도메인") != "파일명|도메인" || union("도메인", "파일명|도메인") != "도메인|파일명" || union("버전", "버전") != "버전" {
		t.Fatal(union("도메인", "파일명|도메인"))
	}
}
