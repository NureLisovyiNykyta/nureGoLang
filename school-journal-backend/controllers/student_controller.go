package controllers

import (
	"net/http"
	"school-journal/database"
	"school-journal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

// CreateStudent godoc
// @Summary Create a new student
// @Description Create a new student record
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student to create"
// @Success 201 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students [post]
func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	if err := validate.Struct(student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error: " + err.Error()})
		return
	}

	if err := database.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student created successfully", "student": student})
}

// GetStudents godoc
// @Summary List students
// @Description Get list of all students
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {array} models.Student
// @Failure 500 {object} map[string]string
// @Router /students [get]
func GetStudents(c *gin.Context) {
	var students []models.Student
	if err := database.DB.Preload("Class").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

// GetStudentByID godoc
// @Summary Get a student by id
// @Description Get student details by UUID
// @Tags students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /students/{id} [get]
func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var student models.Student
	if err := database.DB.Preload("Class").First(&student, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// UpdateStudent godoc
// @Summary Update a student
// @Description Update student by UUID
// @Tags students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Param student body models.Student true "Student data to update"
// @Success 200 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [put]
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var student models.Student
	if err := database.DB.First(&student, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	if err := database.DB.Model(&student).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully", "student": student})
}

// DeleteStudent godoc
// @Summary Delete a student
// @Description Delete student by UUID
// @Tags students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	if err := database.DB.Delete(&models.Student{}, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
