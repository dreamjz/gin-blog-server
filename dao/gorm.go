package dao

import (
	"database/sql"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init(gormDB *gorm.DB) {
	db = gormDB
}

func GormDB() *gorm.DB {
	return db
}

func SqlDB() *sql.DB {
	sqlDB, _ := db.DB()
	return sqlDB
}
