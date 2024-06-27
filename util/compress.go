package util

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/Walter0697/zonai/model"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func CompressImageList(image_list []string, parent *model.ProjectParentModel, configuration model.ProjectConfigurationModel) {
	s := spinner.New(spinner.CharSets[21], 500*time.Millisecond)
	s.Suffix = " Compressing images..."
	s.Start()

	now := time.Now().Format("2006-01-02_15_04_05")
	tarResultFile := path.Join(configuration.OutputImagePath, parent.ProjectName+"_"+now+".gz")

	targetFile, err := os.Create(tarResultFile)
	if err != nil {
		panic(err)
	}

	var fileW io.WriteCloser = targetFile
	tarfileW := tar.NewWriter(fileW)
	defer tarfileW.Close()

	for _, image := range image_list {
		imagePath := path.Join(configuration.OutputImagePath, image)
		file, err := os.Open(imagePath)
		if err != nil {
			panic(err)
		}
		fileInfo, _ := file.Stat()

		header := new(tar.Header)
		header.Name = image
		header.Size = fileInfo.Size()
		header.Mode = int64(fileInfo.Mode())
		header.ModTime = fileInfo.ModTime()

		err = tarfileW.WriteHeader(header)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(tarfileW, file)
		if err != nil {
			panic(err)
		}
	}

	s.Stop()
	resultFileDisplay := color.YellowString(tarResultFile)
	fmt.Println("Images compressed successfully into file: " + resultFileDisplay)
}
