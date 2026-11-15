//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lambdaSum — the sum of the lambdas of the first candidate of each node of a way
package token

func lambdaSum(seg Segment) float64 {
	sum := 0.0
	for _, n := range seg.Nodes {
		sum += n.Contest.Branches[0].Lambda()
	}
	return sum
}
