package gorm

import (
	"fmt"

	"github.com/ragoncsa/todo/config"
	"github.com/ragoncsa/todo/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(conf *config.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		conf.Database.Host,
		conf.Database.DBUser,
		conf.Database.DBPassword,
		conf.Database.DBName,
		conf.Database.Port,
		conf.Database.Timezone)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Task{})
}
