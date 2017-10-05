package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestGetContent(t *testing.T) {
	actual, err := GetContent("./test_data/test.xmind")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile("./test_data/test.xml")
	if err != nil {
		t.Errorf(err.Error())
	}
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %s\nwant %s", actual, expected)
	}
}
