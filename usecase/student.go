package usecase

import (
	"context"
	"time"

	"github.com/mulukenhailu/Binary/domain"
)

type studentUsecase struct {
	studentRepository domain.StudentRepository
	contextTimeout    time.Duration
}


func NewStudentUsecase(studentRepository domain.StudentRepository, contextTimeout time.Duration) domain.StudentUsecase {
	return &studentUsecase{
		studentRepository: studentRepository,
		contextTimeout:    contextTimeout,
	}
}

// Create implements domain.StudentUsecase.
func (su *studentUsecase) Create(c context.Context, createStudentDto *domain.CreateStudentDto) error {
	return su.studentRepository.Create(c, createStudentDto)
}

// Delete implements domain.StudentUsecase.
func (su *studentUsecase) Delete(c context.Context, studentInformationId int32) error {
	return su.studentRepository.Delete(c, studentInformationId)
}

// FetchByStudentId implements domain.StudentUsecase.
func (su *studentUsecase) FetchByStudentId(c context.Context, studentId string) (domain.Student, error) {
	return su.studentRepository.FetchByStudentId(c, studentId)
}

// FetchStudents implements domain.StudentUsecase.
func (su *studentUsecase) FetchStudents(c context.Context) ([]domain.Student, error) {
	return su.studentRepository.FetchStudents(c)
}

// Update implements domain.StudentUsecase.
func (su *studentUsecase) Update(c context.Context, updateStudentDto *domain.UpdateStudentDto) error {
	return su.studentRepository.Update(c, updateStudentDto)
}

