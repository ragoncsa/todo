package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/ragoncsa/todo/authz"
	"github.com/ragoncsa/todo/config"
	docs "github.com/ragoncsa/todo/docs"
	"github.com/ragoncsa/todo/gorm"
	"github.com/ragoncsa/todo/http"
	"github.com/spf13/viper"
)

// @title        Tasks service
// @version      1.0
// @description  This is a sample server that manages tasks.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// We are not using an API key here, but in OpenAPI v2 there is no better way to configure this.
// @securitydefinitions.apikey  JWT
// @in                          header
// @name                        Authorization

// Added workaround due to issues with environment variables in Viper
// https://github.com/spf13/viper/issues/761
func overrideUsingEnvVars(config *config.Config) {
	if host, present := os.LookupEnv("DB_HOST"); present {
		config.Database.Host = host
	}
}

func loadConfig() *config.Config {
	env := "local"
	if v, present := os.LookupEnv("ENV"); present {
		env = v
	}
	viper.SetConfigName(fmt.Sprintf("%s-env", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var conf config.Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}
	overrideUsingEnvVars(&conf)
	return &conf
}

func initServer() *http.Server {
	conf := loadConfig()
	db, err := gorm.Connect(conf)
	if err != nil {
		log.Panicln("failed to connect database")
	}
	log.Println("Successfully connected to database")
	gorm.RunMigration(db)

	docs.SwaggerInfo.BasePath = "/"

	server := http.InitServer(conf)
	tsDB := &gorm.TaskService{DB: db}
	tsHTTP := http.TaskService{
		Service:     tsDB,
		AuthzClient: authz.New(conf),
	}
	server.RegisterRoutes(&tsHTTP)
	return server
}

var server *http.Server
var ginLambda *ginadapter.GinLambdaV2

func init() {
	log.Printf("Initializing Gin server - BEGIN")
	server = initServer()

	if lambdaProxyOn, present := os.LookupEnv("ENABLE_GIN_LAMBDA_PROXY"); present && lambdaProxyOn == "TRUE" {
		ginLambda = ginadapter.NewV2(server.Router())
	}
	log.Printf("Initializing Gin server - END")
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if lambdaProxyOn, present := os.LookupEnv("ENABLE_GIN_LAMBDA_PROXY"); present && lambdaProxyOn == "TRUE" {
		lambda.Start(Handler)
	} else {
		server.Start()
	}
}
