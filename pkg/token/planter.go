//ff:type feature=token type=model
//ff:what planter — the whole unit of a sign of a pair with letters written onto it, read place by place beside the letters after the sign (the places of the unit are those of the letters and the head): its scan, which makes only the words that begin at the unit's head or after the sign, the states it keeps at each place, the sign's own preference, and the states that are the sign alone (the sign takes no part of the word: the words after it are carried over)
package token

type planter struct {
	sc          *scan
	best        [][]*path
	head        int
	signPref    float64
	transparent map[*path]bool
}
