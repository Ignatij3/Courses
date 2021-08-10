package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id			int
	n			int64
	from, to 	int64
}
type Result struct {
	job         Job
	divisor		int64  // 0, if n has no divisors in [job.from; job.to
}

var jobs = make(chan Job, 25)
var results = make(chan Result, 25)
var terminated bool

func worker(wg *sync.WaitGroup) {
	var divisor int64
	for job := range jobs {
		if terminated {break}
		divisor = 0
		for i := job.from; i < job.to; i++  {
			if job.n%i == 0  { 
				divisor = i
				break
			}
		}
		results <- Result{job, divisor}
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func analysis(done chan bool) {
	prime:= true
	for result := range results {
		fmt.Printf("Job id %d:  ", result.job.id)
		if result.divisor > 0 {
			fmt.Printf("divisor %d is found. ", result.divisor)
			prime = false
			terminated = true
		}  else  {
			fmt.Printf("There aren't divisors on [%d ; %d). ", result.job.from, result.job.to)
		}	 	
		fmt.Println()
	}
	done <- prime
}

func allocate(n int64) {
	var start, step, finish int64
	step = 1234567
	jobid:= 0
	for start = 2; start*start <= n; start = finish {
		jobid++
		finish = start + step
		job := Job{jobid, n, start, finish}
		jobs <- job
	}
	close(jobs)
}

func main() {
	startTime := time.Now()

	go allocate(1000000087*1120000093)	

	done := make(chan bool)
	terminated = false
	go analysis(done)
	createWorkerPool(20)  // 20 - number of Workers
	fmt.Println(<-done)

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
