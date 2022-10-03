package main

import (
	"context"
	"log"

	"./geolocation/geolocation"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := geolocation.NewGeoLocationServiceClient(conn)
	response, err := c.GetGeoLocationData(context.Background(), &geolocation.
		Message{Body: "Hello"})
	if err != nil {
		log.Fatalf("Error when retrieving GeoLocation Data: %s", err)
	}
	// display the response from the server
	log.Printf("Response from server: %s", response.Body)
}
