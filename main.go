package main

import (
	"fmt"
	"os"
	"path"
)

// xmindファイルをコマンドライン引数で指定されたファイル形式に変換する
func main() {
	usage :=
		`USAGE: convertXmind format source [distination]
DESCRIPTION:
  convert XMind 8 file to given format.
  if you do not type distination file name,
  it will be source file name that its extension changed to format name.

VERSION: 1.0

FORMATS:
  xlsx  xlsx file for Microsoft Excel
  org   org file for Gnu Emacs
`
	if len(os.Args) > 2 &&
		(os.Args[1] == "-h" ||
			os.Args[1] == "h" ||
			os.Args[1] == "help") {
		fmt.Println(usage)
		os.Exit(0)
	}
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
	fileType := os.Args[1]
	source := os.Args[2]
	var dist string
	if len(os.Args) == 4 {
		dist = os.Args[3]
	} else {
		dist = path.Base(source[:len(source)-len(path.Ext(source))]) +
			"." + fileType
	}
	topic := ParseXML(GetContent(source))
	switch fileType {
	case "xlsx":
		MakeXlsx(topic, dist)
	case "org":
		MakeOrg(topic, dist)
	default:
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
}
