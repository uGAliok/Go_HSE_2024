package main

import (
	"Go_HSE_2024/2_and_3_HW_server/accounts/models"
	"Go_HSE_2024/2_and_3_HW_server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	proto.UnimplementedBankServer
	database *models.BankDatabase
}

// FUNCTIONS
func (s *Server) CreateAccount(ctx context.Context,
	req *proto.CreateAccountRequest) (*proto.CreateAccountReply, error) {
	fmt.Printf("Received: %s, %d\n", req.Name, req.Amount)
	if err := s.database.CreateAccount(req.Name, req.Amount); err != nil {
		return nil, err
	}
	return &proto.CreateAccountReply{Message: fmt.Sprintf("Account created", req.Name, req.Amount)}, nil
}

func (s *Server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountReply, error) {
	fmt.Printf("Received: %s\n", req.Name)
	acc, err := s.database.GetAccount(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.GetAccountReply{Message: fmt.Sprintf("Name: %s\tAmount: %d", acc.Name, acc.Amount)}, nil
}

func (s *Server) DeleteAccount(ctx context.Context,
	req *proto.DeleteAccountRequest) (*proto.DeleteAccountReply, error) {
	fmt.Printf("Received: %s\n", req.Name)
	if err := s.database.DeleteAccount(req.Name); err != nil {
		return nil, err
	}
	return &proto.DeleteAccountReply{Message: fmt.Sprintf("Deleted", req.Name)}, nil
}

func (s *Server) Patch(ctx context.Context,
	req *proto.PatchRequest) (*proto.PatchReply, error) {
	fmt.Printf("Received: %s, %d\n", req.Name, req.Amount)
	if err := s.database.Patch(req.Name, req.Amount); err != nil {
		return nil, err
	}
	return &proto.PatchReply{Message: fmt.Sprintf("%d on account '%s'", req.Amount, req.Name)}, nil
}

func (s *Server) UpdateName(ctx context.Context, req *proto.UpdateNameRequest) (*proto.UpdateNameReply, error) {
	fmt.Printf("Received: %s, %s\n", req.Name, req.NewName)
	err := s.database.UpdateName(req.Name, req.NewName)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateNameReply{Message: fmt.Sprintf("New name of account '%s' ('%s')", req.Name, req.NewName)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5678))

	if err != nil {
		panic(err)
	}
	ser := grpc.NewServer()
	proto.RegisterBankServer(ser, &Server{database: models.TheDatabase()})
	if err := ser.Serve(lis); err != nil {
		panic(err)
	}
}
