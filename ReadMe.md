# deleter

## Deleting Files with Specific Formats in a Directory

This function periodically checks a directory and deletes files with specific formats if they have been present for longer than a specified interval.

## Function Signature

```go
func CheckDirectoryPeriodically(dirPath string, duration time.Duration, formatsToDelete []string,CheckTime time.Duration)

```
Example
Suppose we want to delete all ".txt" and ".log" files in the directory "/var/logs" if they have been present for longer than 1 hour.
```go
package main

import (
	"time"

	"github.com/Alifarid0011/deleter/src/deleter"
)

func main() {
	dirPath := "logs"
	duration := 1 * time.Hour
	formatsToDelete := []string{"txt", "log"}
	go deleter.CheckDirectoryPeriodically(dirPath, duration, formatsToDelete, time.Hour)
	select {}
}

```