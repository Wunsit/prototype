package main

import (
	"log"
	"net/http"
	"server/user"

	"github.com/dgraph-io/dgo/protos/api"

	"github.com/dgraph-io/dgo"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Trying to dial grpc")
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	http.Handle("/users/", user.NewHandler(user.NewDGraphStore(dg)))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
