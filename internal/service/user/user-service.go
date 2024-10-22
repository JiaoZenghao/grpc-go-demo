package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	pb "jiaozenghao.gmail.com/grpc-go-demo/internal/service/user/proto"
)

type UserServer struct {
	userMap map[string]*pb.User
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(_ context.Context, in *wrapperspb.StringValue) (*pb.User, error) {
	id := in.Value
	user := s.userMap[id]

	if user != nil {

		return user, nil

	}

	return nil, status.Errorf(codes.NotFound, "No record found.")
}

func NewUserServer() *UserServer {

	s := &UserServer{userMap: make(map[string]*pb.User)}

	// Populate the productMap with dummy data
	populateDummyUsers(s.userMap)

	return s

}

// populateDummyProducts adds dummy products to the provided product map
func populateDummyUsers(userMap map[string]*pb.User) {
	userMap["1"] = &pb.User{
		Id:   "1",
		Name: "Abert",
	}

	userMap["2"] = &pb.User{
		Id:   "2",
		Name: "Bill",
	}

	userMap["3"] = &pb.User{
		Id:   "3",
		Name: "Candy",
	}

	userMap["4"] = &pb.User{
		Id:   "4",
		Name: "Daniel",
	}

	userMap["5"] = &pb.User{
		Id:   "5",
		Name: "Emliy",
	}
}
