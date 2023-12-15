package pkg

import (
	"app/pkg/controllers"
	// "encoding/json"
	// "github.com/gorilla/mux"
	// // "log"
	// "github.com/gin-gonic/gin"
	// "net/http"
)

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	appConfig.AppName = "GoToko"
	appConfig.AppEnv = "development"
	appConfig.AppPort = "9000"
	appConfig.AppURL = "http://0.0.0.0:9000"

	dbConfig.DBHost = "database-service"
	dbConfig.DBUser = "root"
	dbConfig.DBPassword = "1"
	dbConfig.DBName = "ECOMMERCE"
	dbConfig.DBPort = "3306"
	dbConfig.DBDriver = "mysql"

	// arg := flag.Arg(0)

	server.Initialize(appConfig, dbConfig)
	server.Run("0.0.0.0:" + appConfig.AppPort)
	// server.Router = mux.NewRouter()
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.IndentedJSON(http.StatusOK, map[string]interface{}{"a":1, "b" : "hello"})
	// })
	// router.Run("0.0.0.0:9000")

}
