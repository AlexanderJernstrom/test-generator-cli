package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type FileInfos struct {
	Path    string
	Content string
}

func WalkDir() ([]FileInfos, error) {

	var fileInfos []FileInfos = []FileInfos{}

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			content, err := ioutil.ReadFile(path)

			if err != nil {
				log.Fatal(err)
				return err
			}

			fileInfos = append(fileInfos, FileInfos{Path: path, Content: string(content)})
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return fileInfos, nil
}
