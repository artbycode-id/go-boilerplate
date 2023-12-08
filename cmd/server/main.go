package main

import (
	"log"

	"artbycode.id/go-app/cmd/server/runner"
)

func main() {
	runnerServer := runner.InitializeRunnerServer()
	if err := runnerServer.Run(); err != nil {
		log.Fatal(err)
	}
}
