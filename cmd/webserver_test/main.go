package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/techrail/ground"
	"github.com/valyala/fasthttp"
)

var wg sync.WaitGroup
var t time.Time

func helloWorld(ctx *fasthttp.RequestCtx) {
	fmt.Print("Started serving request...\n")
	time.Sleep(10 * time.Second)
	fmt.Fprintf(ctx, "Hi there! Request Path is %q", ctx.Path())
	fmt.Print("Request served...\n")
}

func main() {
	server := ground.GiveMeAWebServer()

	server.BindPort = 8080
	server.Router.Handle(http.MethodGet, "/hello", helloWorld)
	server.BlockOnStart = true

	wg.Add(1)
	go func() {
		defer wg.Done()
		t = time.Now()
		server.Start()
	}()

	time.Sleep(15 * time.Second)
	fmt.Printf("Shutdown request after: %v\n", time.Since(t))
	server.Stop()

	wg.Wait()
	fmt.Printf("Time spent by server: %v\n", time.Since(t))
	fmt.Print("Server Shutdown...")
}
