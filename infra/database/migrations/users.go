package migrations

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `gorm:"type:uuid;primary_key;"`
	Name      string `gorm:"not null"`
	Phone     string `gorm:"unique;not null"`
	Document  string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
