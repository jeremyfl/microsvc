package cmd

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/productsvc/graph"
	"gitlab.com/jeremylo/microsvc/productsvc/graph/generated"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func loadGrpcConnection() *grpc.ClientConn {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return conn
}

func Serve() {
	conn := loadGrpcConnection()
	defer conn.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	stockRpc := stock.NewStockServiceClient(conn)

	db := initDatabase()
	entities := InitEntities(db, stockRpc)

	r := generated.Config{
		Resolvers: &graph.Resolver{
			Services: entities,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(r))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
