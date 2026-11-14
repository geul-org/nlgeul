//ff:func feature=token type=helper control=iteration dimension=1
//ff:what nearBegins — the places from lo up to e, other than those inside an atom, from which a word with its stem end at e may begin: every one of them is within newStemMax tokens of e
package token

func (sc *scan) nearBegins(e, lo int) []int {
	var out []int
	for b := lo; b < e; b++ {
		if !sc.inside[b] {
			out = append(out, b)
		}
	}
	return out
}
