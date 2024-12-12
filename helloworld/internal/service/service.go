package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewService, NewStudentService)

type Service struct {
	pb.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

func NewService(uc *biz.GreeterUsecase, logger log.Logger) *Service {
	return &Service{uc: uc}
}
