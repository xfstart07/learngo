// Author: xufei
// Date: 2020/7/2

package department

import (
	"context"
	"learngo/microservices/gokit/one-part/config"
	"learngo/microservices/gokit/one-part/protobuf-spec/common"
	"log"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type DManagerServiceServer struct {
	CreateHandler       transport_grpc.Handler
	ListHandler         transport_grpc.Handler
	PersonChangeHandler transport_grpc.Handler
}

func (s *DManagerServiceServer) List(ctx context.Context, req *common.Empty) (*common.DepartmentList, error) {
	_, resp, err := s.ListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*common.DepartmentList), nil
}

func (s *DManagerServiceServer) Create(ctx context.Context, req *common.Department) (*common.Empty, error) {
	_, resp, err := s.CreateHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*common.Empty), nil
}

func (s *DManagerServiceServer) PersonChange(ctx context.Context, req *common.PersonChangeRequest) (*common.Empty, error) {
	_, resp, err := s.PersonChangeHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*common.Empty), nil
}

func NewGrpcServer(srv *DepartmentManager) *DManagerServiceServer {
	listHandler := transport_grpc.NewServer(
		makeListEndpoint(srv),
		DecodeGrpcRequest,
		EndcodeGrpcResponse,
	)

	createHandler := transport_grpc.NewServer(
		makeCreateEndpoint(srv),
		DecodeGrpcRequest,
		EndcodeGrpcResponse)

	personChangeHandler := transport_grpc.NewServer(
		makePersonChangeEndpont(srv),
		DecodeGrpcRequest,
		EndcodeGrpcResponse)

	return &DManagerServiceServer{
		ListHandler:         listHandler,
		CreateHandler:       createHandler,
		PersonChangeHandler: personChangeHandler,
	}
}

func Run(srv *DepartmentManager) {
	listen, err := net.Listen("tcp", config.DepartmentGrpcTransportAddr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	grpcServer := NewGrpcServer(srv)
	common.RegisterDepartmentManagerServiceServer(server, grpcServer)

	log.Fatal(server.Serve(listen))
}
