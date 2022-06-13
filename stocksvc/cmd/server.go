package cmd

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"stocksvc/stock"
)

func Serve() {
	db := initDatabase()
	entities := InitEntities(db)

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcHandler := stock.Server{
		Services: entities,
	}

	grpcServer := grpc.NewServer()

	stock.RegisterStockServiceServer(grpcServer, &rpcHandler)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
