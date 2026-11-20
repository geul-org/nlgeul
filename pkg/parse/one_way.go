//ff:func feature=parse type=helper control=sequence
//ff:what oneWay — is a stretch cut one way at most (Ways is 1 or less)
package parse

import "math/big"

func oneWay(seg Segment) bool {
	return seg.Ways == nil || seg.Ways.Cmp(big.NewInt(1)) <= 0
}
