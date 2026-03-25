package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Enrollment represents a user's enrollment in a course
type Enrollment struct {
	ID               string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserID           string     `gorm:"type:varchar(36);not null;index" json:"user_id"`
	CourseID         string     `gorm:"type:varchar(36);not null;index" json:"course_id"`
	Status           string     `gorm:"type:enum('pending','active','completed','expired','cancelled');default:'pending';index" json:"status"`
	ProgressPercent  float64    `gorm:"type:decimal(5,2);default:0.00" json:"progress_percent"`
	CompletedLessons int        `gorm:"default:0" json:"completed_lessons"`
	TotalLessons     int        `gorm:"default:0" json:"total_lessons"`
	EnrolledAt       time.Time  `json:"enrolled_at"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
	ExpiresAt        *time.Time `json:"expires_at,omitempty"`
	LastAccessedAt   *time.Time `json:"last_accessed_at,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	// Relationships
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Course Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`
}

// BeforeCreate hook to generate UUID
func (e *Enrollment) BeforeCreate(tx *gorm.DB) error {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return nil
}

// TableName returns the table name
func (Enrollment) TableName() string {
	return "enrollments"
}
