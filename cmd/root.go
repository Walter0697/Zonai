/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Zonai",
	Short: "Zonai helps you dockerize your applications with ease.",
	Long: `Zonai is a tool that helps you dockerize your applications, from saving a project to loading it into internal server
	The main idea is to serve system that without internet and heavily relies on internal server, so that they cannot use CI/CD.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
