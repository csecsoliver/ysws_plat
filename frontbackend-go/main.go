package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "olio.ovh/yswsplat/protobufschema"
)

var conn *grpc.ClientConn

func init_project(w http.ResponseWriter, req *http.Request, client pb.ReviewQueueClient) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "!err:%s: %s", "Method not allowed", req.Method)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "!err:%s", err)
		return
	}
	props := strings.Split(string(body), ";")
	if len(props) != 2 {
		fmt.Fprintf(w, "!err:%s, %d, %s", "Malformed request", len(props), body)
		return
	}
	if len(props[0]) == 0 {
		fmt.Fprintf(w, "!err:%s", "Empty field: title.")
		return
	}
	projTicket, err := client.InitProject(ctx, &pb.Project{Name: props[0], Description: props[1]})
	if err != nil {
		fmt.Fprintf(w, "!err:%s", err)
		return
	}
	fmt.Fprintf(w, "!suc:%s", projTicket.GetUuid())
}
func get_project(w http.ResponseWriter, req *http.Request, client pb.ReviewQueueClient) {
	if req.Method != "GET" {
		fmt.Fprintf(w, "!err:%s: %s", "Method not allowed", req.Method)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	uuid := req.URL.Query().Get("uuid")
	if len(uuid) == 0 {
		fmt.Fprintf(w, "!err:%s", "Bad request")
		return
	}
	proj, err := client.GetProject(ctx, &pb.ProjectTicket{Uuid: uuid})
	if err != nil {
		fmt.Fprintf(w, "!err:%s", err)
		return
	}
	fmt.Fprintf(w, "!suc:%s;%s", proj.GetName(), proj.GetDescription())
}

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

	http.HandleFunc("/init_project", func(w http.ResponseWriter, r *http.Request) {
		init_project(w, r, client)
	})
	http.HandleFunc("/get_project", func(w http.ResponseWriter, r *http.Request) {
		get_project(w, r, client)
	})
	log.Println("listening.")
	err = http.ListenAndServe("0.0.0.0:8090", nil)
	log.Println(err)

	// projTicket, err := client.InitProject(ctx, &pb.Project{Name: "asd", Description: "something"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(projTicket.GetUuid())

	// projDetails, err := client.GetProject(ctx, projTicket)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(projDetails.GetName())

	// subTicket, err := client.Enqueue(ctx, &pb.Submission{ProjectUuid: projTicket.GetUuid(), Comment: "comment"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(subTicket.GetUuid())
}
