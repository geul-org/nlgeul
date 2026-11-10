//ff:func feature=cli type=command control=sequence
//ff:what newRootCmd — the nlgeul command with parse and token
package cli

import "github.com/spf13/cobra"

// newRootCmd builds the command tree. Cobra's completion command is
// turned off so that only parse and token are there.
func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "nlgeul",
		Short:         "한국어 문법 음절 표 기반 분석기(실험)",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.CompletionOptions.DisableDefaultCmd = true
	root.AddCommand(newParseCmd(), newTokenCmd())
	return root
}
