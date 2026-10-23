package main

import (
	"fmt"
	"sync"
	"time"
)

// a send-only channel parameter: this function can only write to out
func square(n int, out chan<- int) {
	out <- n * n
}

// a receive-only channel parameter; range stops when the channel is closed
func sum(in <-chan int) int {
	total := 0
	for v := range in {
		total += v
	}
	return total
}

// worker pool: several goroutines read jobs from one channel
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- fmt.Sprintf("worker %d did job %d", id, j)
	}
}

func Concurrency() {
	// go starts a goroutine; an unbuffered channel waits until both sides are ready
	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		go square(i, ch)
	}
	fmt.Println("squares:", <-ch+<-ch+<-ch)

	// buffered channels accept sends until full without a receiver
	buf := make(chan int, 3)
	buf <- 1
	buf <- 2
	buf <- 3
	close(buf)
	fmt.Println("sum of buffered:", sum(buf))

	// receiving from a closed channel returns the zero value and ok == false
	v, ok := <-buf
	fmt.Println("closed channel:", v, ok)

	// WaitGroup waits for a group of goroutines to finish
	jobs := make(chan int, 5)
	results := make(chan string, 5)
	var wg sync.WaitGroup
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(results)
	fmt.Println("results:", len(results))

	// Mutex protects shared data from concurrent writes
	var mu sync.Mutex
	counter := 0
	var cwg sync.WaitGroup
	for i := 0; i < 100; i++ {
		cwg.Add(1)
		go func() {
			defer cwg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	cwg.Wait()
	fmt.Println("counter:", counter)

	// select waits on several channels; time.After gives a timeout
	slow := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond)
		slow <- "done"
	}()
	select {
	case msg := <-slow:
		fmt.Println("received:", msg)
	case <-time.After(10 * time.Millisecond):
		fmt.Println("timed out")
	}
}
