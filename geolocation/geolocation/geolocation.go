package geolocation

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) GetGeoLocationData(ctx context.Context, in *Message) *Message {
	log.Println("Incoming Geolocation Request")
	return &Message{Body: "Hello"}, nil
}
