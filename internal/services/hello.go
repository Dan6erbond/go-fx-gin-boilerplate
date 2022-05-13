package services

import (
	"fmt"

	"github.com/Dan6erbond/go-fx-gin-template/internal/dto"
	"go.uber.org/zap"
)

type HelloWorldService struct {
	logger *zap.Logger
}

func NewHelloWorldService(logger *zap.Logger) *HelloWorldService {
	return &HelloWorldService{
		logger: logger,
	}
}

func (s *HelloWorldService) GetHelloWorld(name string) (*dto.HelloWorldResponse, error) {
	if name == "" {
		return &dto.HelloWorldResponse{
			Message: "Hello world!",
		}, nil
	}
	return &dto.HelloWorldResponse{
		Message: fmt.Sprintf("Hello %s!", name),
	}, nil
}
