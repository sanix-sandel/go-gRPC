package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/client/ecommerce"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	//Set up a conection with the server from the provided address
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//pass the connection and create a stub
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	description := `Meet apple iPhone 11. All-New dual-camera system with Ultra Wide and Night mode.`
	//	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx,
		&pb.Product{Name: name, Description: description})

	if err != nil {
		log.Fatal("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: ", product.String())

}
