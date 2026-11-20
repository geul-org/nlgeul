//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what pullUnit — read the unit a stream is at to its end: every stretch of it queued
package parse

func pullUnit(st *Stream) {
	for target := st.read + 1; st.read < target; {
		st.pull()
	}
}
