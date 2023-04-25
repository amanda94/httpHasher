package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestDoConcurrently(t *testing.T) {

	finish := make(chan bool)
	go func() {
		DoConcurrently(2, []string{
			"http://www.aaa.com",
			"http://www.ddd.com",
			"http://www.google.com",
			"http://www.adjust.com",
			"http://www.bing.com",
		})
		finish <- true
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Printf("Number of goroutines running: %d\n", runtime.NumGoroutine())
		}

	}()
	<-finish

}

func TestDoConcurrently_Empty(t *testing.T) {

	finish := make(chan bool)
	go func() {
		DoConcurrently(2, []string{})
		finish <- true
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Printf("Number of goroutines running: %d\n", runtime.NumGoroutine())
		}

	}()
	<-finish

}
