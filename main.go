package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
)

func main() {
	parallel := flag.Int("parallel", 10, "the largest number of parallel requests")
	flag.Parse()
	urls := flag.Args()
	// fmt.Println("args", parallel, urls)

	DoConcurrently(*parallel, urls)

}

func DoConcurrently(parallel int, urls []string) {
	taskChannel := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func(tasks chan string) {
			// fmt.Println("--------a goroutine start", runtime.NumGoroutine())
			defer func() {
				//	fmt.Println("--------a goroutine end", runtime.NumGoroutine())
				wg.Done()
			}()
			for task := range tasks {
				if !strings.HasPrefix(task, "http") {
					task = fmt.Sprintf("http://%s", task)
				}
				res, e := Hash(task)
				if e != nil {
					fmt.Printf("%s %v\n", task, e)
				} else {
					fmt.Printf("%s %s\n", task, res)
				}
			}
		}(taskChannel)
	}

	for _, u := range urls {
		taskChannel <- u
	}
	close(taskChannel)
	wg.Wait()
}
