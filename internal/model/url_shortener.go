package model

import (
	"time"
)

type URLShortener struct {
	ID          uint       `gorm:"primaryKey"`
	ShortCode   string     `gorm:"uniqueIndex;size:16;not null"` 
	OriginalURL string     `gorm:"type:text;not null"`           
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	ExpiredAt   *time.Time 
	ClickCount  int        `gorm:"default:0"`
}
