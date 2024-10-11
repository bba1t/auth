package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/bba1t/auth/pkg/user_v1"
)

const (
	grpcHost = "localhost:"
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedUserV1Server
}

func main() {
	lis, err := net.Listen("tcp", grpcHost+strconv.Itoa(grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Println(desc.CreateRequest{
		Name:            req.GetName(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
		UserType:        req.GetUserType(),
	})
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Println(desc.GetRequest{
		Id: req.GetId(),
	})
	return &desc.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		UserType:  desc.Role_admin,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	log.Println(desc.UpdateRequest{
		Id:    req.GetId(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	})
	return &empty.Empty{}, nil
}

func (s *server) Delete(_ context.Context, r *desc.DeleteRequest) (*empty.Empty, error) {
	log.Println("id: ", r.GetId())
	return &empty.Empty{}, nil
}
