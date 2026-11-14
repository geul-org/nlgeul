//ff:func feature=token type=helper control=sequence
//ff:what TestJSONWays — a small number of ways is a number, none is 0, more than 2^53 is its digits as a string
package token

import (
	"math/big"
	"testing"
)

func TestJSONWays(t *testing.T) {
	if jsonWays(nil) != int64(0) || jsonWays(big.NewInt(60)) != int64(60) {
		t.Fatal("small")
	}
	big1 := new(big.Int).Lsh(big.NewInt(1), 70)
	if jsonWays(big1) != big1.String() {
		t.Fatal(jsonWays(big1))
	}
	if jsonWays(new(big.Int).Lsh(big.NewInt(1), 54)) != "18014398509481984" {
		t.Fatal("2^54")
	}
}
