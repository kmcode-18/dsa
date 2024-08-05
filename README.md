# dsa
Data Structures and Algorithms


#TODO:
Line Counter in Files using Concurrency
Task: Write a Go program that counts the number of lines in multiple text files concurrently. Use goroutines to read lines from each file and channels to aggregate the counts.
Parallel Matrix Multiplication
Task: Implement a program that multiplies two matrices concurrently. Divide the computation of the result matrix into multiple goroutines, each responsible for calculating a portion of the matrix.
Concurrent Prime Number Finder
Task: Write a Go program that finds all prime numbers up to a given number N using concurrency. Use a worker pool pattern to distribute the workload among multiple goroutines.
File Searcher with Concurrent Pattern Matching
Task: Develop a program that searches for a specific pattern in multiple files concurrently. Each file should be processed by a separate goroutine, and the program should aggregate the results using channels.
Concurrent Log File Analyzer
Task: Write a Go program that analyzes log files concurrently. For example, count the occurrences of different log levels (INFO, WARN, ERROR) across multiple log files using goroutines.
Concurrent Web Scraper
Task: Implement a web scraper that fetches and processes data from multiple URLs concurrently. Each URL should be handled by a separate goroutine, and results should be collected via channels.
Concurrent Image Processor
Task: Create a Go program that applies a series of image processing filters (e.g., grayscale, blur) to a list of images concurrently. Use goroutines to process each image in parallel.
Concurrent Directory File Size Calculator
Task: Write a program that calculates the total size of all files in a directory tree concurrently. Each subdirectory can be processed by a separate goroutine.
Concurrent Word Frequency Counter
Task: Extend the word counter program to count the frequency of each word across multiple files concurrently. Use a channel to aggregate the word counts from each goroutine.
Concurrent Web Server Log Analyzer
Task: Develop a program that processes web server logs concurrently to extract metrics like the number of requests per IP address, the most requested URLs, etc.
