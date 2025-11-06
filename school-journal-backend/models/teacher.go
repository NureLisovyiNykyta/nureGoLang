package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string    `gorm:"size:50;not null" json:"first_name"`
	LastName  string    `gorm:"size:50;not null" json:"last_name"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	SubjectID *uuid.UUID `gorm:"type:uuid" json:"subject_id,omitempty"` // optional
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Subject *Subject `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	Classes []Class  `gorm:"foreignKey:TeacherID" json:"-"`
	Lessons []Lesson `gorm:"foreignKey:TeacherID" json:"-"`
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
