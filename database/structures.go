package database

type Connection struct {
	ID         int    `gorm:"primaryKey"`
	Hostname   string `gorm:"not null"`
	ClientIP   string `gorm:"not null"`
	ProvidedIP string `gorm:"not null"`
	Time       string `gorm:"not null"`
}
