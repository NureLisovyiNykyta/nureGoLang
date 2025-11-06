package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Lesson struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ClassID   uuid.UUID `gorm:"type:uuid;not null" json:"class_id"`
	SubjectID uuid.UUID `gorm:"type:uuid;not null" json:"subject_id"`
	TeacherID uuid.UUID `gorm:"type:uuid;not null" json:"teacher_id"`
	Date      time.Time `gorm:"not null" json:"date"`
	Topic     *string   `gorm:"size:200" json:"topic,omitempty"`
	Homework  *string   `gorm:"type:text" json:"homework,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Class   Class   `gorm:"foreignKey:ClassID" json:"-"`
	Subject Subject `gorm:"foreignKey:SubjectID" json:"-"`
	Teacher Teacher `gorm:"foreignKey:TeacherID" json:"-"`
	Grades  []Grade `gorm:"foreignKey:LessonID" json:"grades,omitempty"`
}

func (l *Lesson) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}
