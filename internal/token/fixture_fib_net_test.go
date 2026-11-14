//ff:func feature=token type=helper control=iteration dimension=1
//ff:what fibNet — a lattice made up over n letters where a word of one letter and a word of two letters begin at every place (a letter's words follow every word that ends there): the ways of cutting it are the ways to write n as a sum of 1 and 2, a Fibonacci number
package token

func fibNet(n int) (*scan, *lattice) {
	sc := &scan{}
	root := sc.registry().root
	ends := make([][]*path, n+1)
	ends[0] = []*path{root}
	for i := 0; i < n; i++ {
		ends[i+1] = append(ends[i+1], addState(sc, ends[i], wordStep(i+1, string(rune('가'+i))+"1", 1, -1)))
		if i+2 <= n {
			ends[i+2] = append(ends[i+2], addState(sc, ends[i], wordStep(i+2, string(rune('가'+i))+"2", 1, -1)))
		}
	}
	return sc, &lattice{n: n, ends: ends[n], all: sc.registry().all}
}
