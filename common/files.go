package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func SaveHtml(url string, data string) error {
	siteName, err := GetSiteName(url)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(siteName+".html", []byte(data), 0644); err != nil {
		return err
	}
	log.Println("Html saved")

	return nil
}

func SaveJson(filename string, data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename+".json", byteData, 0644); err != nil {
		return err
	}
	return nil
}

func LoadJson(filename string, data interface{}) error {
	byteData, err := ioutil.ReadFile(filename + ".json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, data) //Expects data passed by reference
	if err != nil {
		return err
	}
	return nil
}
