package quickstart

import (
	"context"

	echo_pb "github.com/jaimemartinez88/go-grpc-quickstart/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct{}

func (s *Server) Echo(ctx context.Context, request *echo_pb.Request) (*echo_pb.Response, error) {
	log.Infof("client message: %s\n", request.GetMessage())
	return &echo_pb.Response{
		Message: request.GetMessage(),
	}, nil
}
