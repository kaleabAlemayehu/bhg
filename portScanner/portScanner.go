package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		go worker(ports, results, &wg)
	}

	// Send ports to be scanned
	go func() {
		for i := 1; i < 10240; i++ {
			ports <- i
		}
		close(ports) // Close ports when done sending
	}()

	// Collect results
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Close results after all workers are done
	}()

	// Gather open ports from the results channel
	for port := range results {
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// Print the results
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

func worker(ports chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when done
	for p := range ports {
		address := fmt.Sprintf("localhost:%d", p)
		fmt.Printf("scanning... %s \n", address)
		con, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0 // Port closed
			continue
		}
		con.Close()
		results <- p // Port open
	}
}
