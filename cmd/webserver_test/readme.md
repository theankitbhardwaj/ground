# Fasthttp Shutdown test
These files are created to test Fasthttp's Shutdown method.

## 1. ./main.go
1. This file starts the main server, which we will be testing.
2. Each request takes 5 secs to process.
3. Server shutdown is requested after 15 secs of server starting.

## 2. client/main.go
1. This file starts a routine, that sends multiple requests to our server one after the other.


## Results
* Client starts sending requests to server and it starts processing these requests. 
* After first request is processed, second request is sent
* After 5 secs of processing second request, server shutdown is requested
* Server doesn't shutdown for another 5 secs until that request is processed, In the meantime, server closes the listener and doesn't accept any further connections.
* After second request is processed, server shuts down and client also closes since it cannot connect to server.

`First request took 10 secs to process, server shutdown requested after 15 secs, but server actually shuts down after 20 secs, after processing first two requests.`

## Extras
* I tried to send the keep-alive header from client but it doesn't seem to effect the server shutdown, in both cases server shuts down similarly. I am not sure, if the way I tried testing Keep-alive is correct or not.
