package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestMakeOrg(t *testing.T) {
	tmp := "./test_data/tmp"
	xml, err := ioutil.ReadFile("./test_data/test.xml")
	if err != nil {
		t.Errorf(err.Error())
	}
	data, err := ParseXML(xml)
	if err != nil {
		t.Fatal(err)
	}
	err = MakeOrg(data, tmp)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp)
	actual, err := ioutil.ReadFile(tmp)
	if err != nil {
		t.Errorf(err.Error())
	}
	expected, err := ioutil.ReadFile("./test_data/test.org")
	if err != nil {
		t.Errorf(err.Error())
	}
	if !bytes.Equal(actual, expected) {
		t.Errorf("got %s\nwant %s", actual, expected)
	}
}
