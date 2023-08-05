package types

import (
	"time"

	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	Domain     string `gorm:"uniqIndex"`
	Expiration time.Time
}
