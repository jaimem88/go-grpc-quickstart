package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	echo_pb "github.com/jaimemartinez88/go-grpc-quickstart/proto"

	"github.com/jaimemartinez88/go-grpc-quickstart"

	"google.golang.org/grpc"
)

const serverAddr = "localhost:50051"

func main() {

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", serverAddr)
	// error handling omitted
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
	echoServer := &quickstart.Server{}
	echo_pb.RegisterEchoServer(s, echoServer)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	log.Infof("server listening on: %s\n", serverAddr)
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh

	log.Println("server stopping...")
	s.GracefulStop()
	log.Println("server stopped...")
}

// func createGRPCServer(db *sql.DB, deviceProvisioningClient device_pb.ProvisioningServiceClient) (*grpc.Server, error) {
// 	s := grpc.NewServer(
// 		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
// 			grpctrace.UnaryServerInterceptor(grpctrace.WithServiceName(componentName)),
// 			grpc_ctxtags.UnaryServerInterceptor(),
// 			middleware.ReqIDAsTagInterceptor(),
// 			// grpc_logrus: logs stats of each end point along with errors
// 			grpc_logrus.UnaryServerInterceptor(mlog),
// 			grpc_recovery.UnaryServerInterceptor(), // recover from panic and send internal error
// 		)),
// 		grpc.KeepaliveParams(keepalive.ServerParameters{Time: grpcServerKeepalive}),
// 	)

// 	reflection.Register(s)

// 	return s, nil
// }
