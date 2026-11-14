//ff:func feature=cli type=helper control=sequence
//ff:what TestSpacingOption — all gives every name of token.SpacingArticles (a copy of it), none and nothing give none, anything else is the names joined with ","
package cli

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/token"
)

func TestSpacingOption(t *testing.T) {
	all := spacingOption("all")
	if !slices.Equal(all, token.SpacingArticles) {
		t.Fatal(all)
	}
	all[0] = "x"
	if token.SpacingArticles[0] == "x" {
		t.Fatal("not a copy")
	}
	if got := spacingOption("none"); got != nil {
		t.Fatal(got)
	}
	if got := spacingOption(""); got != nil {
		t.Fatal(got)
	}
	if got := spacingOption("41,43다만"); !slices.Equal(got, []string{"41", "43다만"}) {
		t.Fatal(got)
	}
}
