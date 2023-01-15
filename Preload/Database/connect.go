package Database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open(Config.PSqlDsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
