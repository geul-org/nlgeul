//ff:func feature=token type=helper control=iteration dimension=1
//ff:what beginsOf — the places a word that has its stem end at e may begin at (beginnings); every place before e for a scan whose places were not worked out
package token

func (sc *scan) beginsOf(e int) []int {
	if len(sc.begins) == len(sc.s)+1 {
		return sc.begins[e]
	}
	out := make([]int, 0, e)
	for b := 0; b < e; b++ {
		out = append(out, b)
	}
	return out
}
