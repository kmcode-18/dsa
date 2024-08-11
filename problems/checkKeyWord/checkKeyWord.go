package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Question: Write a program that concurrently searches for a keyword in multiple files and reports the results using goroutines and channels.

// Input: Keyword "Go" in files:

// file1: "Go is fun"

// file2: "I love Go"

// file3: "Hello World"

// Expected Output: "Go is fun" from file1, "I love Go" from file2
func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	outChan := make(chan string)
	var wg sync.WaitGroup
	for _, v := range files {
		wg.Add(1)
		go CheckGoKeyword(v, outChan, &wg)
	}
	go func() {
		defer close(outChan)
		wg.Wait()
	}()
	for v := range outChan {
		fmt.Println(v)
	}
}

func CheckGoKeyword(path string, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	//by,err:=io.ReadFile(path)
	by, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(by)
	//scanner.Buffer()
	//scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		v := scanner.Text()
		if strings.Contains(v, "Go") {
			out <- v
		}
	}
}
