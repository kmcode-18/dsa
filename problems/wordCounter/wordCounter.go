package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	files := []string{"large_file1.txt", "large_file2.txt", "large_file3.txt"}
	finalCounts := countWords(files)
	for word, count := range finalCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWords(files []string) map[string]int {
	out := make(map[string]int)
	chanOut := make(chan map[string]int)
	var wg sync.WaitGroup
	for _, f := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			wordCounter(file, chanOut)
		}(f)
	}
	go func() {
		wg.Wait()
		close(chanOut)
	}()
	for fOut := range chanOut {
		for w, c := range fOut {
			out[w] += c
		}
	}
	return out
}

func wordCounter(file string, chanOut chan map[string]int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening file %s : %v\n", file, err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	wordCounter := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, w := range words {
			wrd := strings.ToLower(w)
			wordCounter[wrd]++
		}
	}
	chanOut <- wordCounter
}
