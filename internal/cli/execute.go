//ff:func feature=cli type=command control=sequence
//ff:what Execute — run nlgeul with the process's arguments; the exit code
package cli

import (
	"fmt"
	"os"
)

// Execute returns 0, or 1 after printing the error to standard error.
func Execute() int {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "nlgeul:", err)
		return 1
	}
	return 0
}
