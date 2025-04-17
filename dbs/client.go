package dbs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewMySQL(uri string) (*Database, error) {
	database, err := gorm.Open(mysql.Open(uri), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	sqlDB, _ := database.DB()
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(20)

	return &Database{
		DB: database,
	}, nil
}

func NewPG(uri string) (*Database, error) {
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	sqlDB, _ := database.DB()
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(20)

	return &Database{
		DB: database,
	}, nil
}
