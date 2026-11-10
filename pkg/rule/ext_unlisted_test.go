//ff:func feature=rule type=rule control=sequence
//ff:what TestExtUnlisted — applies when the run right of the dot is not in the list; letters are compared in lower case
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestExtUnlisted(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("ext", []string{"py", "md"})}
	in, _ := ExtUnlisted(markCtx(Mark{Ext: "MD"}), specs)
	out, _ := ExtUnlisted(markCtx(Mark{Ext: "exe"}), specs)
	if in != -1 || out != 1 {
		t.Fatal(in, out)
	}
}
