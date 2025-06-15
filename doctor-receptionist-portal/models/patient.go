package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	Disease    string    `json:"disease"`
	AdmittedAt time.Time `json:"admitted_at"`
	UpdatedBy  uint      `json:"updated_by"`
}
