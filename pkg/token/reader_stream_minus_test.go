//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamMinus — end to end with no word table: two amounts joined by a hyphen are read each as a number and its unit, with the hyphen alone between them (3,000[수]+원[단위:원], -, 1,000[수]+원[단위:원]); a range of two numbers stays one number
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamMinus(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	first := "\n" + strings.Join(bestLines(r, "차액 (3,000원-1,000원)을 낸다"), "\n") + "\n"
	for _, want := range []string{"3,000[수]+원[단위:원]", "-", "1,000[수]+원[단위:원]"} {
		if !strings.Contains(first, "\n"+want+"\n") {
			t.Errorf("%q is not first in\n%s", want, first)
		}
	}
	if got := strings.Join(bestLines(r, "10-20명"), "\n"); strings.Contains(got, "\n-\n") {
		t.Errorf("10-20명 is cut:\n%s", got)
	}
}
