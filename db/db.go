package db

import (
	"os"

	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	if err != nil {
		panic(err)
	}
	connectDb()
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	db, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
}

func autoMigration() {
	db.AutoMigrate(
		&models.League{},
		&models.AdminsLeagues{},
		&models.Player{},
		&models.Rules{},
		&models.Game{},
		&models.Result{},
	)
	db.AutoMigrate()
}

func connectDb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	DBCONNECT := os.Getenv("POSTGRE_CONNECT")

	db, err = gorm.Open(postgres.Open(DBCONNECT), &gorm.Config{})
	if err != nil {
		panic("Error loading open postgres")
	}

	return db
}
