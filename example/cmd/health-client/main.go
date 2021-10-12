package main

import (
	"context"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	address = "localhost:5105"
)

func main() {
	var err error
	c := &grpcServiceInfo{}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("Dial failed:", err)
		panic(err)
	}
	c.Conn = conn
	defer conn.Close()
	ctx := context.Background()

	healthClient := health.NewHealthClient(conn)

	timeout := time.Minute * 1
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := healthClient.Check(ctx, &health.HealthCheckRequest{})
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%v", response)
}

type grpcServiceInfo struct {
	Conn *grpc.ClientConn
}
