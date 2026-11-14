//ff:func feature=token type=helper control=sequence
//ff:what TestOneWay — a stretch with 1 way or none (or no number) is cut one way at most; one with 2 or more is not
package token

import (
	"math/big"
	"testing"
)

func TestOneWay(t *testing.T) {
	if !oneWay(Segment{}) || !oneWay(Segment{Ways: big.NewInt(0)}) || !oneWay(Segment{Ways: big.NewInt(1)}) {
		t.Fatal("one way at most")
	}
	if oneWay(Segment{Ways: big.NewInt(2)}) || oneWay(Segment{Ways: new(big.Int).Lsh(big.NewInt(1), 80)}) {
		t.Fatal("more")
	}
}
