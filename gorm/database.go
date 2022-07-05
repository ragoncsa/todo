package gorm

import (
	"context"
	"fmt"
	"log"
	"os"

	config2 "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/ragoncsa/todo/config"
	"github.com/ragoncsa/todo/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(conf *config.Config) (db *gorm.DB, err error) {
	password := conf.Database.DBPassword
	if lambdaProxyOn, present := os.LookupEnv("ENABLE_GIN_LAMBDA_PROXY"); present && lambdaProxyOn == "TRUE" {
		var dbEndpoint string = fmt.Sprintf("%s:%d", conf.Database.Host, conf.Database.Port)
		cfg, err := config2.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error: " + err.Error())
		}
		authenticationToken, err := auth.BuildAuthToken(
			context.TODO(), dbEndpoint, cfg.Region, conf.Database.DBUser, cfg.Credentials)
		if err != nil {
			log.Panicln("failed to create authentication token: " + err.Error())
		}
		password = authenticationToken
	}

	// on SSL config see: https://www.postgresql.org/docs/9.4/libpq-connect.html#LIBPQ-CONNECT-SSLMODE
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=prefer TimeZone=%s",
		conf.Database.Host,
		conf.Database.DBUser,
		password,
		conf.Database.DBName,
		conf.Database.Port,
		conf.Database.Timezone)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Task{})
}
