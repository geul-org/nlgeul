//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closings — every predicate tail that ends at j: a closing, alone or with the slots before it, a closing read by the own last syllable of a changed stem, the tails of 지 and 하 written onto a main verb's 어/아 form, the tails of 잖 (지+않) and of 겠 right after 어야 (어야 하겠), and the same after the copula stem, said, dropped or fused into a dependent noun (테)
package token

// more tells that 요 follows, chained that a particle is being read
// right after the tail (the copula's graph is told so). A closing that sits in the stem's last
// syllable takes no slot before it. A closing that attaches only to
// the copula stem stands otherwise only after 아니.
func (sc *scan) closings(j int, more, chained bool) []tail {
	if j < 1 {
		return nil
	}
	pieces, from, starts := sc.readGroup(classClosing, j, -1, more)
	var out []tail
	for _, k := range sc.groupKeys(pieces, from, j) {
		out = append(out, sc.closingTails(pieces[k], from[k], starts[k])...)
	}
	out = append(out, sc.irrTails(j, more)...)
	out = append(out, sc.loneHonors(j, false)...)
	out = append(out, sc.helped(out)...)
	out = append(out, sc.negated(out)...)
	out = append(out, sc.obliged(out)...)
	after := append(sc.copulas(out, chained), sc.nounCopulas(out)...)
	return append(out, after...)
}
