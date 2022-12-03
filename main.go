package main

import (
	"blockpost/config"
	"blockpost/genprotos/author"
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

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	c := &dService.AuthorService{
		Stg: inter,
	}
	s := grpc.NewServer()
	author.RegisterAuthorServicesServer(s, c)
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
