//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what longest — the length of the longest tail reading
package mini

func longest(tails []tailFact) int {
	n := 0
	for _, t := range tails {
		n = max(n, t.size)
	}
	return n
}
