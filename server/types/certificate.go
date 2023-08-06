package types

import "time"

type Certificate struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Domain         string    `gorm:"uniqueIndex"`
	ExpirationDate time.Time `gorm:"index"`
	DaysUntilExp   int
	Status         int
}
