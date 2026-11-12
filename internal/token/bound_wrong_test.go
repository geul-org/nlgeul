//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestBoundWrong — a stem of words bound by a slash or a hyphen with a word the table gives as misspelt is wrong; not one of words that are not, one with no sign, or a word alone
package token

import "testing"

func TestBoundWrong(t *testing.T) {
	ix := typoScan("x", layerCost).ix
	for stem, want := range map[string]bool{"입력/역활": true, "역활-입력": true, "입력/출력": false, "역활": false, "입력출력": false, "입력/역활-입력": false} {
		if got := ix.boundWrong(stem); got != want {
			t.Errorf("%s: %v", stem, got)
		}
	}
	if scanOf("x").ix.boundWrong("입력/역활") {
		t.Fatal("no word table")
	}
}
