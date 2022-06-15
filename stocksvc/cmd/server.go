package cmd

import (
	"context"
	"fmt"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/stocksvc/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Serve() {
	ctx := context.Background()

	tp := lib.InitTracer("stock-svc-grpc")
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	db := initDatabase()
	entities := InitEntities(db)

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
