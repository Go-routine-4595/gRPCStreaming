package main

import (
	pb "gRPCStreaming/proto"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type EventService struct {
	pb.UnimplementedEventServer
}

func (e *EventService) GetEvent(r *pb.Request, s pb.Event_GetEventServer) error {

	for i := 0; ; i++ {
		result := &pb.Response{
			MatchedTerm: r.Term,
			Rank:        int32(i),
			Content:     "content: " + strconv.Itoa(i),
		}

		if err := s.Send(result); err != nil {
			log.Printf("Error sendign message to the client: %v", err)
			return err
		}

		d := time.Duration(rand.Intn(500))
		time.Sleep(d * time.Millisecond)
	}

	return nil
}
