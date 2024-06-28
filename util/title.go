package util

import (
	"fmt"

	"github.com/fatih/color"
)

const Version = "v0.1.1"

func DrawTitle() {

	bike2 := color.YellowString(".-.=\\-")
	bike3 := color.YellowString("(_)=='(_)")
	bike4 := color.YellowString("(")
	bike5 := color.YellowString(")")
	bike6 := color.YellowString("\\_")
	bike7 := color.YellowString(")")

	end2 := color.CyanString("░░")
	end3 := color.CyanString("▒")
	end4 := color.CyanString("▓▓▓▓▓")
	end5 := color.CyanString("████")
	end6 := color.CyanString("██")
	end7 := color.CyanString("█")
	end8 := color.YellowString("٩(╹ꇴ ╹๑)۶")

	title := color.CyanString("Zonai")
	titleSeparate := color.YellowString(" ================================ ")

	color.Cyan(`░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░`)
	color.Cyan(`░        ░░░      ░░░   ░░░  ░░░      ░░░        ░░░` + bike2 + end2)
	color.Cyan(`▒▒▒▒▒▒  ▒▒▒  ▒▒▒▒  ▒▒    ▒▒  ▒▒  ▒▒▒▒  ▒▒▒▒▒  ▒▒▒▒` + bike3 + end3)
	color.Cyan(`▓▓▓▓  ▓▓▓▓▓  ▓▓▓▓  ▓▓  ▓  ▓  ▓▓  ▓▓▓▓  ▓▓▓▓▓  ▓▓▓▓▓▓▓▓` + bike4 + end4)
	color.Cyan(`██  ███████  ████  ██  ██    ██        █████  █████████` + bike5 + end5)
	color.Cyan(`█        ███      ███  ███   ██  ████  ██        ███████` + bike6 + end6)
	color.Cyan(`██████████████████████████████████████████████████████████` + bike7 + end7)
	color.Green(`Welcome to ` + title + titleSeparate + end8)
}

func ShowVersion() {
	fmt.Println("Zonai version " + Version)
}

func Divider() {
	fmt.Println("=============================================")
}
