package util

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/briandowns/spinner"
	"github.com/codeclysm/extract"
	"github.com/fatih/color"
)

const (
	DockerImagesDir = "/.images"
	ImageContentDir = "/.content"
)

func ExtractCompression(filename string, dest string) string {
	pwd, _ := os.Getwd()
	destination := path.Join(pwd, dest)
	os.MkdirAll(destination, 0755)
	data, _ := os.ReadFile(filename)
	buffer := bytes.NewBuffer(data)
	err := extract.Tar(context.Background(), buffer, destination, nil)
	if err != nil {
		panic(err)
	}
	return destination
}

type manifestData struct {
	RepoTags []string
}

func ReadDockerTag(imagePath string) string {
	extractedPath := ExtractCompression(imagePath, ImageContentDir)
	manifestFile, err := os.Open(path.Join(extractedPath, "manifest.json"))
	if err != nil {
		panic(err)
	}

	defer manifestFile.Close()

	manifestByte, _ := io.ReadAll(manifestFile)

	var manifest []manifestData
	err = json.Unmarshal(manifestByte, &manifest)
	if err != nil {
		panic(err)
	}

	if len(manifest) == 0 {
		panic("No manifest found")
	}

	if len(manifest[0].RepoTags) == 0 {
		panic("No tag found")
	}

	os.RemoveAll(extractedPath)

	return manifest[0].RepoTags[0]
}

func LoadAllImagesFromGz(gzPath string) []string {
	s := spinner.New(spinner.CharSets[21], 500*time.Millisecond)
	s.Suffix = " Extracting gz files..."
	s.Start()
	imageListFolder := ExtractCompression(gzPath, DockerImagesDir)
	s.Stop()

	// check all files in imagelist folder
	s2 := spinner.New(spinner.CharSets[10], 500*time.Millisecond)
	s2.Suffix = " Analysising images..."
	s2.Start()
	imageListDir, _ := os.Open(imageListFolder)
	imageListFiles, _ := imageListDir.Readdir(0)

	imageTagList := []string{}
	imagePathList := []string{}
	for _, fileInfo := range imageListFiles {
		imageFile := path.Join(imageListFolder, fileInfo.Name())
		imageTag := ReadDockerTag(imageFile)

		imagePathList = append(imagePathList, imageFile)
		imageTagList = append(imageTagList, imageTag)
	}

	lenDisplay := color.YellowString("%d", len(imageTagList))
	fmt.Println("--> Recieved " + lenDisplay + " image(s) in total")
	s2.Stop()

	for _, image := range imagePathList {
		sLoad := spinner.New(spinner.CharSets[43], 500*time.Millisecond)
		sLoad.Suffix = " Loading " + image + "..."
		sLoad.Start()
		dockerCmd := exec.Command("docker", "load", "-i", image)
		stdout, _ := dockerCmd.StdoutPipe()
		dockerCmd.Start()

		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}

		dockerCmd.Wait()
		sLoad.Stop()

		imageNameDisplay := color.YellowString(image)
		fmt.Println("Image loaded successfully: " + imageNameDisplay)
	}

	return imageTagList
}
