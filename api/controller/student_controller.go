package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type StudentController struct{
	StudentUsecase domain.StudentUsecase
}

func (sc *StudentController)Create(c *gin.Context){
	var createStudentDto domain.CreateStudentDto

	err := c.ShouldBindJSON(&createStudentDto)
	if err != nil{


		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{
			field := e.Field()
			switch field {
			case "StudentRegularinfo":
				errorMessage["StudentRegularinfo"] = "Student Regular Info is required"
			case "StudentId":
				errorMessage["StudentId"] = "Student ID is required"
			case "FirstName":
				errorMessage["FirstName"] = "First Name is required"
			case "FatherName":
				errorMessage["FatherName"] = "Father Name is required"
			case "GrandFatherName":
				errorMessage["GrandFatherName"] = "Grandfather Name is required"
			case "YearOfRegistration":
				errorMessage["YearOfRegistration"] = "Year of Registration is required"
			case "PhoneNumber":
				errorMessage["PhoneNumber"] = "Phone Number is required"
			case "Religion":
				errorMessage["Religion"] = "Religion is required"
			case "Sex":
				errorMessage["Sex"] = "Sex is required"
			case "Status":
				errorMessage["Status"] = "Status is required"
			case "RegisteredBy":
				errorMessage["RegisteredBy"] = "Registered By is required"
			case "StudentAdditionalinfo":
				errorMessage["StudentRegularinfo"] = "Student Regular Info is required"
			case "CardNumber":
				errorMessage["CardNumber"] = "Card Number is required"
			case "Photo":
				errorMessage["Photo"] = "Photo is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation Failed ", Errors: errorMessage})
		return 
	}


	//less verbose but not convenient for client dev.
	// validate := validator.New()	// // Validate the User struct
	// err = validate.Struct(createStudentDto)
	// if err != nil {
	// 	// Validation failed, handle the error
	// 	errors := err.(validator.ValidationErrors)
	// 	errorMessages := make(map[string]string)
	// 	for _, err := range errors {
	// 		errorMessages[err.Field()] = err.Tag()
	// 	}
	// 	c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation Failed ", Errors: errorMessages})
	// 	return 

	// }
		
	createStudentDto.StudentId = utils.ConvertToSmallLetter(createStudentDto.StudentId)
	err = sc.StudentUsecase.Create(c, &createStudentDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, domain.SucessResponse{Message: "created successfully."})
}


func (sc *StudentController)Update(c *gin.Context){
	var updateStudentDto domain.UpdateStudentDto

	err := c.ShouldBindJSON(&updateStudentDto)
	if err != nil{
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{
			field := e.Field()
			switch field {
			case "StudentInformationId":
				errorMessage["StudentInformationId"] = "Student Information Id is required"
			case "StudentRegularinfo":
				errorMessage["StudentRegularinfo"] = "Student Regular Info is required"
			case "StudentId":
				errorMessage["StudentId"] = "Student ID is required"
			case "FirstName":
				errorMessage["FirstName"] = "First Name is required"
			case "FatherName":
				errorMessage["FatherName"] = "Father Name is required"
			case "GrandFatherName":
				errorMessage["GrandFatherName"] = "Grandfather Name is required"
			case "YearOfRegistration":
				errorMessage["YearOfRegistration"] = "Year of Registration is required"
			case "PhoneNumber":
				errorMessage["PhoneNumber"] = "Phone Number is required"
			case "Religion":
				errorMessage["Religion"] = "Religion is required"
			case "Sex":
				errorMessage["Sex"] = "Sex is required"
			case "Status":
				errorMessage["Status"] = "Status is required"
			case "RegisteredBy":
				errorMessage["RegisteredBy"] = "Registered By is required"
			case "StudentAdditionalinfo":
				errorMessage["StudentRegularinfo"] = "Student Regular Info is required"
			case "CardNumber":
				errorMessage["CardNumber"] = "Card Number is required"
			case "Photo":
				errorMessage["Photo"] = "Photo is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation Failed ", Errors: errorMessage})
		return 
		}

		updateStudentDto.StudentId = utils.ConvertToSmallLetter(updateStudentDto.StudentId)
		err = sc.StudentUsecase.Update(c, &updateStudentDto)
		if err != nil{
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return 
		}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "updated successfully."})

}

func (sc *StudentController)FetchStudents(c *gin.Context){

	students, err := sc.StudentUsecase.FetchStudents(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK,  students)
}



func (sc *StudentController)Delete(c *gin.Context){

	StudentInformationId := c.Param("studentInformationId")
	bit32, err := utils.ConverParamID(StudentInformationId)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = sc.StudentUsecase.Delete(c, *bit32)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "deleted successfully."})
}


func (sc *StudentController)FetchByStudentId(c *gin.Context){
	studentId := c.Param("studentId")
	studentId = utils.ConvertToSmallLetter(studentId)
	
	student, err := sc.StudentUsecase.FetchByStudentId(c, studentId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK,  student)
}



