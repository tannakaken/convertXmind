package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/tealeg/xlsx"
)

func TestMakeXlsx(t *testing.T) {
	tmp := "./test_data/tmp.xlsx"
	xml, err := ioutil.ReadFile("./test_data/test.xml")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ParseXML(xml)
	if err != nil {
		t.Fatal(err)
	}
	err = MakeXlsx(data, tmp)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp)
	actual, err := xlsx.OpenFile(tmp)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := xlsx.OpenFile("./test_data/test.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	for i, actualSheet := range actual.Sheets {
		expectedSheet := expected.Sheets[i]
		for j, actualRow := range actualSheet.Rows {
			expectedRow := expectedSheet.Rows[j]
			for k, actualCell := range actualRow.Cells {
				expectedCell := expectedRow.Cells[k]
				if actualCell.String() != expectedCell.String() {
					t.Errorf("want %v\ngot %v", expectedCell, actualCell)
				}
			}
		}
	}
}
