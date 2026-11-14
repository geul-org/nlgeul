//ff:func feature=token type=helper control=iteration dimension=1
//ff:what errHas — fail a test unless there is an error and its text holds every part given
package token

import (
	"strings"
	"testing"
)

func errHas(t *testing.T, err error, parts ...string) {
	t.Helper()
	if err == nil {
		t.Fatal("no error")
	}
	for _, p := range parts {
		if !strings.Contains(err.Error(), p) {
			t.Errorf("%q lacks %q", err, p)
		}
	}
}
