//ff:func feature=data type=helper control=iteration dimension=1
//ff:what allDigits — whether a string is not empty and every character of it is a decimal digit
package main

import "unicode"

func allDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
