package prompt

import (
	"fmt"

	"github.com/fatih/color"
)

func ExecuteCli() {
	commandLine := color.YellowString("zonai -h")
	fmt.Println("To use command line tool, please type: " + commandLine + " for more information.")
	Execute()
}
