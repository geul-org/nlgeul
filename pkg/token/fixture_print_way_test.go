//ff:func feature=token type=helper control=iteration dimension=2
//ff:what printWay — a way as the example prints it: the stretch, the rank of the way and how many ways there are, then each word with its candidates as given (the letters, the verdict, the preference and the lambda of each)
package token_test

import (
	"fmt"

	"github.com/geul-org/nlgeul/pkg/token"
)

func printWay(seg token.Segment) {
	fmt.Printf("[%d-%d] way %d of %s\n", seg.Region.Start, seg.Region.End, seg.Rank, seg.Ways)
	for _, n := range seg.Nodes {
		fmt.Printf("  %s:", n.Contest.Text)
		for _, br := range n.Contest.Branches {
			fmt.Printf(" %s (%.2f %.2f %.2f)", br, br.Verdict, br.Pref, br.Lambda())
		}
		fmt.Println()
	}
}
