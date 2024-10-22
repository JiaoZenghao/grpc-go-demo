package service

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	pb "jiaozenghao.gmail.com/grpc-go-demo/internal/service/product/proto"
)

type ProductServer struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductServiceServer
}

func (s *ProductServer) GetProduct(_ context.Context, in *wrapperspb.StringValue) (*pb.Product, error) {

	id := in.Value
	product := s.productMap[id]

	if product != nil {

		return product, nil

	}

	return nil, status.Errorf(codes.NotFound, "No record found.")
}

func (s *ProductServer) GetProductList(in *wrapperspb.StringValue, stream grpc.ServerStreamingServer[pb.Product]) error {

	log.Print("Get request with value: " + in.Value)
	productMap := s.productMap
	for key, product := range productMap {
		log.Print(key, product)
		err := stream.Send(product)
		if err != nil {

			return fmt.Errorf("error sending message to stream : %v",
				err)

		}
		log.Print("Sending Product Found : " + key)

	}

	return nil
}

func NewProductServer() *ProductServer {

	s := &ProductServer{productMap: make(map[string]*pb.Product)}

	// Populate the productMap with dummy data
	populateDummyProducts(s.productMap)

	return s

}

// populateDummyProducts adds dummy products to the provided product map
func populateDummyProducts(productMap map[string]*pb.Product) {
	productMap["1"] = &pb.Product{
		Id:          "1",
		Name:        "Smartphone",
		Description: "A high-end smartphone with a stunning display.",
	}

	productMap["2"] = &pb.Product{
		Id:          "2",
		Name:        "Laptop",
		Description: "A lightweight laptop with powerful processing capabilities.",
	}

	productMap["3"] = &pb.Product{
		Id:          "3",
		Name:        "Headphones",
		Description: "Noise-cancelling over-ear headphones with premium sound quality.",
	}

	productMap["4"] = &pb.Product{
		Id:          "4",
		Name:        "Smartwatch",
		Description: "A smartwatch with fitness tracking and notification features.",
	}

	productMap["5"] = &pb.Product{
		Id:          "5",
		Name:        "Tablet",
		Description: "A 10-inch tablet with a long-lasting battery and crystal-clear display.",
	}
}
