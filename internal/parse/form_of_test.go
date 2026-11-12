//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFormOf — a form is read into its lead and its syllables (면 whole, ㄴ다 under the final consonant ㄴ, 어/아서 after 어/아/여); letters that are not Hangul syllables are an error that names the column and tells how a form is written
package parse

import "testing"

func TestFormOf(t *testing.T) {
	cases := []struct {
		in   string
		lead int
		form string
	}{
		{"면", leadNone, "면"},
		{"ㄴ다", jongN, "다"},
		{"ㄹ", jongL, ""},
		{"ㅁ", jongM, ""},
		{"ㅂ시다", jongB, "시다"},
		{"ㅆ다", jongSS, "다"},
		{"어/아서", leadA, "서"},
	}
	for _, c := range cases {
		lead, form, err := FormOf(rowFor("x", c.in), 1)
		if err != nil || lead != c.lead || form != c.form {
			t.Errorf("%s: %d %q %v", c.in, lead, form, err)
		}
	}
	for _, bad := range []string{"a", "ㄴa", "다 고", "ㅎ다"} {
		_, _, err := FormOf(rowFor("x", bad), 1)
		errHas(t, err, "표 t.tsv 2번째 줄 2번째 칸(앞말)", "\""+bad+"\"", "어/아")
	}
}
