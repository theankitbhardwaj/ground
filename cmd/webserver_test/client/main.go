package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	t := time.Now()
	fmt.Println("Client started")
	wg.Add(1)
	go func() {
		defer wg.Done()
		startSendingRequests(1)
	}()

	// wg.Add(1)
	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	defer wg.Done()
	// 	startSendingRequests(2)
	// }()
	wg.Wait()

	fmt.Printf("Time spent by client: %v\n", time.Since(t))
}

func startSendingRequests(id int8) {
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn was reused: %t", info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	req, err := http.NewRequestWithContext(traceCtx, "GET", "http://localhost:8080/hello", nil)

	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Keep-Alive", "timeout=5, max=1000")

	if err != nil {
		fmt.Printf("P#1WRTZ2: %v", err)
		return
	}
	for i := 0; i < 10000000; i++ {

		// time.Sleep(1 * time.Second)
		fmt.Printf("worker %v sending request %v\n", id, i)
		t := time.Now()
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Printf("P#1Y5S4M: %v\n", err)
			return
		}
		if _, err := io.Copy(io.Discard, res.Body); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("for worker %v request %v : time elapsed %v : status %v\n", id, i, time.Since(t), res.StatusCode)
		res.Body.Close()
	}
}
