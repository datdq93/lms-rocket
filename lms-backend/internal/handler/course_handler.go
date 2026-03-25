package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lms-rocket/lms-backend/internal/domain"
	"github.com/lms-rocket/lms-backend/internal/service"
)

// CourseHandler handles course HTTP requests
type CourseHandler struct {
	courseService service.CourseService
}

// NewCourseHandler creates a new course handler
func NewCourseHandler(courseService service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

// ListCourses returns a list of courses
func (h *CourseHandler) ListCourses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := make(map[string]interface{})
	if level := c.Query("level"); level != "" {
		filters["level"] = level
	}

	courses, total, err := h.courseService.ListCourses(page, limit, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"courses": courses,
			"meta": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// GetCourse returns a specific course by slug
func (h *CourseHandler) GetCourse(c *gin.Context) {
	slug := c.Param("slug")
	course, err := h.courseService.GetCourseBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    course,
	})
}

// CreateRequest represents course creation request
type CreateCourseRequest struct {
	Title            string  `json:"title" binding:"required,min=3,max=255"`
	Slug             string  `json:"slug" binding:"required,min=3,max=255"`
	Description      *string `json:"description"`
	ShortDescription *string `json:"short_description"`
	Price            float64 `json:"price"`
	Level            *string `json:"level"`
	ThumbnailURL     *string `json:"thumbnail_url"`
	TrailerVideoURL  *string `json:"trailer_video_url"`
}

// CreateCourse creates a new course (teacher/admin only)
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := &domain.Course{
		Title:            req.Title,
		Slug:             req.Slug,
		Description:      req.Description,
		ShortDescription: req.ShortDescription,
		Price:            req.Price,
		Level:            req.Level,
		ThumbnailURL:     req.ThumbnailURL,
		TrailerVideoURL:  req.TrailerVideoURL,
	}

	if err := h.courseService.CreateCourse(userID.(string), course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    course,
	})
}

// UpdateCourse updates a course (teacher/admin only)
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.courseService.UpdateCourse(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    course,
	})
}

// DeleteCourse deletes a course (teacher/admin only)
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	if err := h.courseService.DeleteCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Course deleted",
	})
}

// PublishCourse publishes a course (teacher/admin only)
func (h *CourseHandler) PublishCourse(c *gin.Context) {
	id := c.Param("id")

	if err := h.courseService.PublishCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Course published",
	})
}

// UnpublishCourse unpublishes a course (teacher/admin only)
func (h *CourseHandler) UnpublishCourse(c *gin.Context) {
	id := c.Param("id")

	if err := h.courseService.UnpublishCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Course unpublished",
	})
}

// ListCourseStudents returns students enrolled in a course
func (h *CourseHandler) ListCourseStudents(c *gin.Context) {
	// TODO: Implement
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"students": []interface{}{},
		},
	})
}
