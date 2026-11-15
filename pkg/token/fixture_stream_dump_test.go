//ff:func feature=token type=helper control=iteration dimension=2
//ff:what streamDump — what a stream gives with Next to its end, in lines: for each stretch its region, rank, number of ways and lambda, each node with every candidate in order and its lambda, then the trace lines Notes gives after that Next (prefixed note:); the lines Notes gives after the last Next follow
package token

import "fmt"

func streamDump(st *Stream) []string {
	var out []string
	for {
		sg, ok := st.Next()
		if ok {
			out = append(out, fmt.Sprintf("segment %v %d %v %v", sg.Region, sg.Rank, sg.Ways, sg.Lambda))
		}
		for _, n := range sg.Nodes {
			out = append(out, nodeDump(n)...)
		}
		for _, line := range st.Notes() {
			out = append(out, "note:"+line)
		}
		if !ok {
			return out
		}
	}
}
