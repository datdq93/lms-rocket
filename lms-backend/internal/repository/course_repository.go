package repository

import (
	"github.com/lms-rocket/lms-backend/internal/domain"
	"gorm.io/gorm"
)

// CourseRepository defines the interface for course data access
type CourseRepository interface {
	Create(course *domain.Course) error
	FindByID(id string) (*domain.Course, error)
	FindBySlug(slug string) (*domain.Course, error)
	Update(course *domain.Course) error
	Delete(id string) error
	List(page, limit int, filters map[string]interface{}) ([]domain.Course, int64, error)
	ListByTeacher(teacherID string, page, limit int) ([]domain.Course, int64, error)
}

// courseRepository implements CourseRepository
type courseRepository struct {
	db *gorm.DB
}

// NewCourseRepository creates a new course repository
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(course *domain.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) FindByID(id string) (*domain.Course, error) {
	var course domain.Course
	if err := r.db.Preload("Teacher").First(&course, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) FindBySlug(slug string) (*domain.Course, error) {
	var course domain.Course
	if err := r.db.Preload("Teacher").Preload("Sections.Lessons").First(&course, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) Update(course *domain.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) Delete(id string) error {
	return r.db.Delete(&domain.Course{}, "id = ?", id).Error
}

func (r *courseRepository) List(page, limit int, filters map[string]interface{}) ([]domain.Course, int64, error) {
	var courses []domain.Course
	var total int64

	query := r.db.Model(&domain.Course{}).Where("status = ?", "published")

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if level, ok := filters["level"]; ok {
		query = query.Where("level = ?", level)
	}
	if teacherID, ok := filters["teacher_id"]; ok {
		query = query.Where("teacher_id = ?", teacherID)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("Teacher").Offset(offset).Limit(limit).Find(&courses).Error

	return courses, total, err
}

func (r *courseRepository) ListByTeacher(teacherID string, page, limit int) ([]domain.Course, int64, error) {
	var courses []domain.Course
	var total int64

	offset := (page - 1) * limit
	r.db.Model(&domain.Course{}).Where("teacher_id = ?", teacherID).Count(&total)
	err := r.db.Where("teacher_id = ?", teacherID).Offset(offset).Limit(limit).Find(&courses).Error

	return courses, total, err
}
