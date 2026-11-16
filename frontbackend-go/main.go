package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "olio.ovh/yswsplat/protobufschema"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "asd"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.GetMessage())
}
