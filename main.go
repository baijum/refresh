/* Refresh is a command line tool that builds and (re)starts your web
application everytime you save a go or template file.

If the web framework you are using supports the Refresh runner, it
will show build errors on your browser.

Refresh will watch for file events, and every time you
create/modifiy/delete a file it will build and restart the
application.  If `go build` returns an error, it will logs it in the
tmp folder.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/baijum/refresh/runner"
)

func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()

	if *configPath != "" {
		if _, err := os.Stat(*configPath); err != nil {
			fmt.Printf("Can't find config file `%s`\n", *configPath)
			os.Exit(1)
		} else {
			os.Setenv("RUNNER_CONFIG_PATH", *configPath)
		}
	}

	runner.Start()
}
