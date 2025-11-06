package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Subject struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Relations
	Teachers []Teacher `gorm:"foreignKey:SubjectID" json:"-"`
	Lessons  []Lesson  `gorm:"foreignKey:SubjectID" json:"-"`
}

func (s *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
