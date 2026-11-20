//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what placesOf — for each place inside the units of a sentence (the letters of the units one after another, a space inside a span of code counted) the place with the spaces left out, one more than the letters; nil when no unit holds a space, then the two are the same
package parse

import (
	"strings"
	"unicode"
)

func placesOf(units []unit) []int {
	var text strings.Builder
	for _, u := range units {
		text.WriteString(u.text)
	}
	var places []int
	outer, spaced := 0, false
	for _, r := range text.String() {
		places = append(places, outer)
		if unicode.IsSpace(r) {
			spaced = true
			continue
		}
		outer++
	}
	if !spaced {
		return nil
	}
	return append(places, outer)
}
