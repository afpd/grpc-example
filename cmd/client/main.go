package main

import (
	"fmt"

	"dev.home.arpa/devuser/grpc-example/rsocks"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

func HandleIcomingConnection(incomingConn net.Conn, grpcConn *grpc.ClientConn) error {
	log.Printf("Serving %s", incomingConn.RemoteAddr().String())
	c := rsocks.NewTeleConnClient(grpcConn)
	stream, err := c.TeleConn(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	go func() {
		for {

			br, err := stream.Recv()
			if err != nil {
				log.Panicln(err)
			}
			log.Printf("Received: %v", string(br.Body))
			incomingConn.Write(br.Body)
		}
	}()
	go func() {
		buff := make([]byte, 2048)
		for {
			n, err := incomingConn.Read(buff)
			if err != nil {
				return
			}
			if n > 0 {
				stream.Send(&rsocks.Message{Body: buff[0 : n-1]})
			}
		}

	}()
	return err
}

func main() {
	fmt.Println("Starting the client")
	var grpcConn *grpc.ClientConn
	grpcConn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer grpcConn.Close()
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleIcomingConnection(c, grpcConn)
	}

}
