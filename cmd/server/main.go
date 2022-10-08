package main

import (
	"os"

	"artbycode.id/go-app/cmd/server/runner"
)

func main() {
	server := runner.InitializeRunnerServer()
	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
