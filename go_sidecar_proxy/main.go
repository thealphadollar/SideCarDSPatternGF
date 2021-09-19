package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/EwanValentine/distributed-patterns/sidecar-http-grpc/app/transport"
	grpc "google.golang.org/grpc"

	mux "github.com/gorilla/mux"
)

type API struct {
	service pb.ApplicationClient
}

func (api *API) Greet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := &pb.Request{
		Name: params["name"],
	}
	response, err := api.service.FetchGreeting(context.Background(), request)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, response.Reply)
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)

	addr := fmt.Sprintf("0.0.0.0:%d", 6000)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewApplicationClient(conn)
	api := API{
		service: client,
	}

	router.HandleFunc("/user/{name}", api.Greet).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
