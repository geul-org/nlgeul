//ff:func feature=data type=helper control=sequence
//ff:what errBadBlock — the error for a Blocks.txt line whose code points cannot be read
package main

import "fmt"

func errBadBlock(line string) error {
	return fmt.Errorf("Blocks.txt 줄을 읽을 수 없다: %q", line)
}
