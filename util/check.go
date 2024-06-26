package util

import (
	"os"

	"github.com/fatih/color"
)

func CheckSetup() {
	configuration := ReadConfiguration()
	if configuration.OutputImagePath == "" {
		color.Red("OutputImagePath is not set")
		os.Exit(1)
	}
}
