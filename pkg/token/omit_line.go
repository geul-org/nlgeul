//ff:func feature=token type=helper control=sequence
//ff:what omitLine — the line after the last box of a contest of ways that left some ways out: "(n개 생략됨)" starting where the brackets start (one space after the boxes, width + 4 cells wide)
package token

import (
	"fmt"
	"math/big"
	"strings"
)

func omitLine(width int, left *big.Int) string {
	return strings.Repeat(" ", width+4+1) + fmt.Sprintf("(%d개 생략됨)", left)
}
