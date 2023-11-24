package main

import (
	"log"

	cmd "github.com/dreamer-zq/turbo-tester/cmd/tester"
)
func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
