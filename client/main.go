package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aleiphoenix/gopush/api/grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPushServiceClient(conn)

	i := 0

	go func(i *int) {
		for {

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			req := &pb.PushRequest{
				Payload: []byte("1"),
				Ext:     []byte("1"),
			}
			_, err := c.Push(ctx, req)
			if err != nil {
				log.Fatalf("could not push: %v", err)
			}
			*i++
			cancel()
			// log.Printf("response: %v\n", r.Status)
		}
	}(&i)

	go func(i *int) {
		prev := 0
		for {
			log.Printf("cnt: %d\n", *i-prev)
			prev = *i
			time.Sleep(1 * time.Second)
		}
	}(&i)

	for {
		time.Sleep(1 * time.Second)
	}
}
