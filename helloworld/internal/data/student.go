package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/model"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

func (s *studentRepo) GetStudentById(ctx context.Context, id int64) (*model.Student, error) {
	var student model.Student
	if err := s.data.pdb.First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *studentRepo) CreateStudent(ctx context.Context, student *model.Student) error {
	if err := s.data.pdb.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
