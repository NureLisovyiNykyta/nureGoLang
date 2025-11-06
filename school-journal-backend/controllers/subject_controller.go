package controllers

import (
	"net/http"
	"school-journal/database"
	"school-journal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CREATE
// CreateSubject godoc
// @Summary Create a new subject
// @Description Create a new subject record
// @Tags subjects
// @Accept json
// @Produce json
// @Param subject body models.Subject true "Subject to create"
// @Success 201 {object} models.Subject
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subjects [post]
func CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&subject).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, subject)
}

// READ all
// GetSubjects godoc
// @Summary List subjects
// @Description Get list of all subjects
// @Tags subjects
// @Accept json
// @Produce json
// @Success 200 {array} models.Subject
// @Failure 500 {object} map[string]string
// @Router /subjects [get]
func GetSubjects(c *gin.Context) {
	var subjects []models.Subject
	if err := database.DB.Find(&subjects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}

// READ by id
// GetSubjectByID godoc
// @Summary Get a subject by id
// @Description Get subject details by UUID
// @Tags subjects
// @Accept json
// @Produce json
// @Param id path string true "Subject ID"
// @Success 200 {object} models.Subject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /subjects/{id} [get]
func GetSubjectByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var subject models.Subject
	if err := database.DB.First(&subject, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}
	c.JSON(http.StatusOK, subject)
}

// UPDATE
// UpdateSubject godoc
// @Summary Update a subject
// @Description Update subject by UUID
// @Tags subjects
// @Accept json
// @Produce json
// @Param id path string true "Subject ID"
// @Param subject body models.Subject true "Subject data to update"
// @Success 200 {object} models.Subject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subjects/{id} [put]
func UpdateSubject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var subject models.Subject
	if err := database.DB.First(&subject, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&subject)
	c.JSON(http.StatusOK, subject)
}

// DELETE
// DeleteSubject godoc
// @Summary Delete a subject
// @Description Delete subject by UUID
// @Tags subjects
// @Accept json
// @Produce json
// @Param id path string true "Subject ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subjects/{id} [delete]
func DeleteSubject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	database.DB.Delete(&models.Subject{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted"})
}
