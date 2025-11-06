package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Class struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"size:20;not null;unique" json:"name"` // e.g., "10-A"
	TeacherID *uuid.UUID `gorm:"type:uuid" json:"teacher_id,omitempty"` // class teacher (homeroom)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Teacher  *Teacher  `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	Students []Student `gorm:"foreignKey:ClassID" json:"students,omitempty"`
	Lessons  []Lesson  `gorm:"foreignKey:ClassID" json:"lessons,omitempty"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
