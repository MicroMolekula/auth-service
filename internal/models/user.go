package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	YandexID string
}
