package migrations

import "time"

type User struct {
	Id_User   string    `gorm:"primary_key;type:varchar(11);not null"`
	Username  string    `gorm:"type:varchar(20);not null"`
	Password  string    `gorm:"type:varchar(70);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}
