package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/timpointer/golang-demo/rpc/netrpc/demo/api"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &api.Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error :", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
	fmt.Println()
	quotient := new(api.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	fmt.Println("befor quo", quotient.Quo, quotient.Rem)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		log.Fatalf("client.Go %v", replyCall.Error)
	}

	fmt.Println("quo", quotient.Quo, quotient.Rem)
}
