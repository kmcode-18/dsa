package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func main() {
	rootDir := "./"
	files, err := GetFiles(rootDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	sem := make(chan struct{}, 40)        // Semaphore to limit concurrency
	outChan := make(chan int, len(files)) // Buffered channel based on number of files
	var wg sync.WaitGroup

	for _, v := range files {
		wg.Add(1)
		sem <- struct{}{} // Send to semaphore
		go func(v string) {
			defer wg.Done()
			defer func() { <-sem }() // Release from semaphore
			ReadLines(v, outChan)
		}(v)
	}

	go func() {
		wg.Wait()
		close(outChan) // Close the channel after all goroutines are done
	}()

	out := 0
	for v := range outChan {
		fmt.Println("number of goroutines:", runtime.NumGoroutine())
		out += v
	}
	fmt.Println("Total lines:", out)
}

func ReadLines(file string, outChan chan int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		outChan <- 0 // Send a default value to prevent deadlock
		return
	}
	defer f.Close() // Ensure the file is closed after reading

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count++
	}

	// Send the line count to the channel
	select {
	case outChan <- count:
	default:
		fmt.Println("outChan is blocked, unable to send")
	}
}

func GetFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

