//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what withoutLines — the lines of a dump less the trace lines of the rules named (those with "\t<name>\t" in them)
package parse

import "strings"

func withoutLines(lines []string, names ...string) []string {
	var out []string
	for _, l := range lines {
		keep := true
		for _, name := range names {
			keep = keep && !strings.Contains(l, "\t"+name+"\t")
		}
		if keep {
			out = append(out, l)
		}
	}
	return out
}
