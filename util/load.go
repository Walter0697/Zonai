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
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/codeclysm/extract"
	"github.com/fatih/color"
)

const (
	ImageContentDir = "/.content"
)

func ExtractCompression(filename string, destination string) string {
	os.MkdirAll(destination, 0755)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(data)
	err = extract.Tar(context.Background(), buffer, destination, nil)
	if err != nil {
		panic(err)
	}
	return destination
}

type manifestData struct {
	RepoTags []string
}

func ReadDockerTag(imagePath string, parentPath string) string {
	imageContentPath := path.Join(parentPath, ImageContentDir)
	extractedPath := ExtractCompression(imagePath, imageContentPath)
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

func LoadAllImagesFromGz(gzFile string, parentPath string) []string {
	s := spinner.New(spinner.CharSets[21], 500*time.Millisecond)
	s.Suffix = " Extracting gz files..."
	s.Start()

	gzFileInfo := strings.Split(gzFile, ".gz")[0]
	gzFullPath := strings.Split(gzFileInfo, "/")
	gzFileName := gzFullPath[len(gzFullPath)-1]
	dockerImagesDir := "." + gzFileName

	dockerImagePath := path.Join(parentPath, dockerImagesDir)
	imageListFolder := ExtractCompression(gzFile, dockerImagePath)
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
		imageTag := ReadDockerTag(imageFile, parentPath)

		imagePathList = append(imagePathList, imageFile)
		imageTagList = append(imageTagList, imageTag)
	}

	lenDisplay := color.YellowString("%d", len(imageTagList))
	fmt.Println("--> Recieved " + lenDisplay + " image(s) in total")
	s2.Stop()

	for _, image := range imagePathList {
		imageInfo := strings.Split(image, "/")
		imageName := imageInfo[len(imageInfo)-1]

		sLoad := spinner.New(spinner.CharSets[43], 500*time.Millisecond)
		sLoad.Suffix = " Loading " + imageName + "..."
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

		imageNameDisplay := color.YellowString(imageName)
		fmt.Println("Image loaded successfully: " + imageNameDisplay)
	}

	return imageTagList
}
