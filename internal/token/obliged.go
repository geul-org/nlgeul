//ff:func feature=token type=reader control=iteration dimension=2
//ff:what obliged — the readings of 겠 right after the closing 어야 (가야겠다, 먹어야겠네: 어야 하겠 with 하 left out): each tail that begins with the slot 겠, read once more after each tail that ends right before it with that closing, the auxiliary 하 written out between them
package token

import "slices"

// The tails are those of the closings read after the slot. The tail
// before the slot keeps its stem's restored syllable, its gate and its
// head; what follows the slot tells of the closing.
func (sc *scan) obliged(tails []tail) []tail {
	var out []tail
	for _, t := range tails {
		if t.base != 0 || t.only != "" || t.shown != "" || t.as != "" || t.irr != 0 || t.start < 2 || len(t.pieces) < 2 || t.pieces[0].names != guessName {
			continue
		}
		rest := append(partsOf(t.pieces[0], []split{mustPart, {t.pieces[0].text, guessName}}), t.pieces[1:]...)
		for _, l := range sc.closings(t.start, false, false) {
			if l.only != "" || !slices.Contains(mustNames, l.pieces[len(l.pieces)-1].names) {
				continue
			}
			l.pieces = append(append([]piece{}, l.pieces...), rest...)
			l.noYo, l.close = t.noYo, t.close
			out = append(out, l)
		}
	}
	return out
}
