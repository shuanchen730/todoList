package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnect() *gorm.DB {
	return setMysqlConnect(EnvConfig.ConnectionStrings.Mysql)
}

func setMysqlConnect(mysqlConfig MysqlConfig) *gorm.DB {
	if mysqlConfig.Port == "" {
		mysqlConfig.Port = "3306"
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database)
	mysqlClient, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		log.Fatalf("database connection failed : %v \n", err)
	}
	if mysqlClient.Error != nil {
		log.Fatalf("database error %v \n", mysqlClient.Error)
	}
	if EnvConfig.Env == "local" {
		mysqlClient = mysqlClient.Debug()
	}
	return mysqlClient
}
