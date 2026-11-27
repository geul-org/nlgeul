//ff:func feature=data type=reader control=sequence
//ff:what loadBlocks — reads the Blocks.txt file at a path into its blocks and version
package main

import "os"

func loadBlocks(path string) ([]blockRange, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return parseBlocks(f)
}
