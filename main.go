package main

import (
	"fmt"

	"github.com/ragoncsa/todo/config"
	"github.com/ragoncsa/todo/gorm"
	"github.com/ragoncsa/todo/http"

	"github.com/spf13/viper"
)

func loadConfig() *config.Config {
	viper.SetConfigName("local-env")
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
	return &conf
}

func main() {
	conf := loadConfig()
	db, err := gorm.Connect(conf)
	if err != nil {
		panic("failed to connect database")
	}
	gorm.RunMigration(db)

	tsDB := &gorm.TaskService{DB: db}

	server := http.InitServer()

	tsHTTP := http.TaskService{
		Server:  server,
		Service: tsDB,
	}
	tsHTTP.RegisterRoutes()

	http.StartServer(server)
}
