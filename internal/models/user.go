package models

type User struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Email           string
	Password        string
	YandexID        int
	Gender          string
	LevelOfTraining string
	Inventory       string
	Target          string
	Weight          int
	Age             int
	Height          float64
	DesiredWeight   int
	FilledInData    bool
	Details         string
}
