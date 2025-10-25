package helper

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=test1234 dbname=tarmac_fox sslmode=disable"), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect to database")
	}
}
func GetDatabaseInstance() *gorm.DB {
	return db
}
func CloseDatabase() {
	sql, err := db.DB()

	if err != nil {
		panic("failed to close database")
	}

	sql.Close()
}