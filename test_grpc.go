package main

import (
	"context"
	"fmt"
	"net"

	"github.com/ginanjar-template-golang/shared-pkg/interceptor"
	authpb "github.com/ginanjar-template-golang/shared-pkg/proto/pb/authpb"
	responsepb "github.com/ginanjar-template-golang/shared-pkg/proto/pb/responsepb"
	grpcResponse "github.com/ginanjar-template-golang/shared-pkg/response/grpc_response"
	"google.golang.org/grpc"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
}

func (h *AuthHandler) GetUsers(ctx context.Context, req *authpb.PaginationRequest) (*responsepb.PaginationResponse, error) {
	fmt.Println(req.Page)
	// if req.Page <= 0 {
	// 	// Contoh error validasi
	// 	err := appError.CreateResourceError("invalidPage", map[string]any{
	// 		"field": "page",
	// 		"value": req.Page,
	// 	})
	// 	return nil, grpcResponse.FromAppError(ctx, err)
	// }

	// Simulasi data sukses
	data := []map[string]any{
		{"id": 1, "name": "Alice"},
		{"id": 2, "name": "Bob"},
	}

	pagination := grpcResponse.PaginationData{
		Page:     req.Page,
		Size:     req.Limit,
		Limit:    req.Limit,
		TotalRow: 10,
		Results:  data,
	}

	return grpcResponse.PaginationSuccess(ctx, "getUsersSuccess", pagination)

}

// grpc test
func TestGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryRecovery(),
			interceptor.UnaryRequestLogger(),
		),
	)

	authpb.RegisterAuthServiceServer(s, &AuthHandler{})

	fmt.Println("gRPC server running at :50051")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
