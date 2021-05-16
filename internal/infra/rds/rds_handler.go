package rds

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type RDBHandler struct {
	DbService *gorm.DB
}

func NewRDBHandler() (*RDBHandler, error) {
	const (
		Dialect = "mysql"
		DBName  = "{dbname}"
		DBHost  = "{hostname}"
		DBPort  = "3306"
		DBUser  = "{username}"
	)

	dbPassword := "{password}"

	dbProtocol := fmt.Sprintf("tcp(%s:%s)", DBHost, DBPort)
	// About `parseTime=True`, read https://github.com/jinzhu/gorm/issues/18#issuecomment-29351315
	connectTemplate := "%s:%s@%s/%s?timeout=10s&parseTime=True"
	connect := fmt.Sprintf(connectTemplate, DBUser, dbPassword, dbProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		fmt.Println("DB connection initialization error")
		fmt.Println(err.Error())
		return nil, err
	}
	handler := RDBHandler{
		DbService: db,
	}
	return &handler, nil
}
