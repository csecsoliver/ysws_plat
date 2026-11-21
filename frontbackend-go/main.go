package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "olio.ovh/yswsplat/protobufschema"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	client := pb.NewReviewQueueClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	projTicket, err := client.InitProject(ctx, &pb.Project{Name: "asd", Description: "something"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(projTicket.GetUuid())

	projDetails, err := client.GetProject(ctx, projTicket)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(projDetails.GetName())
	
	subTicket, err := client.Enqueue(ctx, &pb.Submission{ProjectUuid: projTicket.GetUuid(), Comment: "comment"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(subTicket.GetUuid())
}

