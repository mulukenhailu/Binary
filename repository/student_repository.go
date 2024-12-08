package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type studentRepository struct {
	pg *database.Queries
}

func NewStudentRepository(db *pgxpool.Pool) domain.StudentRepository {
	return &studentRepository{
		pg: database.New(db),
	}
}


// Create implements domain.StudentRepository.
func (sr *studentRepository) Create(c context.Context, createStudentDto *domain.CreateStudentDto) error {
	studentParams := database.CreateStudentParams{
		Studentid          :createStudentDto.StudentId,
		Firstname          :createStudentDto.FirstName,
		Fathername         :createStudentDto.FatherName,
		Grandfathername    :createStudentDto.GrandFatherName,
		Yearofregistration :createStudentDto.YearOfRegistration,
		Phonenumber        :createStudentDto.PhoneNumber,
		Religion           :createStudentDto.Religion,
		Sex                :createStudentDto.Sex,
		Status             :createStudentDto.Status,
		Cardnumber         :pgtype.Text{String: createStudentDto.CardNumber, Valid: createStudentDto.CardNumber != ""},
		Photo              :pgtype.Text{String: createStudentDto.Photo, 	 Valid: createStudentDto.CardNumber != ""},
		Registeredby       :createStudentDto.RegisteredBy,
	}
	_, err := sr.pg.CreateStudent(c, studentParams)
	return err
}


// Update implements domain.StudentRepository.
func (sr *studentRepository) Update(c context.Context, updateStudentDto *domain.UpdateStudentDto) error {

	studentUpdateParam := database.UpdateStudentParams{
		Studentinformationid :updateStudentDto.StudentInformationId,
		Studentid            :updateStudentDto.StudentId,
		Firstname            :updateStudentDto.FirstName,
		Fathername           :updateStudentDto.FatherName,
		Grandfathername      :updateStudentDto.GrandFatherName,
		Yearofregistration   :updateStudentDto.YearOfRegistration,
		Phonenumber          :updateStudentDto.PhoneNumber,
		Religion             :updateStudentDto.Religion,
		Sex                  :updateStudentDto.Sex,
		Status               :updateStudentDto.Status,
		Cardnumber           :pgtype.Text{String: updateStudentDto.CardNumber, 	Valid: updateStudentDto.CardNumber != ""},
		Photo              	 :pgtype.Text{String: updateStudentDto.Photo, 	 	Valid: updateStudentDto.CardNumber != ""},
		Registeredby         :updateStudentDto.RegisteredBy,
	}
	_, err := sr.pg.UpdateStudent(c, studentUpdateParam)
	return err
}


// Delete implements domain.StudentRepository.
func (sr *studentRepository) Delete(c context.Context, deleteStudentDto *domain.DeleteStudentDto) error {
	_, err := sr.pg.DeleteStudent(c, deleteStudentDto.StudentInformationId)
	return err
}

// FetchStudents implements domain.StudentRepository.
func (sr *studentRepository) FetchStudents(c context.Context) ([]domain.Student, error) {
	dbStudents, err := sr.pg.FetchStudents(c)
	domainStudents := utils.ConvertDbStudentsToDomainStudents(dbStudents)
	return domainStudents, err
}


// FetchByStudentId implements domain.StudentRepository.
func (sr *studentRepository) FetchByStudentId(c context.Context, studentId string) (domain.Student, error) {
	dbStudent, err := sr.pg.FetchByStudentId(c, studentId)
	domainStudent := utils.ConvertDbStudentsToDomainStudents([]database.Student{dbStudent})
	return domainStudent[0], err
}


