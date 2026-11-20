//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what auxLinkedCount — how many of the trace lines are of the two rules of the linking form (their names begin with 문장:연결 꼴 뒤)
package parse

import "strings"

func auxLinkedCount(notes []string) int {
	n := 0
	for _, l := range notes {
		if strings.Contains(l, "\t문장:연결 꼴 뒤") {
			n++
		}
	}
	return n
}
