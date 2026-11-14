//ff:func feature=token type=helper control=sequence
//ff:what TestJSONFloat — a number is kept; infinity and not a number are nothing (null)
package token

import (
	"math"
	"testing"
)

func TestJSONFloat(t *testing.T) {
	if p := jsonFloat(-1.5); p == nil || *p != -1.5 {
		t.Fatal(p)
	}
	if jsonFloat(math.Inf(-1)) != nil || jsonFloat(math.Inf(1)) != nil || jsonFloat(math.NaN()) != nil {
		t.Fatal("not a number kept")
	}
}
