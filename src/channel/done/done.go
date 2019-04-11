package main

import (
    "fmt"
	"sync"
)

func doWork(id int,c chan int, wg *sync.WaitGroup){
    for n := range c {
        fmt.Printf("Worker %d received %c\n",id, n)
        wg.Done()
        //go func(){done <- true}()
    }
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
        wg: wg,
	}
	go doWork(id, w.in, wg)
	return w
}

type worker struct{
	in chan int
	//done chan bool
	wg *sync.WaitGroup
}

func chanDemo(){
	var wg sync.WaitGroup
	var workers [10]worker
	for i:=0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
}
	wg.Add(20)
	for i,worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}
	for i,worker := range workers{
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	wg.Wait()
}

func main() {
    chanDemo()
}
