package controllers

import (
	"net/http"
	"school-journal/database"
	"school-journal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateClass godoc
// @Summary Create a new class
// @Description Create a new class record
// @Tags classes
// @Accept json
// @Produce json
// @Param class body models.Class true "Class to create"
// @Success 201 {object} models.Class
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /classes [post]
func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, class)
}

// GetClasses godoc
// @Summary List classes
// @Description Get list of all classes
// @Tags classes
// @Accept json
// @Produce json
// @Success 200 {array} models.Class
// @Failure 500 {object} map[string]string
// @Router /classes [get]
func GetClasses(c *gin.Context) {
	var classes []models.Class
	if err := database.DB.Preload("Teacher").Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}

// GetClassByID godoc
// @Summary Get a class by id
// @Description Get class details by UUID
// @Tags classes
// @Accept json
// @Produce json
// @Param id path string true "Class ID"
// @Success 200 {object} models.Class
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /classes/{id} [get]
func GetClassByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var class models.Class
	if err := database.DB.Preload("Teacher").First(&class, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	c.JSON(http.StatusOK, class)
}

// UpdateClass godoc
// @Summary Update a class
// @Description Update class by UUID
// @Tags classes
// @Accept json
// @Produce json
// @Param id path string true "Class ID"
// @Param class body models.Class true "Class data to update"
// @Success 200 {object} models.Class
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /classes/{id} [put]
func UpdateClass(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var class models.Class
	if err := database.DB.First(&class, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	var input models.Class
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Model(&class).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}

// DeleteClass godoc
// @Summary Delete a class
// @Description Delete class by UUID
// @Tags classes
// @Accept json
// @Produce json
// @Param id path string true "Class ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /classes/{id} [delete]
func DeleteClass(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	if err := database.DB.Delete(&models.Class{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted"})
}