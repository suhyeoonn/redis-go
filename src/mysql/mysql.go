package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

func GetConfig() DBConfig {
	var dbConfig DBConfig
	yamlFile, _ := os.ReadFile("./src/config/config.yaml")
	err := yaml.Unmarshal([]byte(yamlFile), &dbConfig)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return dbConfig
}

var DB *sql.DB

func Open() {
	config := GetConfig()
	source := fmt.Sprintf("%v:%v@/%v", config.User, config.Password, config.DbName)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	DB = db
}
