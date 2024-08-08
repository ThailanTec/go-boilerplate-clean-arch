package migrations

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Phone    string `gorm:"unique;not null"`
	Document string `gorm:"unique;not null"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
