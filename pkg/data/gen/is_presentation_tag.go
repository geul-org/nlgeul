//ff:func feature=data type=helper control=sequence
//ff:what isPresentationTag — whether a decomposition tag is one of the Arabic presentation forms
package main

import "slices"

func isPresentationTag(tag string) bool {
	return slices.Contains(presentationTags, tag)
}
