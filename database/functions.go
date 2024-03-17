package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"monitoring/config"
)

func Init(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=Europe/Vilnius",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Connection{})
	if err != nil {
		panic(err)
	}
	return db
}

func GetAllConnections(db *gorm.DB) []Connection {
	var conn []Connection
	db.Find(&conn)
	return conn
}
