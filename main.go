package main

import (
	"date-range-generator/command"
	"os"
)

func main() {
	if err := command.GetGenerationCommand().Execute(); err != nil {
		os.Exit(-1)
	}
}
