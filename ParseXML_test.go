package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseXML(t *testing.T) {
	xml, err := ioutil.ReadFile("./test_data/test.xml")
	if err != nil {
		t.Fatal(err)
	}
	topic, err := ParseXML(xml)
	if err != nil {
		t.Fatal(err)
	}
	actual := []byte(fmt.Sprint(*topic))
	expected, err := ioutil.ReadFile("./test_data/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(actual, expected) {
		t.Fatal("want %v\ngot %v", expected, actual)
	}
}
