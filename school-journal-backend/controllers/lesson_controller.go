package controllers

import (
	"net/http"
	"school-journal/database"
	"school-journal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateLesson godoc
// @Summary Create a new lesson
// @Description Create a new lesson record
// @Tags lessons
// @Accept json
// @Produce json
// @Param lesson body models.Lesson true "Lesson to create"
// @Success 201 {object} models.Lesson
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons [post]
func CreateLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&lesson).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, lesson)
}

// GetLessons godoc
// @Summary List lessons
// @Description Get list of all lessons
// @Tags lessons
// @Accept json
// @Produce json
// @Success 200 {array} models.Lesson
// @Failure 500 {object} map[string]string
// @Router /lessons [get]
func GetLessons(c *gin.Context) {
	var lessons []models.Lesson
	if err := database.DB.Preload("Class").Preload("Subject").Preload("Teacher").Find(&lessons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

// GetLessonByID godoc
// @Summary Get a lesson by id
// @Description Get lesson details by UUID
// @Tags lessons
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Success 200 {object} models.Lesson
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /lessons/{id} [get]
func GetLessonByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var lesson models.Lesson
	if err := database.DB.Preload("Class").Preload("Subject").Preload("Teacher").First(&lesson, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

// UpdateLesson godoc
// @Summary Update a lesson
// @Description Update lesson by UUID
// @Tags lessons
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Param lesson body models.Lesson true "Lesson data to update"
// @Success 200 {object} models.Lesson
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons/{id} [put]
func UpdateLesson(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var lesson models.Lesson
	if err := database.DB.First(&lesson, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&lesson)
	c.JSON(http.StatusOK, lesson)
}

// DeleteLesson godoc
// @Summary Delete a lesson
// @Description Delete lesson by UUID
// @Tags lessons
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons/{id} [delete]
func DeleteLesson(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	database.DB.Delete(&models.Lesson{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted"})
}
