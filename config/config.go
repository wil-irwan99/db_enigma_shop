package config

import (
	"fmt"
	"gorm/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Db *gorm.DB
}

func (c *Config) initDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	env := os.Getenv("ENV")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	enigmaDb, err := db.DB()
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected...")
	}

	if env == "dev" {
		c.Db = db.Debug()
	} else if env == "migration" {
		c.Db = db.Debug()
		err := c.Db.AutoMigrate(&model.Customer{}, &model.UserCredential{}, &model.Address{}, &model.Product{})

		if err != nil {
			return
		}
	} else {
		c.Db = db
	}
}

func (c *Config) DbConn() *gorm.DB {
	return c.Db
}

func (c *Config) DbClose() {
	a, err := c.Db.DB()
	if err != nil {
		panic(err)
	}
	a.Close()
}

func NewConfig() Config {
	cfg := Config{}
	cfg.initDB()
	return cfg
}
