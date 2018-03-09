package main

import (
	"encoding/xml"
)

type Content struct {
	Sheet struct {
		Topic Topic `xml:"topic"`
	} `xml:"sheet"`
}

type Topic struct {
	Title    string `xml:"title"`
	Children struct {
		Topics struct {
			Topic []Topic `xml:"topic"`
		} `xml:"topics"`
	} `xml:"children"`
}

// xmindファイルから取り出したcontent.xmlを構造体にし、一番上のレベルのtopicを取り出す。
func ParseXML(xmlData []byte) (*Topic, error) {
	t := Content{}
	err := xml.Unmarshal(xmlData, &t)
	return &t.Sheet.Topic, err
}
