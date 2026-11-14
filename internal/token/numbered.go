//ff:func feature=token type=helper control=iteration dimension=1
//ff:what numbered — give each candidate's last word its number of birth: the words in the order they first come in the list, a word that came before (by where it begins and ends and how it is printed) keeping its number; of two words of the same lambda the one with the smaller number is the first; a state of the net (it has states before it) keeps the number it has, which a state carried over from another scan has from that scan
package token

func (sc *scan) numbered(cands []*path) {
	reg := sc.registry()
	for _, p := range cands {
		if p.preds != nil {
			continue
		}
		k := wordKeyOf(p)
		n, ok := reg.born[k]
		if !ok {
			n = reg.next
			reg.born[k] = n
			reg.next++
		}
		p.last.born = n
	}
}
