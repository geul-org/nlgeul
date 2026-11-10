//ff:func feature=cli type=command control=sequence
//ff:what main — run kopas and exit with its code (0, or 1 after an error)
package main

import (
	"os"

	"github.com/geul-org/koparse/internal/cli"
)

// exit ends the process; tests replace it to observe the exit code.
var exit = os.Exit

func main() {
	exit(cli.Execute())
}
