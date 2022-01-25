package main

import (
	"fmt"
	"log"
	"net"

	"todo_server/src/utils"
	"todo_server/src/todo"
	"google.golang.org/grpc"
)

const(
	connType string = "tcp"
	connPort int = 8080
)

func main(){
	log.Println("Starting server...")

	lis, err := net.Listen(connType, fmt.Sprintf(":%d", connPort))
	utils.Check(err)

	s := todo.Server{}

	grpcServer := grpc.NewServer()

	todo.RegisterToDoServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	utils.Check(err)
}
