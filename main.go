// Refresh is a command line tool that build and restart web
// application when you change Go and other source files.  Refresh
// will watch for file events like create, modifiy or delete and it
// will build and restart the application.

package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/baijum/refresh/runner"
)

func setConfigPath(configPath string) error {
	if configPath != "" {
		if _, err := os.Stat(configPath); err != nil {
			return errors.New("Cannot find config file: " + configPath)
		}
		os.Setenv("REFRESH_CONFIG_PATH", configPath)
	}
	return nil
}

func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()
	err := setConfigPath(*configPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	runner.Start()
}
