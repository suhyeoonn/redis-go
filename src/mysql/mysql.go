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

func Create() (err error) { // TODO model로 이동
	stmt, err := DB.Prepare("INSERT INTO photos VALUES( ?, ?, ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rs, err := stmt.Exec(1, 1, "a", "b", "c")
	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	defer stmt.Close()
	return
}
