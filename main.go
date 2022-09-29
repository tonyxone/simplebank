package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tonyxone/simplebank/api"
	db "github.com/tonyxone/simplebank/db/sqlc"
	"github.com/tonyxone/simplebank/grpc"
	"github.com/tonyxone/simplebank/proto"
	"github.com/tonyxone/simplebank/util"
	grpcServer "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)

	runGrpcServer(config, store)
	//runGinServer(config, store)

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := grpc.NewServer(config, store)
	if err != nil {
		log.Fatal("connot create server:", err)
	}

	grpc := grpcServer.NewServer()
	proto.RegisterSimpleBankServer(grpc, server)
	reflection.Register(grpc)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start grpc server at %s", listener.Addr().String())
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal("cannot start listener:", err)

	}
}
