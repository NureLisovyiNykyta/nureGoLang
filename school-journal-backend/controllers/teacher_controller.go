package controllers

import (
	"net/http"
	"school-journal/database"
	"school-journal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CREATE
// CreateTeacher godoc
// @Summary Create a new teacher
// @Description Create a new teacher record
// @Tags teachers
// @Accept json
// @Produce json
// @Param teacher body models.Teacher true "Teacher to create"
// @Success 201 {object} models.Teacher
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers [post]
func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&teacher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, teacher)
}

// READ all
// GetTeachers godoc
// @Summary List teachers
// @Description Get list of all teachers
// @Tags teachers
// @Accept json
// @Produce json
// @Success 200 {array} models.Teacher
// @Failure 500 {object} map[string]string
// @Router /teachers [get]
func GetTeachers(c *gin.Context) {
	var teachers []models.Teacher
	if err := database.DB.Preload("Subject").Find(&teachers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teachers)
}

// READ by id
// GetTeacherByID godoc
// @Summary Get a teacher by id
// @Description Get teacher details by UUID
// @Tags teachers
// @Accept json
// @Produce json
// @Param id path string true "Teacher ID"
// @Success 200 {object} models.Teacher
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /teachers/{id} [get]
func GetTeacherByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var teacher models.Teacher
	if err := database.DB.Preload("Subject").First(&teacher, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

// UPDATE
// UpdateTeacher godoc
// @Summary Update a teacher
// @Description Update teacher by UUID
// @Tags teachers
// @Accept json
// @Produce json
// @Param id path string true "Teacher ID"
// @Param teacher body models.Teacher true "Teacher data to update"
// @Success 200 {object} models.Teacher
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers/{id} [put]
func UpdateTeacher(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var teacher models.Teacher
	if err := database.DB.First(&teacher, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&teacher)
	c.JSON(http.StatusOK, teacher)
}

// DELETE
// DeleteTeacher godoc
// @Summary Delete a teacher
// @Description Delete teacher by UUID
// @Tags teachers
// @Accept json
// @Produce json
// @Param id path string true "Teacher ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers/{id} [delete]
func DeleteTeacher(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	database.DB.Delete(&models.Teacher{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted"})
}
