//ff:func feature=data type=reader control=iteration dimension=1
//ff:what parseUnicodeData — reads UnicodeData.txt into a table, line by line
package main

import (
	"bufio"
	"io"
)

func parseUnicodeData(r io.Reader) (*ucdTable, error) {
	t := &ucdTable{Chars: map[rune]charInfo{}}
	var first rune
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		if err := addUCDLine(t, line, &first); err != nil {
			return nil, err
		}
	}
	return t, sc.Err()
}
