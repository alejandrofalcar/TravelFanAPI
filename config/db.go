package config

import "fmt"

const (
	DBHost     = "localhost"
	DBPost     = "3306"
	DBUser     = "root"
	DBPassword = "123456"
	DBName     = "travelfan"
)

func GetDBConnection() string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPost,
		DBName)
	return connection
}
