package database

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/go-sql-driver/mysql"
	"github.com/rizkyunm/senabung-api/campaign"
	"github.com/rizkyunm/senabung-api/transaction"
	"github.com/rizkyunm/senabung-api/user"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	dbHost, dbUser, dbPass, dbName, certPath string
	db                                       *gorm.DB
)

func newClient() *gorm.DB {
	var err error

	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	certPath = os.Getenv("CERT_PATH")

	dir, _ := os.Getwd()

	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(dir + certPath)
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}

	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})

	// try to connect to mysql database.
	cfg := mysql.Config{
		User:                 dbUser,
		Passwd:               dbPass,
		Addr:                 dbHost + ":3306",
		Net:                  "tcp",
		DBName:               dbName,
		Loc:                  time.Local,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	cfg.TLSConfig = "custom"

	str := cfg.FormatDSN()

	db, err = gorm.Open(gormMysql.Open(str), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err.Error())
	}

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&campaign.Campaign{},
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
