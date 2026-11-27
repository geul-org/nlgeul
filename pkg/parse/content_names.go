//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what contentNames — the names (shape:content) of the lines of the tables of special letters that write the content asked for, each name once, in the order of the lines
package parse

import "slices"

func contentNames(list []form, content string) []string {
	var out []string
	for _, f := range list {
		if f.content == content && !slices.Contains(out, f.name()) {
			out = append(out, f.name())
		}
	}
	return out
}
