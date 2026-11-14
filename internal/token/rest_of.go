//ff:func feature=token type=helper control=iteration dimension=1
//ff:what restOf — for each node, the most a way can reach from that node on: its own best lambda plus the greatest of the same of the nodes after it (the best lambda alone for a node that ends the region); the nodes after a node have greater places, so one pass from the last back is enough
package token

import "math"

func restOf(best []float64, next [][]int) []float64 {
	rest := make([]float64, len(best))
	for i := len(best) - 1; i >= 0; i-- {
		tail := 0.0
		if len(next[i]) > 0 {
			tail = math.Inf(-1)
		}
		for _, j := range next[i] {
			tail = math.Max(tail, rest[j])
		}
		rest[i] = best[i] + tail
	}
	return rest
}
