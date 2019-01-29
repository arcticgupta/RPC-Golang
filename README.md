# RPC-Golang

This is an example to a Remote Procedural Call using Golang and it's out-of-box library (net)

--------------------------------------
To run:
1. ```go run server.go``` to run the server on the default port (can be changed).
2. ```go run client.go``` on a different terminal to make a remote call to the already set-up server.
--------------------------------------

Rules to construct a method in a RPC server:
1. the method's type is exported.
2. the method is exported.
3. the method has two arguments, both exported (or builtin) types.
4. the method's second argument is a pointer.
5. the method has return type error.

```||||||||server.go||||||||```

```
var set [2]int
var counter int
type Task int
```
Array to store passed variables. Counter to keep track of the number of calls. Data type Task to pass to the called methods.


```func (t *Task) Store(number int, reply *int) ``` to add passed variables to array

```func (t *Task) Multiply(number int, reply *int)``` to multiply elements in the constructed array

```
task := new(Task)
err := rpc.Register(task)
```
to create and register a new task

```	rpc.HandleHTTP()``` make a request handler for the RPC

```	listener, e := net.Listen("tcp", ":1234")``` listens for incoming connections

```	err = http.Serve(listener, nil)``` serves back the reply

-----------------------------------

```||||||||client.go||||||||```
```
var reply,a,b int
log.Println("Enter numbers to multiply")
fmt.Scan(&a)
fmt.Scan(&b)
```
Asks user for input.

```	client, err := rpc.DialHTTP("tcp", "localhost:1234")``` Create a TCP conncetion to localhost on port 1234

```client.Call()``` calls the mentioned method on the server

----------------------------------





