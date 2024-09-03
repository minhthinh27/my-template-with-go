package entity

type Article struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	CreatedAt int64  `gorm:"autoUpdateTime:milli;column:created_at;not null"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at;not null"`
	Author    string `gorm:"size:256;column:author;not null;index"`
	Title     string `gorm:"size:500;column:title;not null"`
}
