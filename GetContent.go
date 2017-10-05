package main

import (
	"archive/zip"
	"io"
)

// xmindファイルからcontent.xmlを取り出す。
func GetContent(src string) ([]byte, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var content *zip.File
	for _, f := range r.File {
		if f.Name == "content.xml" {
			content = f
		}
	}
	if content == nil {
		return nil, err
	}
	rc, err := content.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	buf := make([]byte, content.UncompressedSize)
	_, err = io.ReadFull(rc, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
