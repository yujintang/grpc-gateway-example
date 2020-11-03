package main

import (
	"context"
	"log"
	pb "mastiff/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":1718", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	mastiff := pb.NewMastiffClient(conn)

	data, err := mastiff.CreateId(context.Background(), &pb.CreateIdReq{Prefix: "dml:"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(data)
}
