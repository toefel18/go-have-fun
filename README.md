# Go-have-fun

This repository contains the runnable sources presented in the Golang article written by Rob Brinkman and Christophe Hesters 
published in Java Magazine Q1 2017.

Go (Golang) is a 'new' programming language that just turned 7 and it's quickly gaining traction (it's ranking went from 
the 54th place in Jan 2016 to the 13th place in Jan 2017 as stated in the [Tiobe index](http://www.tiobe.com/tiobe-index/)). 
Big projects such as Docker and Kubernetes are written in Go, as is much of Google's internal software. Go compiles quickly 
to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically 
typed, compiled language that feels like a dynamically typed, interpreted language. 

### Running the code from the article

To get started:

 1. Install [Golang](https://golang.org/dl/) and follow the installation instructions
 1. To help you set-up, watch this [5 minute screencast](https://www.youtube.com/watch?v=XCsL89YtqCs), it's worth it for newcomers ;)
 1. Get the sources and dependencies with one command: 
 
      `go get github.com/toefel18/go-have-fun`

 1. View and run the listings:

 * [listing 1 - time-service](timeservice/listing1-time-service.go)
 
        cd $GOPATH/src/github.com/toefel18/go-have-fun/timeservice
        go run listing1-time-service.go
 
 * [listing 2 - unit test](timeservice/listing2-unit-test_test.go)
  
        cd $GOPATH/src/github.com/toefel18/go-have-fun/timeservice
        go test 
        
 * [listing 3 - benchmark](timeservice/listing3-benchmark_test.go)
   
        cd $GOPATH/src/github.com/toefel18/go-have-fun/timeservice
        go test -bench .
        
 * [listing 4 - structs](structs/listing4-structs.go)
   
        cd $GOPATH/src/github.com/toefel18/go-have-fun/structs
        go run listing4-structs.go
        
 * [listing 5 - arrays](collections/listing5-arrays.go)
   
        cd $GOPATH/src/github.com/toefel18/go-have-fun/collections
        go run listing5-arrays.go
 
 * [listing 6 - maps](collections/listing6-maps.go)
   
        cd $GOPATH/src/github.com/toefel18/go-have-fun/collections
        go run listing6-maps.go
 
 * [listing 7 - channels](goroutines-channels/listing7-producers-1-consumer.go)
   
        cd $GOPATH/src/github.com/toefel18/go-have-fun/goroutines-channels
        go run listing7-producers-1-consumer.go
 
 * [listing 8 - interfaces](interfaces/listing-8-interfaces.go)
  
        cd $GOPATH/src/github.com/toefel18/go-have-fun/interfaces
        go run listing-8-interfaces.go
 
### Start learning Go

A really good starting point is to work through [A Tour of Go](https://tour.golang.org/welcome/1). 
A Tour of Go is an online interactive course developed by Google. The tour contains multiple modules explaining 
Go step-by-step. It contains simple assignments and allows you to test all the code in the browser! 
 
