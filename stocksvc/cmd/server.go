package cmd

import (
	"fmt"
	"gitlab.com/jeremylo/microsvc/stocksvc/handler"
	"gitlab.com/jeremylo/microsvc/stocksvc/stock"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Serve() {
	db := initDatabase()
	entities := InitEntities(db)

	go listener(entities)

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcHandler := handler.RpcHandler{
		Services: entities,
	}

	grpcServer := grpc.NewServer()

	stock.RegisterStockServiceServer(grpcServer, &rpcHandler)

	fmt.Println("Listening on port 9000 for rpc")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
