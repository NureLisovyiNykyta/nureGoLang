package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"school-journal/database"
	"school-journal/models"
)

// ========================
//   STUDENT STATISTICS
// ========================

// @Summary Get average grade of a student by subject
// @Description Returns the average grade for a given student and subject
// @Tags Statistics
// @Produce json
// @Param student_id path string true "Student ID (UUID)"
// @Param subject_id path string true "Subject ID (UUID)"
// @Success 200 {object} map[string]float64
// @Failure 404 {object} map[string]string
// @Router /api/statistics/student/{student_id}/subject/{subject_id}/average [get]
func GetStudentAverageBySubject(c *gin.Context) {
	var avg float64
	studentID := c.Param("student_id")
	subjectID := c.Param("subject_id")

	result := database.DB.
		Model(&models.Grade{}).
		Joins("JOIN lessons ON grades.lesson_id = lessons.id").
		Where("grades.student_id = ? AND lessons.subject_id = ?", studentID, subjectID).
		Select("AVG(grades.grade)").Scan(&avg)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no grades found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"average_grade": avg})
}

// @Summary Get average grade of a class by subject
// @Description Returns the average grade for all students in a class by subject
// @Tags Statistics
// @Produce json
// @Param class_id path string true "Class ID (UUID)"
// @Param subject_id path string true "Subject ID (UUID)"
// @Success 200 {object} map[string]float64
// @Failure 404 {object} map[string]string
// @Router /api/statistics/class/{class_id}/subject/{subject_id}/average [get]
func GetClassAverageBySubject(c *gin.Context) {
	var avg float64
	classID := c.Param("class_id")
	subjectID := c.Param("subject_id")

	result := database.DB.
		Model(&models.Grade{}).
		Joins("JOIN students ON grades.student_id = students.id").
		Joins("JOIN lessons ON grades.lesson_id = lessons.id").
		Where("students.class_id = ? AND lessons.subject_id = ?", classID, subjectID).
		Select("AVG(grades.grade)").Scan(&avg)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no grades found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"average_grade": avg})
}

// ========================
//   RELATED DATA QUERIES
// ========================

// @Summary Get all students in a class
// @Description Returns all students belonging to a given class
// @Tags Classes
// @Produce json
// @Param class_id path string true "Class ID (UUID)"
// @Success 200 {array} models.Student
// @Router /api/classes/{class_id}/students [get]
func GetStudentsByClass(c *gin.Context) {
	classID := c.Param("class_id")
	var students []models.Student

	if err := database.DB.Where("class_id = ?", classID).Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	c.JSON(http.StatusOK, students)
}

// @Summary Get all lessons of a class grouped by subject
// @Description Returns lessons of a class grouped by subject name
// @Tags Classes
// @Produce json
// @Param class_id path string true "Class ID (UUID)"
// @Success 200 {array} map[string]interface{}
// @Router /api/classes/{class_id}/lessons [get]
func GetLessonsByClass(c *gin.Context) {
	classID := c.Param("class_id")
	var lessons []struct {
		SubjectName string
		LessonID    string
		Date        string
	}

	query := `
		SELECT subjects.name AS subject_name, lessons.id AS lesson_id, lessons.date
		FROM lessons
		JOIN subjects ON lessons.subject_id = subjects.id
		WHERE lessons.class_id = ?
		ORDER BY subjects.name, lessons.date;
	`

	if err := database.DB.Raw(query, classID).Scan(&lessons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

// @Summary Get all classes of a teacher
// @Description Returns list of unique classes taught by a teacher
// @Tags Teachers
// @Produce json
// @Param teacher_id path string true "Teacher ID (UUID)"
// @Success 200 {array} models.Class
// @Router /api/teachers/{teacher_id}/classes [get]
func GetClassesByTeacher(c *gin.Context) {
	teacherID := c.Param("teacher_id")
	var classes []models.Class

	err := database.DB.
		Table("classes").
		Joins("JOIN lessons ON lessons.class_id = classes.id").
		Where("lessons.teacher_id = ?", teacherID).
		Group("classes.id").
		Find(&classes).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}

	c.JSON(http.StatusOK, classes)
}

// @Summary Get all grades of a student
// @Description Returns all grades of a student across all subjects
// @Tags Students
// @Produce json
// @Param student_id path string true "Student ID (UUID)"
// @Success 200 {array} models.Grade
// @Router /api/students/{student_id}/grades [get]
func GetGradesByStudent(c *gin.Context) {
	studentID := c.Param("student_id")
	var grades []models.Grade

	if err := database.DB.Preload("Lesson").Where("student_id = ?", studentID).Find(&grades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	c.JSON(http.StatusOK, grades)
}
