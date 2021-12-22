package main

import (
	"enigmacamp.com/omzetsrv/api"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "enigma"
	dbUser := "postgres"
	dbPassword := "P@ssw0rd"
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected")
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")
	serverInfo := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", serverInfo)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := OmzetServer{
		db: db,
	}
	grpcServer := grpc.NewServer()
	api.RegisterOmzetServer(grpcServer, &s)
	log.Println("Server runs on", serverInfo)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
