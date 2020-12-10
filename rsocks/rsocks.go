package rsocks

import (
	"io"
	"log"
)

type Server struct {
}

func (s *Server) TeleConn(srv TeleConn_TeleConnServer) error {
	log.Println("start new server")
	ctx := srv.Context()
	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		log.Printf("%s", string(req.Body))

		//  send  to stream
		resp := Message{Body: []byte("Something")}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("sent %s", string(resp.Body))
	}

}
