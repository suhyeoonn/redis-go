package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"redis-go/src/mysql"
)

func CreatePhotos(data string) (err error) {
	stmt, err := mysql.DB.Prepare("INSERT INTO photos VALUES( ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rs, err := stmt.Exec(data)
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

type StringInterfaceMap map[string]interface{}

type Photo struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

func (m *StringInterfaceMap) Scan(src interface{}) error {
	var source []byte
	_m := make(map[string]interface{})

	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))

	case nil:
		return nil
	default:
		return errors.New("incompatible type for StringInterfaceMap")
	}
	err := json.Unmarshal(source, &_m)
	if err != nil {
		return err
	}
	*m = StringInterfaceMap(_m)
	return nil
}

func (m StringInterfaceMap) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}

func GetPhotos() (photos []Photo) {
	rows, err := mysql.DB.Query("SELECT * FROM photos WHERE id = 1")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var photo Photo
		rows.Scan(&photo.Id, &photo.Data)
		photos = append(photos, photo)
	}
	defer rows.Close()
	return
}
