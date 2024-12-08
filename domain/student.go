package domain

import (
	"context"
)


type Student struct {
  	StudentInformationId  	int32                	`json:"student_information_id"`
	StudentRegularinfo    	*StudentRegularInfo  	`json:"Student_regular_info"`
	StudentAdditionalInfo   *StudentAdditionalInfo	`json:"student_additional_info"`
  	RegisteredDate        	string               	`json:"registered_date"`
}

type StudentRegularInfo struct {
	StudentId             string `json:"student_id" binding:"required" validate:"required"`
	FirstName             string `json:"first_name" binding:"required" validate:"required"`
	FatherName            string `json:"father_name" binding:"required" validate:"required"`
	GrandFatherName       string `json:"grand_father_name" binding:"required" validate:"required"`
	YearOfRegistration    string `json:"year_of_registration" binding:"required" validate:"required"`
	PhoneNumber           string `json:"phone_number" binding:"required" validate:"required"`
	Religion              string `json:"religion" binding:"required" validate:"required"`
	Sex                   string `json:"sex" binding:"required" validate:"required"`
	Status                string `json:"status" binding:"required" validate:"required"`
  	RegisteredBy          string `json:"registered_by" binding:"required" validate:"required"`
}

type StudentAdditionalInfo struct {
	CardNumber  string `json:"card_number" binding:"required" validate:"required,omitempty"`
	Photo       string `json:"photo" binding:"required" validate:"required,omitempty"`
}

type CreateStudentDto struct {
	StudentId             string `json:"student_id" binding:"required" validate:"required"`
	FirstName             string `json:"first_name" binding:"required" validate:"required"`
	FatherName            string `json:"father_name" binding:"required" validate:"required"`
	GrandFatherName       string `json:"grand_father_name" binding:"required" validate:"required"`
	YearOfRegistration    string `json:"year_of_registration" binding:"required" validate:"required"`
	PhoneNumber           string `json:"phone_number" binding:"required" validate:"required"`
	Religion              string `json:"religion" binding:"required" validate:"required"`
	Sex                   string `json:"sex" binding:"required" validate:"required"`
	Status                string `json:"status" binding:"required" validate:"required"`
  	RegisteredBy          string `json:"registered_by" binding:"required" validate:"required"`
	CardNumber  		  string `json:"card_number"`
	Photo       		  string `json:"photo"`
}

type UpdateStudentDto struct {
  	StudentInformationId  int32  `json:"student_information_id" binding:"required"`
  	StudentId             string `json:"student_id" binding:"required" validate:"required"`
	FirstName             string `json:"first_name" binding:"required" validate:"required"`
	FatherName            string `json:"father_name" binding:"required" validate:"required"`
	GrandFatherName       string `json:"grand_father_name" binding:"required" validate:"required"`
	YearOfRegistration    string `json:"year_of_registration" binding:"required" validate:"required"`
	PhoneNumber           string `json:"phone_number" binding:"required" validate:"required"`
	Religion              string `json:"religion" binding:"required" validate:"required"`
	Sex                   string `json:"sex" binding:"required" validate:"required"`
	Status                string `json:"status" binding:"required" validate:"required"`
	RegisteredBy          string `json:"registered_by" binding:"required" validate:"required"`
	CardNumber  		  string `json:"card_number"`
	Photo       		  string `json:"photo"`
}

type DeleteStudentDto struct{
  StudentInformationId  int32 `json:"student_information_id" binding:"required"`
}

type FetchByStudentIdDto struct {
  StudentId string  `json:"student_id" binding:"required"`
}

type StudentRepository interface {
	Create(c context.Context, createStudentDto *CreateStudentDto)           error
	Update(c context.Context, updateStudentDto *UpdateStudentDto)           error
	Delete(c context.Context, deleteStudentDto *DeleteStudentDto)           error
	FetchStudents(c context.Context)                                        ([]Student, error)
	FetchByStudentId(c context.Context, studentId string) 					(Student, error)
}

type StudentUsecase interface {
	Create(c context.Context, createStudentDto *CreateStudentDto)           error
	Update(c context.Context, updateStudentDto *UpdateStudentDto)           error
	Delete(c context.Context, deleteStudentDto *DeleteStudentDto)           error
	FetchStudents(c context.Context)                                        ([]Student, error)
	FetchByStudentId(c context.Context, studentId string) 					(Student, error)
}



