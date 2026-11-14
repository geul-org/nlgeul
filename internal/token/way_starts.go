//ff:func feature=token type=helper control=iteration dimension=1
//ff:what wayStarts — the line where each way's boxes start when the ways are put one after the other (the first line of each way's first box)
package token

func wayStarts(ways []boxWay) []int {
	starts := make([]int, len(ways))
	at := 0
	for i, w := range ways {
		starts[i] = at
		at += len(w.lines)
	}
	return starts
}
