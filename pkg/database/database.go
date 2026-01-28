package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() (*gorm.DB, error) {
	//1.  set datasource db config
	dsn := "host=localhost user=postgres password=123 dbname=quiz_01 port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	//2. open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to Connect", err)
	}

	//3. test create schema oe di db
	//db.Exec("CREATE SCHEMA IF NOT EXISTS OE")

	//3.1 create variable sqlDB agar bisa akses semua function di gorm.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	//3.2 set connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connected successfully")

	//4. return value gormDB
	return db, nil

}
