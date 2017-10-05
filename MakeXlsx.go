package main

import (
	"github.com/tealeg/xlsx"
)

func makeXlsx(topic *Topic, sheet *xlsx.Sheet, level int, row *int, index int) {
	sheet.Cell(*row, level).SetInt(index + 1)
	sheet.Cell(*row, level+1).Value = topic.Title
	*row++
	for index, childTopic := range topic.Children.Topics.Topic {
		makeXlsx(&childTopic, sheet, level+1, row, index)
	}
}

// Topic構造体を再帰的にxlsxファイルにして出力する。
func MakeXlsx(topic *Topic, dist string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	row := 2
	sheet.Cell(1, 1).Value = topic.Title
	for index, childTopic := range topic.Children.Topics.Topic {
		makeXlsx(&childTopic, sheet, 1, &row, index)
	}
	err = file.Save(dist)
	if err != nil {
		return err
	}
	return nil
}
