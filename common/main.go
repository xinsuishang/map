package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	common          = "suaC1KQbM7PeYpF4p5vlhvfDgban1T0bPvSbN1Mi"
	ret             = "2nU83IRk8VBzKDJX9SlxPIzaqgg="
	testCounter     int64
	generateCounter int64
	match           = false
)

func matchTest(s, t string, wg *sync.WaitGroup) {
	defer wg.Done()
	data := []byte(s)
	hash := sha1.New()
	hash.Write(data)
	atomic.AddInt64(&testCounter, 1)
	value := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	if value == t {
		match = true
		fmt.Println(strings.Replace(s, common, "", -1))
	}
}

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	sChan := make(chan string, 100)

	func() {
		wg.Add(1)
		go recursion(sChan, "", 3, 6, wg)
		wg.Done()
	}()

	go func() {
		for {
			tmp := <-sChan
			matchTest(tmp+common, ret, wg)
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("cost: %s, testCounter: %d, generateCounter: %d\n", time.Since(start), testCounter, generateCounter)
		}
	}()
	time.Sleep(1 * time.Second)
	wg.Wait()
}

func recursion(sChan chan string, s string, min, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	if len(s) > max || match {
		return
	}
	if len(s) >= min {
		wg.Add(1)
		sChan <- s
		atomic.AddInt64(&generateCounter, 1)
	}
	func() {
		for i := 0; i < 26; i++ {
			recursion(sChan, s+string(i+97), min, max, wg)
		}
	}()
}
