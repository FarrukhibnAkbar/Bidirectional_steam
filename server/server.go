package main

import(
	"io"
	"log"
	"net"
	"app/proto"
	db "app/database"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedAirportServer
}


func (*Server) Status(s proto.Airport_StatusServer) error {

	ctx := s.Context ()

	st := db.GetBase() 

	for{
		select {
		case <- ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := s.Recv()

		if err == io.EOF {
			log.Fatalf("Exit")
			return nil
		}

		if st[req.Id].Status {
			if err := s.Send(&proto.Res{ Result: st[req.Id].Description}); err != nil {log.Fatalf("%v", err)} 
		}

		log.Println(req.Id)
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("%v", err)
	}

	s := grpc.NewServer()

	proto.RegisterAirportServer(s, &Server {})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("%v", err)
	}
}