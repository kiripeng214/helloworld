package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper("data/student", logger),
	}
}

func (s studentRepo) ListStudent(ctx context.Context) ([]*biz.Student, error) {
	panic("implement me")
}
