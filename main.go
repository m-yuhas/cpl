/*
main.go - Entry Point for Code Execution
(C) 2017 Michael Yuhas
*/

/*
Decare this as part of package 'main' so it has somewhere to go
*/
package main

/*
Use the following packages:
fmt - printing to console
bufio - io functions for reading script file
os - os utilities for files
strings - string processing utilities
cpl/variable - variable class
cpl/parser - parser class
cpl/messages - error, warning, and info messages
*/
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "cpl/variable"
    "cpl/parser"
    "cpl/messages"
)

/*
main() - main entry point
Arguments: None
Returns: None
*/
func main() {
  //Check if user wants to start program in terminal mode or execute a script
  if len(os.Args) <= 1 {
    fmt.Println(messages.CLIHeaderText)
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
      lines[i] = strings.TrimSpace(lines[i])
      if strings.HasPrefix(lines[i],"函数") {
        line := strings.TrimPrefix(lines[i],"函数")
        line = strings.TrimSpace(line)
        nameAndArgs := strings.Split(line,"要")
        name := strings.TrimSpace(nameAndArgs[0])
        arglist := strings.Split(strings.TrimSpace(line),string(','))
        tempVar := variable.Variable{}
        tempVar.TypeCode = 10
        i++
        tempFuncVal := []string{}
        for i < len(lines) {
          if strings.TrimSpace(lines[i]) == "结束函数" {
            break
          }
          tempFuncVal = append(tempFuncVal,lines[i])
          i++
        }
        tempVar.FuncVal = tempFuncVal
        tempArgList := []string{}
        for _, arg := range arglist {
          tempArgList = append(tempArgList,arg)
        }
        tempVar.FuncArgs = tempArgList
        variableMap[0][name] = tempVar
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
