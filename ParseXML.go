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
  Title string `xml:"title"`
  Children struct {
    Topics struct {
      Topic []Topic `xml:"topic"`
    } `xml:"topics"`
  } `xml:"children"`
}

func ParseXML(xmlData []byte) Topic {
  t := Content{}
  xml.Unmarshal(xmlData, &t)
  return t.Sheet.Topic
}

