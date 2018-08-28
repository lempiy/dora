package main

import (
	"github.com/lempiy/dora/shared/help"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

// TODO: move to env variables
const port = ":7000"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	dir := help.GetCurrentDir()
	creds, err := credentials.NewServerTLSFromFile(dir+"/cert/server.crt", dir+"/cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	log.Printf("Parser server is now listening on port %s", port)
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
