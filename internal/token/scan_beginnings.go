//ff:func feature=token type=reader control=iteration dimension=1
//ff:what beginnings — for each place e the places from which a word may begin that has its stem end at e, in order: those farther than newStemMax tokens where a word that is not a new word may begin (farBegins), then every place within that many tokens that is not inside an atom; a new word's stem is no longer, and the place 0 is also tried where the word is the whole unit (extend, whole)
package token

func (sc *scan) beginnings() {
	sc.begins = make([][]int, len(sc.s)+1)
	lo := 0
	for e := range sc.begins {
		for sc.count[e]-sc.count[lo] > newStemMax {
			lo++
		}
		sc.begins[e] = sc.farBegins(e, lo)
		sc.begins[e] = append(sc.begins[e], sc.nearBegins(e, lo)...)
	}
}
