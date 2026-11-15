//ff:func feature=rule type=helper control=sequence
//ff:what nearAt — the i-th of the words on one side of a candidate (0 is the nearest); an empty one when there are not so many
package rule

func nearAt(ns []Near, i int) Near {
	if i >= len(ns) {
		return Near{}
	}
	return ns[i]
}
