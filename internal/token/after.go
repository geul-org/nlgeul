//ff:func feature=token type=helper control=iteration dimension=1
//ff:what after — the runs that begin at or after a place (those inside the atom a unit begins with are left out)
package token

func (sc *scan) after(runs [][2]int, from int) [][2]int {
	var out [][2]int
	for _, r := range runs {
		if r[0] >= from {
			out = append(out, r)
		}
	}
	return out
}
