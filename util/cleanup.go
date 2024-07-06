package util

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

const (
	inputWarning  = "- all images (.tar) in the input folder will be deleted."
	outputWarning = "- all extracted folder (start with .) in the output folder will be deleted."
)

func ShowWarning(display []string) bool {
	for _, item := range display {
		color.Yellow(item)
	}

	prompt := promptui.Prompt{
		Label: "Are you sure you want to continue? [y/n]",
	}

	result, _ := prompt.Run()

	return strings.ToLower(result) == "y"
}

func CleanupAll(autoYes bool) {
	if !autoYes {
		if !ShowWarning([]string{inputWarning, outputWarning}) {
			return
		}
	}

	cleanupOutput()
	cleanupInput()
}

func CleanupOutputFolder(autoYes bool) {
	if !autoYes {
		if !ShowWarning([]string{outputWarning}) {
			return
		}
	}

	cleanupOutput()
}

func CleanupInputFolder(autoYes bool) {
	if !autoYes {
		if !ShowWarning([]string{inputWarning}) {
			return
		}
	}
	cleanupInput()
}

func cleanupOutput() {
	configuration := ReadConfiguration()
	outputFolder := configuration.OutputImagePath

	// confirm that the path exists
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		color.Red("The output folder does not exist.")
		return
	}

	// read all files in the output folder
	files, err := os.ReadDir(outputFolder)
	if err != nil {
		color.Red("Error reading the output folder.")
		return
	}

	s := spinner.New(spinner.CharSets[5], 500*time.Millisecond)
	s.Suffix = " Cleaning up the output folder..."
	s.Start()

	// delete all files in the output folder
	count := 0
	for _, file := range files {
		filename := file.Name()
		if path.Ext(filename) != ".tar" {
			continue
		}

		err := os.Remove(fmt.Sprintf("%s/%s", outputFolder, filename))
		if err != nil {
			color.Red("Error deleting file %s", filename)
			continue
		}

		count++
	}

	s.Stop()

	color.Green("Deleted %d files in the output folder.", count)
}

func cleanupInput() {
	configuration := ReadConfiguration()
	inputFolder := configuration.InputImagePath

	// confirm that the path exists
	if _, err := os.Stat(inputFolder); os.IsNotExist(err) {
		color.Red("The output folder does not exist.")
		return
	}

	// read all files in the output folder
	files, err := os.ReadDir(inputFolder)
	if err != nil {
		color.Red("Error reading the output folder.")
		return
	}

	s := spinner.New(spinner.CharSets[5], 500*time.Millisecond)
	s.Suffix = " Cleaning up the input folder..."
	s.Start()

	// delete all files in the output folder
	count := 0
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		filename := file.Name()
		if !strings.HasPrefix(filename, ".") {
			continue
		}
		if !strings.Contains(filename, "_") {
			continue
		}

		err := os.RemoveAll(fmt.Sprintf("%s/%s", inputFolder, filename))
		if err != nil {
			color.Red("Error deleting file %s", filename)
			continue
		}

		count++
	}

	s.Stop()

	color.Green("Deleted %d files in the input folder.", count)
}
