package settings

import (
	"encoding/json"
	"log"
	"os"
)

type Setting struct {
	Database struct {
		Host     string `json: "host"`
		User     string `json: "user"`
		Password string `json: "password"`
		Dbname   string `json: "Dbname"`
		Port     int    `json: "port"`
	} `json: "database"`
}

func LoadSetting() (Setting, error) {
	file, err := os.Open("dbinfo.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var config Setting
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config, err
}
