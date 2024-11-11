package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/lms-backend/models"
	"github.com/joshua468/lms-backend/services"
)

// GetAllCourses returns a list of all courses
func GetAllCourses(c *gin.Context) {
	courses, err := services.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

// GetCourseByID returns a specific course by ID
func GetCourseByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	course, err := services.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"course": course})
}

// CreateCourse handles requests to create a new course
func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"course": course})
}

// UpdateCourse handles requests to update an existing course
func UpdateCourse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.ID = uint(id)

	if err := services.UpdateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

// DeleteCourse handles requests to delete a course
func DeleteCourse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteCourse(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
