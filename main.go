package main

import (
  "os"
  "fmt"
)


func main() {
  usage := "Usage: convertXmind [xlsx|org] source dist\n"
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
