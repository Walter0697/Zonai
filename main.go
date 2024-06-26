/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
