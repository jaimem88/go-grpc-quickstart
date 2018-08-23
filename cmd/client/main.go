package main

import (
	"bufio"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jaimemartinez88/go-grpc-quickstart/proto"

	"github.com/jaimemartinez88/go-grpc-quickstart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	serverAddr          = "localhost:50051"
	grpcClientKeepalive = 120 * time.Second
)

func main() {

	conn, err := grpc.Dial(serverAddr,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: grpcClientKeepalive, PermitWithoutStream: true}),
	)
	if err != nil {
		log.Fatalf("failed to dial org grpc : %s %s", serverAddr, err)
	}

	echoClient := echo.NewEchoClient(conn)
	c := quickstart.NewClient(echoClient)

	log.Infoln("starting client...")
	mes := make(chan string, 1)
	ch := read(os.Stdin) //Reading from Stdin
	log.Info("type message: ")

	go func() {
		for m := range ch {
			t1 := time.Now()
			response, err := c.Echo(m)
			if err != nil {
				mes <- err.Error()
				continue
			}
			mes <- response + "\nduration: " + time.Since(t1).String()
		}
	}()
	for anu := range mes {
		log.Printf(anu) //Writing to Stdout
		log.Info("type message: ")

	}
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh
	log.Info("client stopped...")
}

/* Function to run the groutine to run for stdin read */
func read(r io.Reader) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			s := scan.Text()
			lines <- s
		}
	}()
	return lines
}
