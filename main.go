/*
Copyright © 2024 Walter Cheng <waltercheng621@gmail.com>
*/
package main

import (
	"github.com/Walter0697/zonai/cmd"
	"github.com/Walter0697/zonai/util"
)

func main() {
	util.InitializeFolder()
	cmd.Execute()
}
