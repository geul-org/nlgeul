//ff:func feature=cli type=helper control=selection
//ff:what spacingOption — the spacing rules --spacing asks for: all of them for "all", none for "none" (or nothing), else the names given joined with ","
package cli

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/parse"
)

func spacingOption(v string) []string {
	switch v {
	case "all":
		return append([]string{}, parse.SpacingArticles...)
	case "none", "":
		return nil
	}
	return strings.Split(v, ",")
}
