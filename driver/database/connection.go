package database

import (
	"github.com/rizkyunm/senabung-api/campaign"
	"github.com/rizkyunm/senabung-api/transaction"
	"github.com/rizkyunm/senabung-api/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	dbHost, dbUser, dbPass, dbName string
	db                             *gorm.DB
)

func newClient() *gorm.DB {
	var err error

	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err.Error())
	}

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&campaign.Campaign{},
		&campaign.CampaignImage{},
		&user.User{},
		&transaction.Transaction{},
	); err != nil {
		panic(err.Error())
	}

	return db
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	return newClient()
}
