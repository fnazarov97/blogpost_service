package main

import (
	"blockpost/config"
	"blockpost/genprotos/article"
	"blockpost/genprotos/author"
	aService "blockpost/services/article"
	dService "blockpost/services/author"
	"blockpost/storage"
	"blockpost/storage/postgres"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf := config.Load()
	AUTH := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDatabase,
	)
	var inter storage.StorageI
	inter, err := postgres.InitDB(AUTH)
	if err != nil {
		panic(err)
	}

	println("gRPC server running port:9000 with tcp protocol!")

	listener, err := net.Listen("tcp", conf.GRPCPort)
	if err != nil {
		panic(err)
	}

	c := &dService.AuthorService{
		Stg: inter,
	}
	s := grpc.NewServer()
	author.RegisterAuthorServicesServer(s, c)

	c1 := &aService.ArticleService{
		Stg: inter,
	}
	article.RegisterArticleServicesServer(s, c1)
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
