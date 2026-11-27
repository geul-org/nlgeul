//ff:type feature=data type=model
//ff:what blockRange — one block of Blocks.txt: first and last code point and the block name
package main

type blockRange struct {
	Lo   rune
	Hi   rune
	Name string
}
