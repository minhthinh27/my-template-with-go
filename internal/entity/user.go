package entity

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	CreatedAt int64  `gorm:"autoUpdateTime:milli;column:created_at;not null"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at;not null"`
	Name      string `gorm:"size:500;column:name;not null"`
}
