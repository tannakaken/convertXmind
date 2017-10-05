package main

import (
	"archive/zip"
	"io"
	"log"
)

// xmindファイルからcontent.xmlを取り出す。
func GetContent(src string) []byte {
	r, err := zip.OpenReader(src)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	var content *zip.File
	for _, f := range r.File {
		if f.Name == "content.xml" {
			content = f
		}
	}
	if content == nil {
		log.Fatal("content not found")
	}
	rc, err := content.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()
	buf := make([]byte, content.UncompressedSize)
	_, err = io.ReadFull(rc, buf)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}
