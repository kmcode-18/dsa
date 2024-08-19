package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Define two matrices A and B
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	B := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}
	t := time.Now()
	C, err := ParallelMatrixMultiplication(A, B)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(time.Since(t).Nanoseconds())

	// Print the result matrix
	fmt.Println("Resultant Matrix C:")

	for _, row := range C {
		fmt.Println(row)
	}
}

// Function to multiply two matrices concurrently
func ParallelMatrixMultiplication(A, B [][]int) ([][]int, error) {
	//get dimension of matrices
	n, m := len(A), len(A[0])
	p, q := len(B), len(B[0])
	if m != p {
		return nil, fmt.Errorf("incompatible matrix dimensions: A is %dx%d, B is %dx%d", n, m, p, q)
	}
	// Initialize the result matrix
	c := make([][]int, n)
	for i := range c {
		c[i] = make([]int, q)
	}
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		for j := 0; j < q; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				c[i][j] = 0
				for k := 0; k < m; k++ {
					c[i][j] += A[i][k] * B[k][j]
				}
			}(i, j)
		}
	}
	wg.Wait()
	return c, nil
}
