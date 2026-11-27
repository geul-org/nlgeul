//ff:func feature=data type=helper control=sequence
//ff:what addUCDLine — puts one UnicodeData.txt line in the table; a First line is remembered and the Last line after it makes a range
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addUCDLine(t *ucdTable, line string, first *rune) error {
	f := strings.Split(line, ";")
	if len(f) < 9 {
		return fmt.Errorf("UnicodeData.txt 줄의 칸이 모자라다: %q", line)
	}
	code, err := strconv.ParseInt(f[0], 16, 32)
	if err != nil {
		return fmt.Errorf("UnicodeData.txt 코드가 아니다: %q", f[0])
	}
	name := f[1]
	if strings.HasSuffix(name, ", First>") {
		*first = rune(code)
		return nil
	}
	if strings.HasSuffix(name, ", Last>") {
		t.Ranges = append(t.Ranges, charRange{Lo: *first, Hi: rune(code), Category: f[2]})
		return nil
	}
	if strings.HasPrefix(name, "<") {
		name = ""
	}
	t.Chars[rune(code)] = charInfo{Name: name, Category: f[2], Decomp: f[5], Numeric: f[8]}
	return nil
}
