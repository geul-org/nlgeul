//ff:func feature=data type=helper control=iteration dimension=1
//ff:what arrowDirection — the one direction an arrow name reads as (오른쪽, 양쪽, 오른쪽 아래 ...), or empty when none or more than one direction is read
package main

import "strings"

func arrowDirection(name string) string {
	words := strings.Fields(name)
	found := map[string]bool{}
	for i := 0; i < len(words); {
		d, n := directionAt(words, i)
		if d != "" {
			found[d] = true
		}
		i += n
	}
	if len(found) != 1 {
		return ""
	}
	for d := range found {
		return d
	}
	return ""
}
