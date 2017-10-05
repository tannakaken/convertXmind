package main

import(
  "testing"
  "io/ioutil"
  "bytes"
  "os"
)

func TestMakeOrg(t *testing.T) {
  tmp := "./test_data/tmp"
  xml, err := ioutil.ReadFile("./test_data/test.xml")
  if err != nil {
    t.Errorf(err.Error())
  }
  data := ParseXML(xml)
  MakeOrg(data,tmp)
  defer os.Remove(tmp)
  actual, err := ioutil.ReadFile(tmp)
  if err != nil {
    t.Errorf(err.Error())
  }
  expected, err := ioutil.ReadFile("./test_data/test.org")
  if err != nil {
    t.Errorf(err.Error())
  }
  if !bytes.Equal(actual,expected) {
    t.Errorf("got %s\nwant %s", actual, expected)
  }
}
