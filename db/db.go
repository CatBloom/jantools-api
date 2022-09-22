package db

import (
	"database/sql"
	"os"

	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/lib/pq"
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
}

func connectDb() *gorm.DB {
	// ローカル環境用
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }
	// url := os.Getenv("POSTGRE_CONNECT")
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=require"

	sqlDB, err := sql.Open("postgres", connection)
	if err != nil {
		panic("Error loading open postgres")
	}
	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	// ローカル環境用
	// db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic("Error loading open postgres")
	}

	return db
}
