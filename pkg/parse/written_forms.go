//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what writtenForms — the words given and the wrong ways a closed table writes them (a written form to the word it stands for): the forms a stem of those words may be written in
package parse

import "slices"

func writtenForms(words []string, typos map[string]string) []string {
	out := slices.Clone(words)
	for wrote, right := range typos {
		if slices.Contains(words, right) {
			out = append(out, wrote)
		}
	}
	slices.Sort(out)
	return out
}
