//ff:func feature=token type=helper control=sequence
//ff:what partsBound — the words a stem is made of when it is words bound by one kind of sign, a slash or a hyphen, with a letter on both sides of each sign; nil when it holds no such sign, both kinds, or an empty word
package token

import (
	"slices"
	"strings"
)

func partsBound(stem string) []string {
	slash, hyphen := strings.Contains(stem, "/"), strings.Contains(stem, "-")
	if slash == hyphen {
		return nil
	}
	sign := "-"
	if slash {
		sign = "/"
	}
	parts := strings.Split(stem, sign)
	if slices.Contains(parts, "") {
		return nil
	}
	return parts
}
