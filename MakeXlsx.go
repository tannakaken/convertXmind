package main

import (
  "log"
  "github.com/tealeg/xlsx"
)

func toXlsx(topic Topic, sheet *xlsx.Sheet, level int, row *int, index int) {
  sheet.Cell(*row,level).SetInt(index+1)
  sheet.Cell(*row,level+1).Value = topic.Title
  *row++
  for index,childTopic := range topic.Children.Topics.Topic {
    toXlsx(childTopic, sheet, level+1,row,index)
  }
}

func MakeXlsx(topic Topic, dist string) {
  file := xlsx.NewFile()
  sheet, err := file.AddSheet("Sheet1")
  if err != nil {
    log.Fatal(err)
  }
  row := 2
  sheet.Cell(1,1).Value = topic.Title
  for index,childTopic := range topic.Children.Topics.Topic {
    toXlsx(childTopic, sheet, 1, &row, index)
  }
  err = file.Save(dist)
  if err != nil {
    log.Fatal(err)
  }
}

