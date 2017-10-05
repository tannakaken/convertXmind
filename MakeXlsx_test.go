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
		t.Error(err)
	}
	data, err := ParseXML(xml)
	if err != nil {
		t.Fatal(err)
	}
	MakeXlsx(data, tmp)
	defer os.Remove(tmp)
	actual, err := xlsx.OpenFile(tmp)
	if err != nil {
		t.Error(err)
	}
	expected, err := xlsx.OpenFile("./test_data/test.xlsx")
	if err != nil {
		t.Error(err)
	}
	for i, actualSheet := range actual.Sheets {
		expectedSheet := expected.Sheets[i]
		for j, actualRow := range actualSheet.Rows {
			expectedRow := expectedSheet.Rows[j]
			for k, actualCell := range actualRow.Cells {
				expectedCell := expectedRow.Cells[k]
				if actualCell.String() != expectedCell.String() {
					t.Errorf("got %v\nwant %v", actualCell, expectedCell)
				}
			}
		}
	}
}
