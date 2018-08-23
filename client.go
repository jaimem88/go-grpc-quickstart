package quickstart

// func newClient() {
// 	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	// error handling omitted
// 	client := pb.NewGreeterClient(conn)
// 	// ...
// }

import (
	"context"
	"fmt"

	echo_pb "github.com/jaimemartinez88/go-grpc-quickstart/proto"
)

type Client struct {
	echoClient echo_pb.EchoClient
}

func NewClient(echoClient echo_pb.EchoClient) *Client {
	return &Client{
		echoClient: echoClient,
	}
}
func (c *Client) Echo(message string) (string, error) {
	ctx := context.Background()
	response, err := c.echoClient.Echo(ctx, &echo_pb.Request{Message: message})

	return fmt.Sprintf("server response: %s", response.GetMessage()), err
}
