//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestKeptName — a name of the hand-written table or the table of kinds is kept (○ 가림:원 and 기호:원, → 화살표:오른쪽); a name of the table of shapes (ⓐ 원:로마자), a name the letter is not given and a letter in no table are not
package parse

import "testing"

func TestKeptName(t *testing.T) {
	for _, c := range []struct {
		r    rune
		name string
		want bool
	}{{'○', "가림:원", true}, {'○', "기호:원", true}, {'→', "화살표:오른쪽", true}, {'ⓐ', "원:로마자", false}, {'○', "원:로마자", false}, {'가', "가림:원", false}} {
		if got := keptName(c.r, c.name); got != c.want {
			t.Errorf("keptName(%c, %s) = %v", c.r, c.name, got)
		}
	}
}
