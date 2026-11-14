//ff:func feature=token type=reader control=sequence
//ff:what TestNetStep — the function is worked through by the net over the sentences of the tests and made-up long units, and each is read to at least one stretch
package token

import "testing"

func TestNetStep(t *testing.T) {
	driveAll(t)
}
