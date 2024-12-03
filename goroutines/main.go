package main

import (
	"fmt"
	"sync"
	"time"
)

// wait group is just a counter
var wg = sync.WaitGroup{}
var m = sync.Mutex{}
var dbdata = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	fmt.Println("Start")
	t0 := time.Now()

	for i := 0; i < len(dbdata); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("Execution time: ", time.Since(t0))
	fmt.Println("Results: ", results)
}

func dbCall(i int) {
	// simulate a DB call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("result is: ", dbdata[i])

	// locks memory from another process while one process uses it
	// lock the memory not the process or it will loose it concurrency
	m.Lock()
	results = append(results, dbdata[i])
	m.Unlock()

	wg.Done()
}
