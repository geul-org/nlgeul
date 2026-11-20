//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what wordLabels — the labels of a word's parts (Part.Label: a new word's stem as 새 낱말), joined with " + "
package parse

import "strings"

func wordLabels(w Eojeol) string {
	var labels []string
	for _, p := range w.Parts {
		labels = append(labels, p.Label())
	}
	return strings.Join(labels, " + ")
}
