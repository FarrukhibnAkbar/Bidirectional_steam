package main

import(
	"io"
	"log"
	"time"
	"context"
	"math/rand"
	"app/proto"
	"google.golang.org/grpc"
)


func main() {

	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("%v", err)
	}

	stub := proto.NewAirportClient(conn)

	stream, err := stub.Status(context.Background())

	ctx := stream.Context()

	if err != nil {
		log.Fatalf("%v", err)
	}

	done := make(chan bool)

	go func (){
		for {
			random := int64(rand.Intn(9))

			req := proto.Req { Id: random}

			err := stream.Send(&req)
		
			if err != nil {
				log.Fatalf("%v", err)
			}

			time.Sleep(time.Second * 3)
		}
	}()

	go func() {

		for{

			res, err := stream.Recv()

			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				log.Fatalf("%v", err)
			}	

			log.Println(res.Result)
		}
	}()

	go func() {
		<-ctx.Done()

		if err := ctx.Err(); err != nil {log.Fatalf("%v", err)}

		close(done)
	}()

	<- done

}

