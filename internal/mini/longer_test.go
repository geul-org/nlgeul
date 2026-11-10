//ff:func feature=mini type=helper control=sequence
//ff:what TestLonger — applies only when the rival is strictly longer
package mini

import "testing"

func TestLonger(t *testing.T) {
	if longer(2, 4) != 1 || longer(4, 4) != -1 || longer(4, 2) != -1 {
		t.Fatal("longer")
	}
}
