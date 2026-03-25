package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Course represents a course in the LMS
type Course struct {
	ID                string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	Title             string     `gorm:"type:varchar(255);not null;index" json:"title"`
	Slug              string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Description       *string    `gorm:"type:text" json:"description,omitempty"`
	ShortDescription  *string    `gorm:"type:varchar(500)" json:"short_description,omitempty"`
	Price             float64    `gorm:"type:decimal(10,2);default:0" json:"price"`
	Currency          string     `gorm:"type:varchar(3);default:'USD'" json:"currency"`
	ThumbnailURL      *string    `gorm:"type:varchar(500)" json:"thumbnail_url,omitempty"`
	TrailerVideoURL   *string    `gorm:"type:varchar(500)" json:"trailer_video_url,omitempty"`
	TeacherID         string     `gorm:"type:varchar(36);not null;index" json:"teacher_id"`
	Status            string     `gorm:"type:enum('draft','published','archived');default:'draft';index" json:"status"`
	Level             *string    `gorm:"type:enum('beginner','intermediate','advanced')" json:"level,omitempty"`
	DurationMinutes   int        `gorm:"default:0" json:"duration_minutes"`
	TotalLessons      int        `gorm:"default:0" json:"total_lessons"`
	PublishedAt       *time.Time `json:"published_at,omitempty"`
	MetaTitle         *string    `gorm:"type:varchar(255)" json:"meta_title,omitempty"`
	MetaDescription   *string    `gorm:"type:varchar(500)" json:"meta_description,omitempty"`
	IsFeatured        bool       `gorm:"default:false;index" json:"is_featured"`
	DeletedAt         *time.Time `gorm:"index" json:"-"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`

	// Relationships
	Teacher   User      `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	Sections  []Section `gorm:"foreignKey:CourseID" json:"sections,omitempty"`
}

// BeforeCreate hook to generate UUID
func (c *Course) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

// TableName returns the table name
func (Course) TableName() string {
	return "courses"
}
