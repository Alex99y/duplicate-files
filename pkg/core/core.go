package core

import (
	"fmt"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/Alex99y/duplicate-files/pkg/cmd"
	"github.com/Alex99y/duplicate-files/pkg/crypto"
	"github.com/Alex99y/duplicate-files/pkg/structures"
	"github.com/Alex99y/duplicate-files/pkg/utils"
)

const threadRetryBeforeReturn = 1

var wg sync.WaitGroup

// StructureInfo contains the configuration to start the process
type StructureInfo struct {
	folderQueue *structures.QueueWithSync
	resultMap   *structures.MapWithSync
}

func (s *StructureInfo) processFile() {
	// Process file
}

func (s *StructureInfo) processFolder(id int) {
	// Retries before return
	retriesLeft := threadRetryBeforeReturn
	for {
		// Dequeue next file to process
		nextFolderToProcess := s.folderQueue.Dequeue()

		if nextFolderToProcess != nil {
			retriesLeft = threadRetryBeforeReturn
			file := nextFolderToProcess.(string)
			isDir, err := utils.IsDirectory(file)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if isDir {
				// Process folder
				files := utils.ReadFilesFromDirectory(file)
				for _, f := range files {
					s.folderQueue.Enqueue(path.Join(file, f))
				}
			} else {
				// Process regular file
				fileContent := utils.ReadFile(file)
				fileHash := crypto.GetFileHash(fileContent)
				s.resultMap.AddElement(fileHash, file)
				fileContent = nil
			}
		} else {
			if retriesLeft == 0 {
				break
			} else {
				time.Sleep(500 * time.Millisecond)
				retriesLeft--
			}
		}
	}

	// End task
	wg.Done()
}

// Start function will start the thread process
func Start(config cmd.CobraInterface) {

	// Prepare queues
	structure := StructureInfo{
		// Contains the folder/files to process
		folderQueue: structures.NewQueue(),
		// Contains the result (duplicated files)
		resultMap: structures.NewMap(),
	}
	structure.folderQueue.Enqueue(config.RootFolder)

	// Total threads to improve paralellism
	threads := runtime.NumCPU() / 2
	if int(config.NumberOfThreads) < threads {
		threads = int(config.NumberOfThreads)
	}
	runtime.GOMAXPROCS(threads)
	wg.Add(threads)

	// Start searching
	for i := 0; i < threads; i++ {
		go structure.processFolder(i)
	}

	// Wait until goroutines ends
	wg.Wait()

	resultMap := structure.resultMap.GetMap()
	gotDuplicates := false

	for key, files := range resultMap {
		if len(files) > 1 {
			gotDuplicates = true
			fmt.Println("Duplicated files (" + key + "):")
			for _, file := range files {
				fmt.Println(file)
			}
			fmt.Print("\n")
		}
	}
	if gotDuplicates == false {
		fmt.Println("No duplicated files found")
	}
}
