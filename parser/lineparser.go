package parser

import (
  "fmt"
  "strings"
  "variable"
  "os"
)

func ParseLine( line string ) {
  variableMap := []map[string]variable.Variable{}
  line = strings.Trim(line," \t\r\n")
  if strings.HasPrefix(line,"注意") {
    fmt.Println("comment activated")
  } else if strings.HasPrefix(line,"输出") {
    fmt.Println("output activated")
  } else if strings.HasPrefix(line,"如果") {
    fmt.Println("if activated")
  } else if strings.HasPrefix(line,"否则") {
    fmt.Println("else activated")
  } else if strings.HasPrefix(line,"结束支") {
    fmt.Println("end branch")
  } else if strings.HasPrefix(line,"从") {
    fmt.Println("for loop")
  } else if strings.HasPrefix(line,"当") {
    fmt.Println("while loop")
  } else if strings.HasPrefix(line,"结束圈") {
    fmt.Println("end loop")
  } else if strings.HasPrefix(line,"跳出") {
    fmt.Println("break")
  } else if strings.HasPrefix(line,"函数") {
    fmt.Println("function")
  } else if strings.HasPrefix(line,"结束函数") {
    fmt.Println("end function")
  } else if strings.HasPrefix(line,"离去") {
    os.Exit(1)
  } else if strings.Contains(line,"=") {
    AlgebraicParser(strings.SplitN(line,"=",-1)[1],variableMap)
  } else {
    tmp := variable.Variable{}
    tmp = AlgebraicParser(line,variableMap)
    fmt.Println(tmp.IntVal)

  }
}
