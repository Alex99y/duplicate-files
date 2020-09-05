package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

// Max file size is 4gb
const maxFileSize = 4294967296

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// IsDirectory check if the file is or not a directory
func IsDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if fi.Size() > maxFileSize {
		return false, errors.New("Ignore large files: " + fi.Name() + "\n")
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	default:
		return false, nil
	}
}

// ReadFile returns the content of the file
func ReadFile(file string) []byte {
	content, err := ioutil.ReadFile(file)
	check(err)

	return content
}

// ReadFilesFromDirectory returns all file inside a directory
func ReadFilesFromDirectory(dir string) []string {
	var files []string
	f, err := os.Open(dir)
	check(err)
	fileInfo, err := f.Readdir(-1)
	defer f.Close()
	check(err)
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files
}
