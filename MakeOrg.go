package main

import (
	"os"
)

func makeOrg(file *os.File, topic *Topic, level int) {
	for i := 0; i < level; i++ {
		file.Write([]byte("*"))
	}
	file.Write([]byte(" "))
	file.Write([]byte(topic.Title))
	file.Write([]byte("\n"))
	for _, childTopic := range topic.Children.Topics.Topic {
		makeOrg(file, &childTopic, level+1)
	}
}

// topic構造体を再帰的にgnu emcasのorgファイルにして出力する。
func MakeOrg(topic *Topic, dist string) error {
	file, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer file.Close()
	makeOrg(file, topic, 1)
	return nil
}
