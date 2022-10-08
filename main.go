package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":10000")
	fmt.Println("This is grpc Server")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	//geolocation server
	ch1 := geolocation.Server{}
	geolocation.RegisterGeoLocationServiceServer(grpcServer, &ch1)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
