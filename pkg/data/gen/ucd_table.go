//ff:type feature=data type=model
//ff:what ucdTable — the characters of UnicodeData.txt by code point, and the ranges written as First/Last pairs
package main

type ucdTable struct {
	Chars  map[rune]charInfo
	Ranges []charRange
}
