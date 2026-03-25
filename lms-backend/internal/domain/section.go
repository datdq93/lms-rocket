package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Section represents a section/chapter in a course
type Section struct {
	ID          string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	CourseID    string    `gorm:"type:varchar(36);not null;index" json:"course_id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	SortOrder   int       `gorm:"not null;default:0" json:"sort_order"`
	IsPublished bool      `gorm:"default:true" json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Course  Course   `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Lessons []Lesson `gorm:"foreignKey:SectionID" json:"lessons,omitempty"`
}

// BeforeCreate hook to generate UUID
func (s *Section) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// TableName returns the table name
func (Section) TableName() string {
	return "sections"
}
