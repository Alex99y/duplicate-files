package core

import (
	"fmt"
	"path"
	"sync"

	"github.com/Alex99y/duplicate-files/pkg/cmd"
	"github.com/Alex99y/duplicate-files/pkg/crypto"
	"github.com/Alex99y/duplicate-files/pkg/structures"
	"github.com/Alex99y/duplicate-files/pkg/utils"
)

var wg sync.WaitGroup

func processFolder(file string) {
	// Retries before return
	isDir, err := utils.IsDirectory(file)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}
	if isDir {
		// Process folder
		files := utils.ReadFilesFromDirectory(file)
		for _, f := range files {
			wg.Add(1)
			go processFolder(path.Join(file, f))
		}
	} else {
		// Process regular file
		fileContent := utils.ReadFile(file)
		fileHash := crypto.GetFileHash(fileContent)
		structures.AddElement(fileHash, file)
		fileContent = nil
	}

	// End task
	wg.Done()
}

// Start function will start the thread process
func Start(config cmd.CobraInterface) {

	// Excecute first thread
	wg.Add(1)

	// Start searching
	go processFolder(config.RootFolder)

	// Wait until goroutines ends
	wg.Wait()

	resultMap := structures.GetMap()
	gotDuplicates := false

	resultMap.Range(func(key interface{}, value interface{}) bool {
		files := value.([]string)
		if len(files) > 1 {
			gotDuplicates = true
			fmt.Println("Duplicated files (" + key.(string) + "):")
			for _, file := range files {
				fmt.Println(file)
			}
			fmt.Print("\n")
		}
		return true
	})

	if gotDuplicates == false {
		fmt.Println("No duplicated files found")
	}
}
