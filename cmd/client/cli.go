package main

import (
	"log"
	"os"

	"grpc-gateway/cmd/client/internal"
)

func main() {
	var rootCmd = internal.CobraCommands()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)

	}
}
