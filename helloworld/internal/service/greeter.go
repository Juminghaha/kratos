package service

import (
	"context"

	pb "helloworld/api/helloworld/v1"
)

func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *Service) SayGoodbye(ctx context.Context, req *pb.GoodbyeRequest) (*pb.GoodbyeReply, error) {
	//goodbye, err := s.uc.ByeBye(ctx, req.Name)
	//if err != nil {
	//	return nil, err
	//}

	return &pb.GoodbyeReply{Message: "Goodbye, " + req.Name + "!"}, nil
}
