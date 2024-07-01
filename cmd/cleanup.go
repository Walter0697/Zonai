/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup [input | output] (-a | --all)",
	Short: "Clean up all the unzipped files, uncompressed images in input and output folder",
	Long: `By using this command, you can clean up all unnecessary files in input or output folder
	For input folder, you will remove all the folders with docker images file(.tar) in it, but it won't remove .gz file so that you can still revert or deploy the project.
	For output folder, you will remove all docker images file(.tar) in it since .gz is the one you should move and deploy.
	
	Therefore, cleaning up folders will not affect anything about deployment, it just saves spaces on your disk.`,
	Example: `
	zonai cleanup input
	zonai cleanup -a`,
	ValidArgs: []string{"input", "output", "-a", "--all"},
	Run: func(cmd *cobra.Command, args []string) {
		allFlags, _ := cmd.Flags().GetBool("all")

		if allFlags {
			util.CleanupAll()
			return
		}

		if len(args) == 0 {
			color.Red("--> Please provide a type: [input | output]")
			os.Exit(1)
		}

		if args[0] != "input" && args[0] != "output" {
			color.Red("--> Please provide a valid type: [input | output]")
			os.Exit(1)
		}

		if args[0] == "input" {
			util.CleanupInputFolder()
		} else {
			util.CleanupOutputFolder()
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
	cleanupCmd.Flags().BoolP("all", "a", false, "Clean up all the unzipped files, uncompressed images in input and output folder")
}
