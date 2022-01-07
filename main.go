package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Uchencho/core-accounts/internal/app"
	"github.com/Uchencho/core-proto/generated/accounts"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = "4444"
)

func getListener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
	return lis
}

func main() {

	mockInsert := func(oa *app.Option) {
		oa.InsertAccount = func(a accounts.Account) error {
			return nil
		}
	}

	a := app.NewApp(mockInsert)
	grpcServer := grpc.NewServer()

	accounts.RegisterClientServer(grpcServer, a)

	log.Println(fmt.Sprintf("Starting server on address: %s:%s", host, port))
	if err := grpcServer.Serve(getListener()); err != nil {
		log.Fatal("failed to serve: " + err.Error())
	}
}
