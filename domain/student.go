package domain

import "context"

type Student struct {
  StudentInformationId  int32                `json:"student_information_id" binding:"required"`
	StudentRegularData    *StudentRegularInfo  `json:"Student"  binding:"required"`
	StudentAdditionInfo   *StudentAdditionInfo `json:"sudent_addition_info"`
  RegisteredDate        string               `json:"registered_date" binding:"required"`
}

type StudentRegularInfo struct {
	StudentId             string `json:"student_id" binding:"required"`
	FirstName             string `json:"first_name" binding:"required"`
	FatherName            string `json:"father_name" binding:"required"`
	GrandFatherName       string `json:"grand_father_name" binding:"required"`
	YearOfRegistration    string `json:"year_of_registration" binding:"required"`
	PhoneNumber           string `json:"phone_number" binding:"required"`
	Religion              string `json:"religion" binding:"required"`
	Sex                   string `json:"sex" binding:"required"`
	Status                string `json:"status" binding:"required"`
  RegisteredBy          string `json:"registered_by" binding:"required"`
}

type StudentAdditionInfo struct {
	CardNumber  string `json:"card_number" binding:"required"`
	Photo       string `json:"photo" binding:"required"`
}

type CreateStudentDto struct {
	StudentRegularInfo  *StudentRegularInfo   `json:"Student"  binding:"required"`
	StudentAdditionInfo *StudentAdditionInfo  `json:"sudent_addition_info"`
}

type UpdateStudentDto struct {
  StudentInformationId    int32                 `json:"student_information_id" binding:"required"`
  StudentRegularInfo      *StudentRegularInfo   `json:"Student"  binding:"required"`
	StudentAdditionInfo     *StudentAdditionInfo  `json:"sudent_addition_info"`
}

type DeleteStudentDto struct{
  StudentInformationId  int32 `json:"student_information_id" binding:"required"`
}

type FetchByStudentId struct {
  StudentId string  `json:"student_id" binding:"required"`
}

type StudentRepository interface {
	Create(c context.Context, createStudentDto *CreateStudentDto)           error
	Update(c context.Context, updateStudentDto *UpdateStudentDto)           error
	Delete(c context.Context, deleteStudentDto *DeleteStudentDto)           error
	FetchStudents(c context.Context)                                        ([]Student, error)
	FetchByStudentId(c context.Context, createStudentDto *FetchByStudentId) (Student, error)
}



