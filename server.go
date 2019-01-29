package main 

import (
		"log"
		"net/http"
		"net"
		"net/rpc"

)
var set [2]int
var counter int
type Task int

func (t *Task) Store(number int, reply *int) error{
	set[counter]=number
	*reply= number
	counter ++
	return nil
}

func (t *Task) Multiply(number int, reply *int) error{
	*reply= set[0]*set[1]
	counter =0
	
	return nil

	
}


func main() {
	counter =0
	task := new(Task)
	// Publish the receivers methods
	err := rpc.Register(task)
	
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}