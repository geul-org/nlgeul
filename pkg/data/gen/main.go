//ff:func feature=data type=command control=sequence
//ff:what main — makes the special-character table: go run ./pkg/data/gen <UnicodeData.txt> <Blocks.txt> > pkg/data/special-forms.tsv
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "쓰기: go run ./pkg/data/gen <UnicodeData.txt> <Blocks.txt> > pkg/data/special-forms.tsv")
		os.Exit(2)
	}
	if err := run(os.Args[1], os.Args[2], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
