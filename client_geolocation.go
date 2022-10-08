package main

import (
	"context"
	"log"

	"github.com/rajajamal/API-Gateway/geolocation/geolocation"
	"google.golang.org/grpc"
)

func geoServer() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := geolocation.NewGeoLocationServiceClient(conn)
	response, err := c.GetGeoLocationData(context.Background(), &geolocation.
		Pesan{Body: "Hello"})
	if err != nil {
		log.Fatalf("Error when retrieving GeoLocation Data: %s", err)
	}
	// display the response from the server
	log.Printf("Response from server: %s", response.Body)
}
