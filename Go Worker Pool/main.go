package main

import (
	"fmt"
	"sync"
	"time"
)

type Work struct {
	amount int
	name   string
}

func (w *Work) GiveAmount() int {
	return w.amount
}

func worder(id int, workChan <-chan Work, reslutChan chan<- string) {
	for work := range workChan {
		reslutChan <- fmt.Sprintf("Worker %v: Received Work %v. Processing...\n", id, work.name)
		time.Sleep(time.Duration(work.GiveAmount()) * time.Second)
		reslutChan <- fmt.Sprintf("Worker %v: Work %v done.\n", id, work.name)
	}
	fmt.Printf("Worker %v: terminated.\n", id)
}

func main() {
	fmt.Println("Helllo")
	wg := sync.WaitGroup{}
	workChan := make(chan Work, 5)
	reslutChan := make(chan string, 10)
	defer close(reslutChan)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worder(i, workChan, reslutChan)
		}()
	}
	then := time.Now()
	workChan <- Work{
		name:   "work1",
		amount: 2,
	}
	workChan <- Work{
		name:   "work2",
		amount: 3,
	}
	workChan <- Work{
		name:   "work3",
		amount: 1,
	}
	workChan <- Work{
		name:   "work4",
		amount: 5,
	}
	workChan <- Work{
		name:   "work5",
		amount: 10,
	}
	workChan <- Work{
		name:   "work6",
		amount: 3,
	}
	close(workChan)
	i := 0
	for msg := range reslutChan {
		i++
		fmt.Println(msg)
		if i == 12 {
			break
		}
	}
	wg.Wait()
	fmt.Printf("Time: %v", time.Since(then))
	fmt.Println("All Workers terminated. Exiting...")
}
