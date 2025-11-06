package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Grade struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	StudentID uuid.UUID `gorm:"type:uuid;not null" json:"student_id"`
	LessonID  uuid.UUID `gorm:"type:uuid;not null" json:"lesson_id"`
	Value     int       `gorm:"not null" json:"value"` // use 1-12 or 0-100 depending on grading scheme
	Comment   *string   `gorm:"type:text" json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Student Student `gorm:"foreignKey:StudentID" json:"-"`
	Lesson  Lesson  `gorm:"foreignKey:LessonID" json:"-"`
}

func (g *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}
