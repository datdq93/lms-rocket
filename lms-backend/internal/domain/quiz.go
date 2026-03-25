package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Quiz represents a quiz question
type Quiz struct {
	ID             string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	LessonID       string    `gorm:"type:varchar(36);not null;index" json:"lesson_id"`
	Question       string    `gorm:"type:text;not null" json:"question"`
	QuestionType   string    `gorm:"type:enum('single_choice','multiple_choice','fill_blank');not null" json:"question_type"`
	Options        string    `gorm:"type:json;not null" json:"options"`        // JSON array
	CorrectAnswers string    `gorm:"type:json;not null" json:"correct_answers"` // JSON array
	Explanation    *string   `gorm:"type:text" json:"explanation,omitempty"`
	Points         int       `gorm:"default:1" json:"points"`
	SortOrder      int       `gorm:"default:0" json:"sort_order"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// Relationships
	Lesson Lesson `gorm:"foreignKey:LessonID" json:"lesson,omitempty"`
}

// BeforeCreate hook to generate UUID
func (q *Quiz) BeforeCreate(tx *gorm.DB) error {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	return nil
}

// TableName returns the table name
func (Quiz) TableName() string {
	return "quizzes"
}
