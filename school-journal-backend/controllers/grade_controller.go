package controllers

import (
	"net/http"
	"school-journal/database"
	"school-journal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateGrade godoc
// @Summary Create a new grade
// @Description Create a new grade record
// @Tags grades
// @Accept json
// @Produce json
// @Param grade body models.Grade true "Grade to create"
// @Success 201 {object} models.Grade
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /grades [post]
func CreateGrade(c *gin.Context) {
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&grade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, grade)
}

// GetGrades godoc
// @Summary List grades
// @Description Get list of all grades
// @Tags grades
// @Accept json
// @Produce json
// @Success 200 {array} models.Grade
// @Failure 500 {object} map[string]string
// @Router /grades [get]
func GetGrades(c *gin.Context) {
	var grades []models.Grade
	if err := database.DB.Preload("Student").Preload("Lesson").Find(&grades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grades)
}

// GetGradeByID godoc
// @Summary Get a grade by id
// @Description Get grade details by UUID
// @Tags grades
// @Accept json
// @Produce json
// @Param id path string true "Grade ID"
// @Success 200 {object} models.Grade
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /grades/{id} [get]
func GetGradeByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var grade models.Grade
	if err := database.DB.First(&grade, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
		return
	}
	c.JSON(http.StatusOK, grade)
}

// UpdateGrade godoc
// @Summary Update a grade
// @Description Update grade by UUID
// @Tags grades
// @Accept json
// @Produce json
// @Param id path string true "Grade ID"
// @Param grade body models.Grade true "Grade data to update"
// @Success 200 {object} models.Grade
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /grades/{id} [put]
func UpdateGrade(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var grade models.Grade
	if err := database.DB.First(&grade, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
		return
	}
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&grade)
	c.JSON(http.StatusOK, grade)
}

// DeleteGrade godoc
// @Summary Delete a grade
// @Description Delete grade by UUID
// @Tags grades
// @Accept json
// @Produce json
// @Param id path string true "Grade ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /grades/{id} [delete]
func DeleteGrade(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	database.DB.Delete(&models.Grade{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Grade deleted"})
}
