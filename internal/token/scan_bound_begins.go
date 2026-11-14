//ff:func feature=token type=reader control=iteration dimension=1
//ff:what boundBegins — the places before lo from which a stem that ends at e may begin when it is words bound by one sign, a slash or a hyphen (partsBound): the stem's last word is the letters after the last sign and fits (lastPartFits); each word before it, as far back as they all stand, fits (partFits) and is parted from the one before by the same sign; the first word is a part of one of them; none for letters with no sign before e
package token

func (sc *scan) boundBegins(e, lo int) []int {
	if sc.parts == nil {
		return nil
	}
	x := sc.parts[e]
	if x == 0 || x == e || !sc.lastPartFits(x, e) {
		return nil
	}
	var out []int
	for end := x - 1; end > 0; {
		y := sc.parts[end]
		if y == end {
			break
		}
		out = append(out, sc.firstBegins(y, end, lo)...)
		if y == 0 || sc.s[y-1] != sc.s[x-1] || !sc.partFits(y, end) {
			break
		}
		end = y - 1
	}
	return out
}
