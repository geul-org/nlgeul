//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what branchesInOrder — whether branches stand best first: a candidate never has a higher verdict, or at the same verdict a higher preference, than the one before it (the order of a candidate with no word table)
package parse

func branchesInOrder(bs []Branch) bool {
	for i := 1; i < len(bs); i++ {
		if bs[i].Verdict > bs[i-1].Verdict || bs[i].Verdict == bs[i-1].Verdict && bs[i].Pref > bs[i-1].Pref {
			return false
		}
	}
	return true
}
