//ff:type feature=data type=model
//ff:what charRange — a range of code points that UnicodeData.txt writes as a First/Last pair, with the one category the whole range has
package main

type charRange struct {
	Lo       rune
	Hi       rune
	Category string
}
