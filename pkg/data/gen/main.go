//ff:func feature=data type=command control=sequence
//ff:what main — makes a table: go run ./pkg/data/gen forms|groups <UnicodeData.txt> <Blocks.txt> > the table file
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "쓰기: go run ./pkg/data/gen forms|groups <UnicodeData.txt> <Blocks.txt> > 표 파일")
		os.Exit(2)
	}
	if err := run(os.Args[1], os.Args[2], os.Args[3], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
