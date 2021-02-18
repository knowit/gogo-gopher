package repositories

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
)

func InitDB() *gorm.DB {
	sqlDB, sqlOpenError := sql.Open("postgres", "mydb_dsn")

	if sqlOpenError != nil {
		log.Fatalf("Error while opening sqlDB driver to postgres. Error %v\n", sqlOpenError)
	}

	gormDB, gormOpenError := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if gormOpenError != nil {

	}

	return gormDB
}
