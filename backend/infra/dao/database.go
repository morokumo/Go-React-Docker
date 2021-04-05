package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := "database"
	user := "postgres"
	password := "postgres"
	dbname := "postgres"
	port := "5432"
	sslMode := "disable"
	timezone := "Asia/Tokyo"
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname
	dsn += " port=" + port + " sslmode=" + sslMode + " TimeZone=" + timezone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}
