package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "variable"
    "parser"
)

/*
type Variable struct {
  type_code uint8
  intVal int64
}

func (v *Variable) setType(type_code uint8) {
  v.type_code = type_code
}

func (v *Variable) setValue(value int64) {
  v.intVal = value
}
*/

func main() {
  fmt.Println("欢迎中华电脑语言第0.2版本!\n©2017 － 麦克尔 余哈斯")
  var variableMap = make(map[string]variable.Variable)
  inputBuffer := bufio.NewReader(os.Stdin)
  fmt.Printf(">>>")
  userInput, _ := inputBuffer.ReadString('\n')
  userInput = strings.Replace(userInput,"\r\n","",-1)
  parser.ParseLine(userInput)






  testObj := variable.Variable{}
  testObj.SetType(1)
  testObj.SetValue(1000)
  variableMap["A"] = testObj
  fmt.Println(variableMap["A"])

}
