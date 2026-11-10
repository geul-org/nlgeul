//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFusedPast — 했 갔 봤 됐 셨 hold a fused 었/았; 있 and 가 do not
package parse

import "testing"

func TestFusedPast(t *testing.T) {
	for _, r := range "했갔섰봤줬졌냈셌됐셨" {
		if !fusedPast(r) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range "있가간A" {
		if fusedPast(r) {
			t.Errorf("%c", r)
		}
	}
}
