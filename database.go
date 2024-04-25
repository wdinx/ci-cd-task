package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
}

func InitDB(config Config) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("error when connecting to the database: %v", err)
	}

	log.Println("connected to the database")

	Migrate(DB)

}

func Migrate(db *gorm.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS posts(id INT PRIMARY KEY AUTO_INCREMENT,title VARCHAR(255) NOT NULL,content VARCHAR(255) NOT NULL);")
	log.Println("database migration success")
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitConfigMySQL() Config {
	return Config{
		DBName: os.Getenv("DBNAME"),
		DBUser: os.Getenv("DBUSER"),
		DBPass: os.Getenv("DBPASS"),
		DBHost: os.Getenv("DBHOST"),
		DBPort: os.Getenv("DBPORT"),
	}
}
