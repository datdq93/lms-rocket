package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Lesson represents a lesson within a section
type Lesson struct {
	ID                  string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	SectionID           string    `gorm:"type:varchar(36);not null;index" json:"section_id"`
	Title               string    `gorm:"type:varchar(255);not null" json:"title"`
	Description         *string   `gorm:"type:text" json:"description,omitempty"`
	ContentType         string    `gorm:"type:enum('video','text','file','quiz');not null" json:"content_type"`
	VideoURL            *string   `gorm:"type:varchar(500)" json:"video_url,omitempty"`
	VideoDuration       int       `gorm:"default:0" json:"video_duration"`
	TextContent         *string   `gorm:"type:longtext" json:"text_content,omitempty"`
	FileURL             *string   `gorm:"type:varchar(500)" json:"file_url,omitempty"`
	FileName            *string   `gorm:"type:varchar(255)" json:"file_name,omitempty"`
	FileSize            int       `gorm:"default:0" json:"file_size"`
	IsFreePreview       bool      `gorm:"default:false" json:"is_free_preview"`
	SortOrder           int       `gorm:"not null;default:0" json:"sort_order"`
	IsPublished         bool      `gorm:"default:true" json:"is_published"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	// Relationships
	Section Section `gorm:"foreignKey:SectionID" json:"section,omitempty"`
	Quizzes []Quiz  `gorm:"foreignKey:LessonID" json:"quizzes,omitempty"`
}

// BeforeCreate hook to generate UUID
func (l *Lesson) BeforeCreate(tx *gorm.DB) error {
	if l.ID == "" {
		l.ID = uuid.New().String()
	}
	return nil
}

// TableName returns the table name
func (Lesson) TableName() string {
	return "lessons"
}
