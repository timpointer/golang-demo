package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/timpointer/golang-demo/rpc/netrpc/demo/api"
)

func main() {
	arith := new(api.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)

}
