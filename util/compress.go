package util

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Walter0697/zonai/model"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func addFileToTarWriter(filePath string, tarWriter *tar.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not open file '%s', got error '%s'", filePath, err.Error()))
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return errors.New(fmt.Sprintf("Could not get stat for file '%s', got error '%s'", filePath, err.Error()))
	}

	filePathInfo := strings.Split(filePath, "/")
	// last element is the file name
	filename := filePathInfo[len(filePathInfo)-1]
	header := &tar.Header{
		Name:    filename,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not write header for file '%s', got error '%s'", filePath, err.Error()))
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not copy the file '%s' data to the tarball, got error '%s'", filePath, err.Error()))
	}

	return nil
}

func CompressImageList(image_list []string, parent *model.ProjectParentModel, configuration model.ProjectConfigurationModel) {
	s := spinner.New(spinner.CharSets[21], 500*time.Millisecond)
	s.Suffix = " Compressing images..."
	s.Start()

	now := time.Now().Format("2006-01-02_15_04_05")
	tarResultFile := path.Join(configuration.OutputImagePath, parent.ProjectName+"_"+now+".tar")
	file, err := os.Create(tarResultFile)
	if err != nil {
		s.Stop()
		panic(err)
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, image := range image_list {
		filePath := path.Join(configuration.OutputImagePath, image)
		err := addFileToTarWriter(filePath, tarWriter)
		if err != nil {
			s.Stop()
			panic(err)
		}
	}

	s.Stop()

	resultFileDisplay := color.YellowString(tarResultFile)
	fmt.Println("Images compressed successfully into file: " + resultFileDisplay)
}
