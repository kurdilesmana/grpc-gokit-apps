package endpoints

import (
	"context"
	"grpc-gokit-apps/service"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	Add      endpoint.Endpoint
	Multiply endpoint.Endpoint
}

// MathReq struct holds the endpoint request definition
type MathReq struct {
	Num_a float32
	Num_b float32
}

// MathResp struct holds the endpoint response definition
type MathResp struct {
	Result float32
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Add:      makeAddEndpoint(s),
		Multiply: makeMultiplyEndpoint(s),
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Add(ctx, req.Num_a, req.Num_b)
		return MathResp{Result: result}, nil
	}
}

func makeMultiplyEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Multiply(ctx, req.Num_a, req.Num_b)
		return MathResp{Result: result}, nil
	}
}
