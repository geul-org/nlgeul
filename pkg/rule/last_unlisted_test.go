//ff:func feature=rule type=rule control=sequence
//ff:what TestLastUnlisted — applies when the name's last run is not in the list
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestLastUnlisted(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("tld", []string{"com", "kr"})}
	in, _ := LastUnlisted(markCtx(Mark{Ext: "co", Last: "kr"}), specs)
	out, _ := LastUnlisted(markCtx(Mark{Ext: "py", Last: "py"}), specs)
	if in != -1 || out != 1 {
		t.Fatal(in, out)
	}
}
