package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "variable"
    "parser"
    //"ioutil"
)

var variableMap = map[string]variable.Variable{}

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("欢迎中华电脑语言第0.2版本!\n©2017 － 麦克尔 余哈斯")
    //var variableMap = make(map[string]variable.Variable)
    for {
      inputBuffer := bufio.NewReader(os.Stdin)
      fmt.Printf(">>>")
      userInput, _ := inputBuffer.ReadString('\n')
      userInput = strings.Replace(userInput,"\r\n","",-1)
      parser.ParseLine(userInput)
    }
  } else {
    f, err := os.Open(os.Args[1])
    if err != nil {
      fmt.Println("Error opening file")
      panic(err)
    }
    defer f.Close()
    var lines []string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
      lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
      fmt.Println(os.Stderr,err)
    }
    for _, line := range lines {
      fmt.Println(line)
    }
  }






/*
  testObj := variable.Variable{}
  testObj.SetType(1)
  testObj.SetValue(1000)
  variableMap["A"] = testObj
  fmt.Println(variableMap["A"])
  */

}
