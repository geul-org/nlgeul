//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what complementKept — checks that a sentence read with complementRule has the same words, each with as many candidates under the same names, as the stream with no rule
package parse

import "testing"

func complementKept(t *testing.T, s string) {
	t.Helper()
	plain, _ := complementNodes(complementStream(t, s, false, false))
	ruled, _ := complementNodes(complementStream(t, s, false, true))
	if len(plain) != len(ruled) {
		t.Fatal(s, len(plain), len(ruled))
	}
	for i := range plain {
		complementSame(t, s, plain[i], ruled[i])
	}
}
