package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	prdSrv "jiaozenghao.gmail.com/grpc-go-demo/internal/service/product"
	product "jiaozenghao.gmail.com/grpc-go-demo/internal/service/product/proto"
	userSrv "jiaozenghao.gmail.com/grpc-go-demo/internal/service/user"
	user "jiaozenghao.gmail.com/grpc-go-demo/internal/service/user/proto"
)

func StartGrpcServer() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	product.RegisterProductServiceServer(grpcServer, prdSrv.NewProductServer())
	user.RegisterUserServiceServer(grpcServer, userSrv.NewUserServer())
	grpcServer.Serve(lis)

}
