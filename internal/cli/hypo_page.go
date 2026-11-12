//ff:func feature=cli type=command control=sequence
//ff:what hypoPage — the page nlgeul hypo is asked for and its size, from --page and --per-page; page 0 when no page is asked for; a page or a size below 1, and a size without a page, are refused
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func hypoPage(cmd *cobra.Command) (int, int, error) {
	page, _ := cmd.Flags().GetInt("page")
	perPage, _ := cmd.Flags().GetInt("per-page")
	if !cmd.Flags().Changed("page") {
		if cmd.Flags().Changed("per-page") {
			return 0, 0, fmt.Errorf("--per-page는 --page와 함께 준다")
		}
		return 0, perPage, nil
	}
	if page < 1 {
		return 0, 0, fmt.Errorf("--page는 1 이상이어야 한다: %d", page)
	}
	if perPage < 1 {
		return 0, 0, fmt.Errorf("--per-page는 1 이상이어야 한다: %d", perPage)
	}
	return page, perPage, nil
}
