package pkg


import (
	"flag"

	"test/ecommerce2/pkg/controllers"

	// "github.com/joho/godotenv"
)



func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	appConfig.AppName = "GoToko"
	appConfig.AppEnv = "development"
	appConfig.AppPort = "9000"
	appConfig.AppURL = "http://127.0.0.1:9000"

	dbConfig.DBHost = "localhost"
	dbConfig.DBUser = "tam"
	dbConfig.DBPassword = "1"
	dbConfig.DBName = "ECOMMERCE"
	dbConfig.DBPort = "3306"
	dbConfig.DBDriver = "mysql"

	flag.Parse()
	// arg := flag.Arg(0)

	server.Initialize(appConfig, dbConfig)
	server.Run("127.0.0.1:" + appConfig.AppPort)
	// if arg != "" {
	// 	server.InitCommands(appConfig, dbConfig)
	// } else {
	// 	server.Initialize(appConfig, dbConfig)
	// 	server.Run(":" + appConfig.AppPort)
	// }
}