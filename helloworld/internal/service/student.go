package service

import (
	"context"
	"helloworld/internal/biz"
	"helloworld/internal/model"

	pb "helloworld/api/student/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServer
	uc *biz.StudentUsecase
}

func NewStudentService(uc *biz.StudentUsecase) *StudentService {
	return &StudentService{
		uc: uc,
	}
}

func (s *StudentService) GetStudent(ctx context.Context, req *pb.StudentReq) (*pb.UserReply, error) {
	student, err := s.uc.GetStudentById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{
		Code:    200,
		Message: "student",
		Data:    student,
	}, nil
}
func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CrStuReq) (*pb.Resp, error) {
	err := s.uc.CreateStudent(ctx, &model.Student{
		Name:  req.Name,
		Age:   req.Age,
		Total: req.Total,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Resp{
		Message: "success",
	}, nil
}
