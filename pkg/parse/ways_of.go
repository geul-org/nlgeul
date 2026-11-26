//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what waysOf — the candidates of a sign of a pair in one way the signs before it may stand open, each with the way the signs stand open after it (the two lists judge gives, side by side)
package parse

func waysOf(candidates []pairing, after []string) []pairWay {
	out := make([]pairWay, len(candidates))
	for i, rd := range candidates {
		out[i] = pairWay{rd, after[i]}
	}
	return out
}
