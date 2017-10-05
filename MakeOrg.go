package main

import (
  "os"
  "log"
)

func makeOrg(file *os.File ,topic Topic, level int) {
  for i := 0; i < level; i++ {
    file.Write([]byte("*"))
  }
  file.Write([]byte(" "))
  file.Write([]byte(topic.Title))
  file.Write([]byte("\n"))
  for _, childTopic := range topic.Children.Topics.Topic {
    makeOrg(file, childTopic, level + 1)
  }
}

// topic構造体を再帰的にgnu emcasのorgファイルにして出力する。
func MakeOrg(topic Topic,dist string) {
  file, err := os.Create(dist)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  makeOrg(file,topic,1)
}
