package cmd

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ravilushqa/otelgqlgen"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/productsvc/domain"
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
	ctx := context.Background()

	tp := lib.InitTracer("product-svc")
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

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
			Tracer: domain.Tracer,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(r))
	srv.Use(otelgqlgen.Middleware())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
