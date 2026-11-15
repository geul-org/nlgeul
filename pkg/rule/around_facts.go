//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what aroundFacts — the yes/no facts of a candidate judged in its sentence as flags (Around…, in the order of their constants), for AroundIs
package rule

func aroundFacts(a Around) int {
	have := 0
	for i, yes := range []bool{a.Self.Glued, nearAt(a.Before, 0).Nominal} {
		if yes {
			have |= 1 << i
		}
	}
	return have
}
