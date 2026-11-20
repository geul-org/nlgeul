//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDirPath — end to end with no word table: a name, a slash, a name and the slash that ends it is one path (pkg/parser/[경로]) with particles after it; a name with only its end slash is no path
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamDirPath(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ sentence, first string }{
		{"pkg/parser/ 를 본다", "pkg/parser/[경로]"},
		{"plans/security/에 둔다", "plans/security/[경로]+에[부사격:에]"},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(r, c.sentence), "\n") + "\n"
		if !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
	}
	if got := streamText(r, "files/에 둔다"); strings.Contains(got, "[경로]") {
		t.Errorf("files/에 is read as a path:\n%s", got)
	}
}
