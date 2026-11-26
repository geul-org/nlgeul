//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what shortOrigin — is a closing with a particle after it (지는, 고는) what a closing of the table of shortened closings is short for (진, 곤)
package parse

func shortOrigin(pair string) bool {
	for _, row := range shortClosings {
		if row.origin == pair {
			return true
		}
	}
	return false
}
