package common

import (
	"os"
	"log"
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
