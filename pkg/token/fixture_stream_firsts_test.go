//ff:func feature=token type=helper control=iteration dimension=1
//ff:what streamFirsts — what Next gives until the end of the stream, in order
package token

func streamFirsts(st *Stream) []Segment {
	var out []Segment
	for {
		sg, ok := st.Next()
		if !ok {
			return out
		}
		out = append(out, sg)
	}
}
