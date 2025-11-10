package models

import "time"

type Professor struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SubjectID string `json:"subject_id"`
}

type Subject struct {
	ID                string      `json:"id" gorm:"primaryKey"`
	Name              string      `json:"name"`
	SubjectFees       string      `json:"subject_fees"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	SubjectProfessors []Professor `json:"subjectProfessors" gorm:"-"` // ignore for now
}

func (Subject) TableName() string {
	return `"Subject"`
}
