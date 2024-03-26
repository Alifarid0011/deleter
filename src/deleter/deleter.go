package deleter

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func CheckDirectoryPeriodically(dirPath string, duration time.Duration, formatsToDelete []string, CheckTime time.Duration) {
	// Repeatedly check the directory at the specified interval
	for range time.Tick(CheckTime) {
		log.Printf("Locking directory: %s", dirPath)
		// Open the directory
		dir, err := os.Open(dirPath)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		// Read the directory entries
		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		// Close the directory after reading
		err = dir.Close()
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		// Iterate over the files in the directory
		for _, fileInfo := range fileInfos {
			// Check if the entry is a regular file
			if fileInfo.Mode().IsRegular() {
				// Check if the file format is in the formats to delete list
				if isInFormatsToDeleteList(filepath.Ext(fileInfo.Name()), formatsToDelete) {
					// Calculate the time elapsed since the file was created
					elapsed := time.Since(fileInfo.ModTime())
					// Check if more than the specified interval has elapsed since the file creation
					if elapsed > duration {
						// Construct the file path
						filePath := filepath.Join(dirPath, fileInfo.Name())
						// Delete the file
						err := os.Remove(filePath)
						if err != nil {
							log.Println("Error:", err)
							continue
						}
						log.Printf("Deleted file: %s", filePath)
					}
				}
			}
		}
	}
}

// Function to check if a format exists in the list of formats to delete
func isInFormatsToDeleteList(format string, formats []string) bool {
	for _, f := range formats {
		// Check if the format is in the list
		if format == "."+f || format == f {
			return true
		}
	}
	return false
}
