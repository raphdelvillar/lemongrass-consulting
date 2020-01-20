package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	createZipFile()
}

func createZipFile() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("test.zip", flags, 0644)

	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}

	defer file.Close()

	var files = []string{"test1.txt", "test2.txt", "test3.txt"}
	for _, filename := range files {
		zipw := zip.NewWriter(file)
		defer zipw.Close()

		var z Zip
		z.writer = zipw
		z.filename = filename
		if err := z.compress(); err != nil {
			log.Fatal(err)
		}
	}
}
