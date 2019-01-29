package main

import (
	"fmt"
	"log"
	"net/rpc"
)

var set [2]int

func main() {
	
	var reply,a,b int
	log.Println("Enter numbers to multiply")
	fmt.Scan(&a)
	fmt.Scan(&b)

	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	client.Call("Task.Store",a, &reply)
	client.Call("Task.Store",b, &reply)
	client.Call("Task.Multiply",1,&reply)
	log.Printf ("Numbers multiplied are %v, %v  " , a, b)
	log.Println ("Answer is ", reply)

}