package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
}

// Service interface describes a service that adds numbers
type Service interface {
	Add(ctx context.Context, num_a, num_b float32) (float32, error)
	Multiply(ctx context.Context, num_a, num_b float32) (float32, error)
}

// NewService returns a Service with all of the expected dependencies
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) Add(ctx context.Context, num_a, num_b float32) (float32, error) {
	return num_a + num_b, nil
}

func (s *service) Multiply(ctx context.Context, num_a, num_b float32) (float32, error) {
	return num_a * num_b, nil
}
