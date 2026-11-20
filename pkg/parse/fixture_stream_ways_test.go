//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what streamWays — all the ways of a stretch a stream gives, the first and then Again until there is none
package parse

func streamWays(st *Stream, first Segment) []Segment {
	out := []Segment{first}
	for {
		next, ok := st.Again(first.Region)
		if !ok {
			return out
		}
		out = append(out, next)
	}
}
