package main

import (
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"

	"google.golang.org/grpc"
)

func main() {
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)
	conn, gerr := grpc.Dial("my-service", grpc.WithBalancer(b), grpc.WithBlock())
}
