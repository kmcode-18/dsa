package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int { return rand.Intn(54554541) }
	intStream := repeatFunc(done, randNumFetcher)
	// for str := range take(done, intStream, 20) {
	// 	fmt.Println(str)
	// }
	//native
	// primeStream := primeFinder(done, intStream)
	// for prime := range take(done, primeStream, 20) {
	// 	fmt.Println(prime)
	// }

	//fanout for int stream
	numCPUs := runtime.NumCPU()

	primeFinderChannels := make([]<-chan int, numCPUs)
	for i := 0; i < numCPUs; i++ {
		primeFinderChannels[i] = primeFinder(done, intStream)
	}

	//fanin of primes
	primes := fanIn(done, primeFinderChannels)
	for prime := range take(done, primes, 20) {
		fmt.Println(prime)
	}
	fmt.Println(time.Since(t))
}

func fanIn[T any, K any](done <-chan K, inputChanels []<-chan T) <-chan T {
	fanedInStream := make(chan T)
	var wg sync.WaitGroup
	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fanedInStream <- i:
			}
		}
	}

	for _, ch := range inputChanels {
		wg.Add(1)
		go transfer(ch)
	}
	go func() {
		defer close(fanedInStream)
		wg.Wait()
	}()
	return fanedInStream
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randInt int) bool {
		for i := randInt - 1; i > 1; i-- {
			if randInt%i == 0 {
				return false
			}
		}
		return true
	}
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case i := <-randIntStream:
				if isPrime(i) {
					primes <- i
				}
			}
		}
	}()
	return primes
}
func take[T any, K any](done <-chan K, inpStream <-chan T, n int) <-chan T {
	outStream := make(chan T)
	go func() {
		defer close(outStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case outStream <- <-inpStream:
			}
		}
	}()
	return outStream
}
func repeatFunc[T any, K any](done <-chan K, fu func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fu():
			}
		}
	}()
	return stream
}
