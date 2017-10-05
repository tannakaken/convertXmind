package main

import (
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"testing"
)

func TestMakeXlsx(t *testing.T) {
	tmp := "./test_data/tmp.xlsx"
	xml, err := ioutil.ReadFile("./test_data/test.xml")
	if err != nil {
		t.Errorf(err.Error())
	}
	data := ParseXML(xml)
	MakeXlsx(data, tmp)
	defer os.Remove(tmp)
	actual, err := xlsx.OpenFile(tmp)
	if err != nil {
		t.Errorf(err.Error())
	}
	expected, err := xlsx.OpenFile("./test_data/test.xlsx")
	if err != nil {
		t.Errorf(err.Error())
	}
	for i, actualSheet := range actual.Sheets {
		expectedSheet := expected.Sheets[i]
		for j, actualRow := range actualSheet.Rows {
			expectedRow := expectedSheet.Rows[j]
			for k, actualCell := range actualRow.Cells {
				expectedCell := expectedRow.Cells[k]
				if actualCell.String() != expectedCell.String() {
					t.Errorf("got %s\nwant %s", actualCell.String(),
						expectedCell.String())
				}
			}
		}
	}
}
