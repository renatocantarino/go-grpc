package main

import (
	"net"

	"github.com/renatocantarino/go-grpc/internals"
	"github.com/renatocantarino/go-grpc/internals/database"
	"github.com/renatocantarino/go-grpc/internals/pb"
	"github.com/renatocantarino/go-grpc/internals/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := internals.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	reflection.Register(grpcServer)

	lister, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lister); err != nil {
		panic(err)
	}

}
