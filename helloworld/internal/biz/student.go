package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "helloworld/api/student/v1"
	"helloworld/internal/model"
)

type StudentRepo interface {
	GetStudentById(ctx context.Context, id int64) (*model.Student, error)
	CreateStudent(ctx context.Context, student *model.Student) error
}

type StudentUsecase struct {
	repo   StudentRepo
	logger *log.Helper
}

func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsecase {
	return &StudentUsecase{
		repo:   repo,
		logger: log.NewHelper(logger),
	}
}
func (uc *StudentUsecase) GetStudentById(ctx context.Context, id int64) (*v1.UserReply_Data, error) {
	user, _ := uc.repo.GetStudentById(ctx, id)
	return &v1.UserReply_Data{
		Id:    user.Id,
		Name:  user.Name,
		Age:   user.Age,
		Total: user.Total,
	}, nil
}
func (uc *StudentUsecase) CreateStudent(ctx context.Context, student *model.Student) error {
	return uc.repo.CreateStudent(ctx, student)
}
