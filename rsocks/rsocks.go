package rsocks

import (
	"io"
	"log"
	"time"
)

type Server struct {
}

func (s *Server) TeleConn(srv TeleConn_TeleConnServer) error {
	log.Println("start new server")
	ctx := srv.Context()
	go func(lsrv TeleConn_TeleConnServer) {
		for {
			req, err := lsrv.Recv()
			if err == io.EOF {
				log.Println("EOF, exiting the connection...")
				return
			}
			if err != nil {
				log.Printf("receive error %v", err)
				return
			}
			log.Printf("%s", string(req.Body))
		}
	}(srv)

	go func(lsrv TeleConn_TeleConnServer) {
		for {
			s := time.Now().String()
			resp := Message{Body: []byte(s)}
			if err := lsrv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			time.Sleep(3 * time.Second)
		}

	}(srv)
	<-ctx.Done()
	return ctx.Err()
}
