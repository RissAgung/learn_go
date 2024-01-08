package migrations

import "time"

type BlackListToken struct {
	Id_Token  int64     `gorm:"primary_key;"`
	Token     string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}
