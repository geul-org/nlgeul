//ff:func feature=data type=reader control=iteration dimension=1
//ff:what parseBlocks — reads Blocks.txt into its blocks and the version written on its first line
package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func parseBlocks(r io.Reader) ([]blockRange, string, error) {
	var out []blockRange
	version := ""
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if m := versionLine.FindStringSubmatch(line); m != nil {
			version = m[1]
		}
		m := blockLine.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		lo, err1 := strconv.ParseInt(m[1], 16, 32)
		hi, err2 := strconv.ParseInt(m[2], 16, 32)
		if err1 != nil || err2 != nil {
			return nil, "", errBadBlock(line)
		}
		out = append(out, blockRange{Lo: rune(lo), Hi: rune(hi), Name: strings.TrimSpace(m[3])})
	}
	return out, version, sc.Err()
}
