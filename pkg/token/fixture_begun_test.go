//ff:func feature=token type=helper control=sequence
//ff:what begun — how many units a stream has begun to read: those read to their end and the one being read
package token

func begun(st *Stream) int {
	if st.cur != nil {
		return st.read + 1
	}
	return st.read
}
