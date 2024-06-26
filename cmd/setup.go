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

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup [flags] [setup value]",
	Short: "Setup the configuration for Zonai",
	Long: `We don't need much for Zonai setup, just an Output Image Path for storing your image.
	If you have any specific needs (like a M-series Mac), you might need other Docker build command`,
	Example: `
	zonai setup -o /path/to/image
	zonai setup --output-image-path=/path/to/image
	zonai setup -d "docker build -t"
	zonai setup --docker-build-command="docker build -t"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		outputImagePath := cmd.Flags().Lookup("output-image-path").Value
		dockerBuildCommand := cmd.Flags().Lookup("docker-build-command").Value
		if outputImagePath.String() == "" && dockerBuildCommand.String() == "" {
			color.Red("--> Please provide at least one flag: output-image-path or docker-build-command")
			os.Exit(1)
		}

		configuration := util.ReadConfiguration()
		if outputImagePath.String() != "" {
			configuration.OutputImagePath = outputImagePath.String()
			color.Cyan("Set up OutputImagePath...")
		}

		if dockerBuildCommand.String() != "" {
			configuration.DockerBuildCommand = dockerBuildCommand.String()
			color.Cyan("Set up DockerBuildCommand...")
		}

		util.SaveConfiguration(configuration)
		color.Cyan("Done!")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.PersistentFlags().StringP("output-image-path", "o", "", "Output Image Path")
	rootCmd.PersistentFlags().StringP("docker-build-command", "d", "", "Docker Build Command")
}