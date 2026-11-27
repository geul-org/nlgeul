//ff:func feature=data type=helper control=sequence
//ff:what directionAt — the direction read at a word of an arrow name and how many words it takes (a pair takes two), or empty and one
package main

func directionAt(words []string, i int) (string, int) {
	if i+1 < len(words) {
		if d, ok := directionPairs[[2]string{words[i], words[i+1]}]; ok {
			return d, 2
		}
	}
	return directionWords[words[i]], 1
}
