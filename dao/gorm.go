package dao

import (
	"database/sql"

	"gorm.io/gorm"
)

var (
	db            *gorm.DB
	customSession *gorm.Session
)

func Init(gormDB *gorm.DB) {
	db = gormDB
	customSession = &gorm.Session{
		QueryFields: true,
	}
}

func GormDB() *gorm.DB {
	return db
}

func SqlDB() *sql.DB {
	sqlDB, _ := db.DB()
	return sqlDB
}
