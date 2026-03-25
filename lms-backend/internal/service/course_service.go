package service

import (
	"github.com/lms-rocket/lms-backend/internal/domain"
	"github.com/lms-rocket/lms-backend/internal/repository"
)

// CourseService defines course operations
type CourseService interface {
	CreateCourse(teacherID string, course *domain.Course) error
	GetCourse(id string) (*domain.Course, error)
	GetCourseBySlug(slug string) (*domain.Course, error)
	UpdateCourse(id string, updates map[string]interface{}) (*domain.Course, error)
	DeleteCourse(id string) error
	ListCourses(page, limit int, filters map[string]interface{}) ([]domain.Course, int64, error)
	ListTeacherCourses(teacherID string, page, limit int) ([]domain.Course, int64, error)
	PublishCourse(id string) error
	UnpublishCourse(id string) error
}

// courseService implements CourseService
type courseService struct {
	courseRepo repository.CourseRepository
}

// NewCourseService creates a new course service
func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return &courseService{courseRepo: courseRepo}
}

func (s *courseService) CreateCourse(teacherID string, course *domain.Course) error {
	course.TeacherID = teacherID
	course.Status = "draft"
	return s.courseRepo.Create(course)
}

func (s *courseService) GetCourse(id string) (*domain.Course, error) {
	return s.courseRepo.FindByID(id)
}

func (s *courseService) GetCourseBySlug(slug string) (*domain.Course, error) {
	return s.courseRepo.FindBySlug(slug)
}

func (s *courseService) UpdateCourse(id string, updates map[string]interface{}) (*domain.Course, error) {
	course, err := s.courseRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if title, ok := updates["title"].(string); ok {
		course.Title = title
	}
	if description, ok := updates["description"].(string); ok {
		course.Description = &description
	}

	if err := s.courseRepo.Update(course); err != nil {
		return nil, err
	}

	return course, nil
}

func (s *courseService) DeleteCourse(id string) error {
	return s.courseRepo.Delete(id)
}

func (s *courseService) ListCourses(page, limit int, filters map[string]interface{}) ([]domain.Course, int64, error) {
	return s.courseRepo.List(page, limit, filters)
}

func (s *courseService) ListTeacherCourses(teacherID string, page, limit int) ([]domain.Course, int64, error) {
	return s.courseRepo.ListByTeacher(teacherID, page, limit)
}

func (s *courseService) PublishCourse(id string) error {
	course, err := s.courseRepo.FindByID(id)
	if err != nil {
		return err
	}
	course.Status = "published"
	return s.courseRepo.Update(course)
}

func (s *courseService) UnpublishCourse(id string) error {
	course, err := s.courseRepo.FindByID(id)
	if err != nil {
		return err
	}
	course.Status = "draft"
	return s.courseRepo.Update(course)
}
