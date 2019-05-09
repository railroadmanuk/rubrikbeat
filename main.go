package main

import (
	"os"

	"github.com/railroadmanuk/rubrikbeat/cmd"

	_ "github.com/railroadmanuk/rubrikbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
