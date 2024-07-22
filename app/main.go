package main

import (
	"log"
	"todoList/config"
	_cardHandler "todoList/domain/Card/delivery/http"
	_mysqlCard "todoList/domain/Card/repository/mysql"
	_cardUsecase "todoList/domain/Card/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB

func initSetup() {

	// 綁定環境變數
	err := config.InitialEnvConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	// Mysql 資料庫
	MysqlDb = config.NewMysqlConnect()
}

func StartAPIServer() {
	initSetup()
	g := gin.New()

	mysqlCard := _mysqlCard.NewMysqlCardRepository(MysqlDb)
	cardUsecase := _cardUsecase.NewCardUsecase(mysqlCard)
	_cardHandler.NewCardHandler(g, cardUsecase)
	_ = g.Run(":" + config.EnvConfig.Port)

}

func main() {
	StartAPIServer()

}
