//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what linesInOrder — whether the lines wanted are all among the lines given, in the order wanted (other lines may stand between them)
package parse

func linesInOrder(lines, want []string) bool {
	at := 0
	for _, line := range lines {
		if at < len(want) && line == want[at] {
			at++
		}
	}
	return at == len(want)
}
