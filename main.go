package main

import (
  "os"
  "fmt"
)

// xmindファイルをコマンドライン引数で指定されたファイル形式に変換する
func main() {
  usage :=
`USAGE: convertXmind format source dist
DESCRIPTION:
  convert XMind 8 file to given format.

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
  if len(os.Args) != 4 {
    fmt.Fprint(os.Stderr,usage)
    os.Exit(1)
  }
  fileType := os.Args[1]
  source := os.Args[2]
  dist := os.Args[3]
  topic := ParseXML(GetContent(source))
  switch fileType {
  case "xlsx":
    MakeXlsx(topic,dist)
  case "org":
    MakeOrg(topic,dist)
  default:
    fmt.Fprint(os.Stderr,usage)
    os.Exit(1)
  }
}
