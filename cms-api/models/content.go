package models

import "time"

type Content struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    Title      string    `gorm:"not null" json:"title"`
    Body       string    `gorm:"not null" json:"body"`
    CategoryID uint      `json:"category_id"`
    AuthorID   uint      `json:"author_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}
