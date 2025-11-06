package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string    `gorm:"size:50;not null" json:"first_name" binding:"required"`
	LastName  string    `gorm:"size:50;not null" json:"last_name" binding:"required"`
	ClassID   *uuid.UUID `gorm:"type:uuid" json:"class_id,omitempty"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
	Email     *string    `gorm:"size:100;unique" json:"email,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Relations
	Class  *Class  `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Grades []Grade `gorm:"foreignKey:StudentID" json:"grades,omitempty"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
