package routes

import (
	"github.com/gin-gonic/gin"
	"school-journal/controllers"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Students
	api.POST("/students", controllers.CreateStudent)
	api.GET("/students", controllers.GetStudents)
	api.GET("/students/:id", controllers.GetStudentByID)
	api.PUT("/students/:id", controllers.UpdateStudent)
	api.DELETE("/students/:id", controllers.DeleteStudent)

	// Teachers
	api.POST("/teachers", controllers.CreateTeacher)
	api.GET("/teachers", controllers.GetTeachers)
	api.GET("/teachers/:id", controllers.GetTeacherByID)
	api.PUT("/teachers/:id", controllers.UpdateTeacher)
	api.DELETE("/teachers/:id", controllers.DeleteTeacher)

	// Classes
	api.POST("/classes", controllers.CreateClass)
	api.GET("/classes", controllers.GetClasses)
	api.GET("/classes/:id", controllers.GetClassByID)
	api.PUT("/classes/:id", controllers.UpdateClass)
	api.DELETE("/classes/:id", controllers.DeleteClass)

	// Subjects
	api.POST("/subjects", controllers.CreateSubject)
	api.GET("/subjects", controllers.GetSubjects)
	api.GET("/subjects/:id", controllers.GetSubjectByID)
	api.PUT("/subjects/:id", controllers.UpdateSubject)
	api.DELETE("/subjects/:id", controllers.DeleteSubject)

	// Lessons
	api.POST("/lessons", controllers.CreateLesson)
	api.GET("/lessons", controllers.GetLessons)
	api.GET("/lessons/:id", controllers.GetLessonByID)
	api.PUT("/lessons/:id", controllers.UpdateLesson)
	api.DELETE("/lessons/:id", controllers.DeleteLesson)

	// Grades
	api.POST("/grades", controllers.CreateGrade)
	api.GET("/grades", controllers.GetGrades)
	api.GET("/grades/:id", controllers.GetGradeByID)
	api.PUT("/grades/:id", controllers.UpdateGrade)
	api.DELETE("/grades/:id", controllers.DeleteGrade)
}