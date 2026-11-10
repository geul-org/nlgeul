//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAuxHead — the letters begin with an auxiliary verb's stem as it stands (주, 버리), with a final consonant under its last syllable (볼, 준), or with 어/아 fused into it (줘, 봐, 둬, 놔, 버려) and ㅆ under that (줬, 봤); not 조 for 주, not 봔, not a stem that only shares the first letters (버티), not letters shorter than the stem, not a sign that is no syllable
package token

import "testing"

func TestAuxHead(t *testing.T) {
	cases := []struct {
		rest, stem string
		want       bool
	}{
		{"주다", "주", true}, {"버리다", "버리", true},
		{"볼까", "보", true}, {"준다", "주", true},
		{"줘", "주", true}, {"봐", "보", true}, {"둬", "두", true}, {"놔", "놓", true}, {"버려", "버리", true},
		{"줬다", "주", true}, {"봤다", "보", true},
		{"조", "주", false}, {"봔", "보", false}, {"버티다", "버리", false}, {"가다", "주", false},
		{"버", "버리", false}, {"x", "주", false}, {"보x", "보x", true}, {"보y", "보x", false},
	}
	for _, c := range cases {
		if got := auxHead([]rune(c.rest), []rune(c.stem)); got != c.want {
			t.Errorf("%s after %s: %v", c.stem, c.rest, got)
		}
	}
}
