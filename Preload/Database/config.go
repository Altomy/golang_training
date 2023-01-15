package Database

import "fmt"

type ConfigStruct struct {
	DBHost     string `yaml:"DBHost"`
	DBUser     string `yaml:"DBUser"`
	DBPassword string `yaml:"DBPassword"`
	DBName     string `yaml:"DBName"`
	DBPort     string `yaml:"DBPort"`
	DBSSLMode  string `yaml:"DBSSlMode"`
	DBTimeZone string `yaml:"DBTimeZone"`
}

var Config = ConfigStruct{
	DBHost:     "localhost",
	DBUser:     "postgres",
	DBPassword: "00962s00962S!",
	DBName:     "training",
	DBPort:     "5432",
	DBSSLMode:  "disable",
	DBTimeZone: "Asia/Amman",
}

func (db *ConfigStruct) PSqlDsn() string {

	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", db.DBHost, db.DBUser, db.DBPassword, db.DBName, db.DBPort, db.DBSSLMode, db.DBTimeZone)
}
