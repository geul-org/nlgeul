//ff:func feature=token type=reader control=iteration dimension=1
//ff:what partsOf — a piece written out as the parts it is fused from: each part has the piece's verdict and kind of tail, and the first its preference
package token

func partsOf(pc piece, parts []split) []piece {
	if len(parts) == 0 {
		return []piece{pc}
	}
	var out []piece
	for i, p := range parts {
		q := piece{p.text, p.name, pc.verdict, 0, pc.tail, false, ""}
		if i == 0 {
			q.pref = pc.pref
		}
		out = append(out, q)
	}
	return out
}
