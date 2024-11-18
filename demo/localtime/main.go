package main

import (
	"bytes"
	"github.com/26huitailang/zip"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	contents := []byte("Hello World")
	fzip, err := os.Create(`./test.zip`)
	if err != nil {
		log.Fatalln(err)
	}
	fh := &zip.FileHeader{
		Name:   `test.txt`,
		Method: zip.Deflate,
	}
	fh.SetModTime(time.Now(), time.Local)
	zipw := zip.NewWriter(fzip)
	w, err := zipw.CreateHeader(fh)
	if err != nil {
		log.Fatal(err)
	}
	defer zipw.Close()
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Flush()
}
