package main

import (
	"Preload/Database"
	"Preload/models"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func main() {
	Database.ConnectDatabase()
	Database.DB.AutoMigrate(&models.Student{})
	Database.DB.AutoMigrate(&models.Bill{})
	Database.DB.AutoMigrate(&models.Profile{})
	Database.DB.AutoMigrate(&models.Course{})

	router := gin.Default()

	studentGroup := router.Group("/student")
	studentGroup.POST("", func(context *gin.Context) {
		var student models.Student
		context.ShouldBindJSON(&student)
		err := student.Save(context)
		if err != nil {
			context.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
			return
		}

		context.JSON(200, student)
	})

	studentGroup.GET("", func(context *gin.Context) {
		var students []models.Student
		err := Database.DB.Preload("Courses").Find(&students).Error
		if err != nil {
			context.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
			return
		}

		context.JSON(200, students)
	})

	// courses_id => 1-3-2-3
	studentGroup.POST("/:student_id/courses/:courses_id", func(context *gin.Context) {
		studentID := context.Param("student_id")
		coursesID := context.Param("courses_id")

		var student models.Student
		err := Database.DB.Take(&student, studentID).Error
		if err != nil {
			context.AbortWithStatusJSON(500, gin.H{"message": err.Error(), "type": "StudentNotFound"})
			return
		}

		// Regexp
		validRegex := regexp.MustCompile("^[1-8](,[1-8])*$")
		if validRegex.MatchString(coursesID) != true {
			context.AbortWithStatusJSON(500, gin.H{"message": "syntax error"})
			return
		}

		// courses list IDs
		coursesIDs := strings.Split(coursesID, ",")
		// 1,2,3,4

		for _, ID := range coursesIDs {
			// Get the course
			var course models.Course
			// Assign the course from Database
			err = Database.DB.Where("id = ?", ID).Take(&course).Error
			// If err == nil => if the course available
			if err == nil {
				// Append the course to students course list
				student.Courses = append(student.Courses, course)
			}
		}

		// Student => courses
		// Saving the student
		err = Database.DB.Save(&student).Error
		if err != nil {
			context.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "Validate Success",
			"student": student,
		})

	})

	billGroup := router.Group("/bill")
	billGroup.POST("", func(context *gin.Context) {
		var bill models.Bill
		context.ShouldBindJSON(&bill)
		err := bill.Save(context)
		if err != nil {
			context.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
			return
		}
		context.JSON(200, bill)
	})

	profileGroup := router.Group("/profile")
	profileGroup.POST("", func(context *gin.Context) {
		var profile models.Profile
		context.ShouldBindJSON(&profile)
		err := profile.Save()
		if err != nil {
			context.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
			return
		}
		context.JSON(200, profile)
	})

	courseGroup := router.Group("/course")
	courseGroup.POST("", func(context *gin.Context) {
		var course models.Course
		context.ShouldBindJSON(&course)
		err := course.Save()
		if err != nil {
			context.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
			return
		}
		context.JSON(200, course)
	})

	router.Run(":8080")

}
