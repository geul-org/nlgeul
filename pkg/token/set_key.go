//ff:func feature=token type=helper control=iteration dimension=1
//ff:what setKey — a span and a set of states as one key: the span's place and the states' places, each as a variable-length number
package token

import "encoding/binary"

func setKey(node int, set []int) string {
	b := binary.AppendUvarint(nil, uint64(node))
	for _, id := range set {
		b = binary.AppendUvarint(b, uint64(id))
	}
	return string(b)
}
