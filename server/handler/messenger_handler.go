package handler

import (
	"context"
	"log"
	"time"

	pb "github.com/kkk777-7/grpc-web/messenger/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MessengerHandler struct {
	pb.UnimplementedMessengerServer
	requests []*pb.MessageRequest
}

func NewMessengerHandler() *MessengerHandler {
	return &MessengerHandler{}
}

func (s *MessengerHandler) GetMessages(_ *emptypb.Empty, stream pb.Messenger_GetMessagesServer) error {
	for _, r := range s.requests {
		if err := stream.Send(&pb.MessageResponse{Message: r.GetMessage()}); err != nil {
			return err
		}
	}

	previousCount := len(s.requests)

	for {
		currentCount := len(s.requests)
		if previousCount < currentCount && currentCount > 0 {
			r := s.requests[currentCount-1]
			log.Printf("Sent: %v", r.GetMessage())
			if err := stream.Send(&pb.MessageResponse{Message: r.GetMessage()}); err != nil {
				return err
			}
		}
		previousCount = currentCount
	}
}

func (s *MessengerHandler) CreateMessage(ctx context.Context, r *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Received: %v", r.GetMessage())
	newR := &pb.MessageRequest{Message: r.GetMessage() + ": " + time.Now().Format("2006-01-02 15:04:05")}
	s.requests = append(s.requests, newR)
	return &pb.MessageResponse{Message: r.GetMessage()}, nil
}
