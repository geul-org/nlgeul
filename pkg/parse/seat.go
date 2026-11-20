//ff:type feature=parse type=model
//ff:what seat — one word of the line the rules of a sentence judge: which of the regions not yet judged it is in, its node, and how its first candidate reads it as it stands now
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

type seat struct {
	region int
	node   Node
	near   rule.Near
}
