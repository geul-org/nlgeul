//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestFarKnownStems — a stem of a table has no limit of length (a new word's stem has, newStemMax): known words of more than ten letters stand in the middle of a unit and at its end, bare and with a tail, as a verb whose last syllable a piece was fused into (…수립하 + 였 = …수립했), as the changed form of a verb's stem (…를잘듣 → …를잘들), as a wrong way to write a word of the table, and as words bound by a slash or a hyphen; the stem of a table's word is not cut by the limit to the new words
package token

import "testing"

func TestFarKnownStems(t *testing.T) {
	r := streamReader(t, "words-far.tsv")
	for _, c := range []struct{ in, want string }{
		{"대한민국임시정부수립기념일을마신다", "대한민국임시정부수립기념일[열쇠:체언]+을[목적격:을] 마시[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]"},
		{"고양이가대한민국임시정부수립기념일을마신다", "고양이[열쇠:체언]+가[주격:가] 대한민국임시정부수립기념일[열쇠:체언]+을[목적격:을]"},
		{"대한민국임시정부수립기념일", "대한민국임시정부수립기념일[열쇠:체언]"},
		{"대한민국임시정부수립했다", "대한민국임시정부수립하[열쇠:동사]+(였)[선어말:ㅆ]+다[맺음:다]"},
		{"대한민국임시정부를잘들었다", "대한민국임시정부를잘듣[열쇠:동사]+었[선어말:었]+다[맺음:다]"},
		{"대한민국임시정부수립기렴일을마신다", "대한민국임시정부수립기념일[열쇠:체언,오타정정:대한민국임시정부수립기렴일]+을[목적격:을]"},
		{"대한민국임시정부/수립기념일을마신다", "대한민국임시정부/수립기념일[대비:/]+을[목적격:을] 마시[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]"},
		{"대한민국임시정부-수립기념일을마신다", "대한민국임시정부-수립기념일[붙임:-]+을[목적격:을] 마시[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]"},
	} {
		if !unitHas(r, c.in, c.want) {
			t.Error(c.in, "has no candidate", c.want)
		}
	}
}
