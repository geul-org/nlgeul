//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestFormShaped — a letter of the tables of special letters has the shape its line gives (Ⅲ and ⅳ 로마 숫자, ① 원) and no other; a letter the tables lack has none
package parse

import "testing"

func TestFormShaped(t *testing.T) {
	for _, c := range []struct {
		r     rune
		shape string
		want  bool
	}{{'Ⅲ', romanShape, true}, {'ⅳ', romanShape, true}, {'①', "원", true}, {'①', romanShape, false}, {'I', romanShape, false}, {'가', "원", false}} {
		if got := formShaped(c.r, c.shape); got != c.want {
			t.Errorf("formShaped(%c, %s) = %v", c.r, c.shape, got)
		}
	}
}
