package transports

import (
	"context"
	"grpc-gokit-apps/endpoints"
	"grpc-gokit-apps/pb"

	"github.com/go-kit/kit/log"

	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	add gt.Handler
	// Embed the unimplemented server
	pb.UnimplementedMathServiceServer
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.MathServiceServer {
	return &gRPCServer{
		add: gt.NewServer(
			endpoints.Add,
			decodeMathRequest,
			encodeMathResponse,
		),
	}
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return endpoints.MathReq{Num_a: req.NumA, Num_b: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResp)
	return &pb.MathResponse{Result: resp.Result}, nil
}
