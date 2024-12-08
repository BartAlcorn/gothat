package main

/*
	! go run ./cmd/localizer

	Downloads the various script/library files from CDNs,
	renames them and saves them to named folders within
	./pkg/embedded/assets/js/

	 * Edit the Scripts slice in ./pkg/embedded/assets/js/embedded.go
*/

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bartalcorn/gothat/pkg/embedded"
)

type RemoteFile struct {
	Folder string
	Name   string
	URL    string
}

func main() {
	scripts := embedded.Scripts
	sheets := embedded.StyleSheets
	embedPath := "./pkg/embedded/assets/js/"

	htmxPath := embedPath + "htmx"
	ensurePath(htmxPath)

	for _, scrpt := range scripts {
		body, err := fetch(scrpt.URL)
		if err != nil {
			break
		}
		path := embedPath + scrpt.Folder
		ensurePath(path)

		if err := os.WriteFile(path+"/"+scrpt.Name, body, 0666); err != nil {
			log.Fatal(err)
		}
	}

	for _, sheet := range sheets {
		body, err := fetch(sheet.URL)
		if err != nil {
			break
		}
		path := embedPath + sheet.Folder
		ensurePath(path)

		if err := os.WriteFile(path+"/"+sheet.Name, body, 0666); err != nil {
			log.Fatal(err)
		}
	}

}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching CDN:", err)
	}

	if resp.StatusCode != 200 {
		log.Println("Error fetching CDN:", url, resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err
}

func ensurePath(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
