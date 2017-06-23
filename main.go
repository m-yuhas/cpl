package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "cpl/variable"
    "cpl/parser"
    //"ioutil"
)

//var variableMap = map[string]variable.Variable{}

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("欢迎中华电脑语言第0.2版本!\n©2017 － 迈克尔 余哈斯")
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
      fmt.Println("错误：不能开文件")
      panic(err)
    }
    defer f.Close()
    var lines []string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
      lines = append(lines, strings.TrimSpace(scanner.Text()))
    }
    if err := scanner.Err(); err != nil {
      fmt.Println(os.Stderr,err)
    }
    variableMap := []map[string]variable.Variable{}
    variableMap = append(variableMap, map[string]variable.Variable{})

    for i := 0; i < len(lines); i++ {
      lines[i] = strings.TrimSpace(lines[i]])
      if line.HasPrefix("函数") {
        line = strings.TrimPrefix(lines[i],"函数")
        line = strings.TrimSpace(lines[i])
        nameAndArgs = strings.Split(line,"要")
        name = strings.TrimSpace(nameAndArgs[0])
        arglist = strings.Split(strings.TrimSpace(line),',')
        variableMap[0][name] = variable.Variable{}
        variableMap[0][name].SetType(10)
        variableMap[0][name].FuncVal = []string{}
        i++
        for i < len(lines) {
          if strings.TrimeSpace(lines[i]) == "结束函数" {
            break
          }
          variableMap[0][name].FuncVal = append(variableMap[0][name].FuncVal,lines[i])
          i++
        }
        for _, arg = range arglist {
          variableMap[0][name].FuncArgs = append(variableMap[0][name].FuncArgs,arg)
        }
      }
    }
    parser.ParseScript(lines,variableMap)
/*
    for _, line := range lines {
      fmt.Println(line)
    }
    */
  }






/*
  testObj := variable.Variable{}
  testObj.SetType(1)
  testObj.SetValue(1000)
  variableMap["A"] = testObj
  fmt.Println(variableMap["A"])
  */

}
