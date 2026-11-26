//ff:func feature=parse type=reader control=sequence
//ff:what politeNamed — whether a name is that of an old polite slot of the index
package parse

import "slices"

func (ix *index) politeNamed(name string) bool {
	return slices.ContainsFunc(ix.polite, func(e *entry) bool { return e.name == name })
}
