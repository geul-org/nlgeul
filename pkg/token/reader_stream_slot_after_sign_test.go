//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamSlotAfterSign — the letters 였다 right after a closing bracket read as the copula 이 with a fused 었 first, and no candidate of them reads the slot 였 with no stem; right after a stem of 하다 (하였다) the slot 였 is still read
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamSlotAfterSign(t *testing.T) {
	got, _ := copulaBranches(t, "그건 차(1대)였다.", "였다", true)
	if len(got) == 0 || got[0].String() != "+이[서술격:이]+(었)[선어말:ㅆ]+다[맺음:다]" {
		t.Fatalf("the first candidate is the copula, got %v", got)
	}
	for _, b := range got {
		if strings.Contains(b.String(), "였[선어말:였]") {
			t.Errorf("a slot with no stem is read: %v", b)
		}
	}
	done, _ := copulaBranches(t, "그는 하였다.", "하였다", true)
	if len(done) == 0 || done[0].String() != "하[열쇠:동사]+였[선어말:였]+다[맺음:다]" {
		t.Errorf("the slot after 하 is read, got %v", done)
	}
}
