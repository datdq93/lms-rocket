package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system (Student, Teacher, or Admin)
type User struct {
	ID                string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	Email             string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash      string     `gorm:"type:varchar(255);not null" json:"-"`
	Name              string     `gorm:"type:varchar(255);not null" json:"name"`
	Role              string     `gorm:"type:enum('student','teacher','admin');not null;index" json:"role"`
	AvatarURL         *string    `gorm:"type:varchar(500)" json:"avatar_url,omitempty"`
	EmailVerified     bool       `gorm:"default:false" json:"email_verified"`
	ResetToken        *string    `gorm:"type:varchar(255);index" json:"-"`
	ResetTokenExpires *time.Time `json:"-"`
	LastLoginAt       *time.Time `json:"last_login_at,omitempty"`
	IsActive          bool       `gorm:"default:true;index" json:"is_active"`
	DeletedAt         *time.Time `gorm:"index" json:"-"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`

	// Relationships
	Courses     []Course     `gorm:"foreignKey:TeacherID" json:"courses,omitempty"`
	Enrollments []Enrollment `gorm:"foreignKey:UserID" json:"enrollments,omitempty"`
}

// BeforeCreate hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

// TableName returns the table name
func (User) TableName() string {
	return "users"
}
