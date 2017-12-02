package common

import (
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
)

func SaveHtml(url string, content string) error {
	siteName, err := GetSiteName(url)
	if err != nil {
		return err
	}
	filename := siteName + ".html"

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := f.WriteString(content)
	f.Sync()

	log.Printf("wrote %d bytes\n", b)
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

	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return err
	}

	return nil
}
