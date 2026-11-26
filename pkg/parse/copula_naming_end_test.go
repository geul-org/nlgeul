//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestCopulaNamingEnd — 음 and a syllable with ㅁ or ㄻ under it end a naming closing the copula is read after; 기, 고 and a syllable with another final consonant do not
package parse

import "testing"

func TestCopulaNamingEnd(t *testing.T) {
	for r, want := range map[rune]bool{'음': true, '됨': true, '함': true, '삶': true, '기': false, '고': false, '된': false} {
		if got := copulaNamingEnd(r); got != want {
			t.Errorf("copulaNamingEnd(%q) = %v", r, got)
		}
	}
}
