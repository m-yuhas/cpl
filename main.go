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
  // Check if user wants to start program in terminal mode or execute a script
  if len(os.Args) <= 1 {
    fmt.Println(messages.CLIHeaderText)
    cli_main()
  } else {
    f, err := os.Open(os.Args[1])
    if err != nil {
      fmt.Println("错误：不能开文件")
      panic(err)
    }
    defer f.Close() //TODO: make sure that this is the right way to open/close files
    var lines []string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
      lines = append(lines, strings.TrimSpace(scanner.Text()))
    }
    if err := scanner.Err(); err != nil {
      fmt.Println(os.Stderr,err)
    }


    lines = find_comments(lines)

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
}

/*
cli_main() - top level function to handle the CLI
Arguments: None
Returns: None
*/
func cli_main() {
  // Loop continuously and execute commands provided by user
  for {
    inputBuffer := bufio.NewReader(os.Stdin)
    fmt.Printf(">>>")
    userInput, err := inputBuffer.ReadString('\n')
    // Handle error if one occurs
    if err != nil {
      fmt.Println(messages.CLIInputError)
      os.Exit(1)
    }
    userInput = strings.Replace(userInput,"\r\n","",-1)
    parser.ParseLine(userInput)
  }
}

/*
find_comments - finds and removes comments from an array of lines
Arguments:
  str_arr - Array of strings
Returns:
  []string - Modified Array of strings with comments removed
*/
func find_comments( str_arr []string ) []string {
    var output_str_arr []string
    // Loop through array and search for commented blocks
    for i := 0; i < len(str_arr); i++ {
        // Check for start of commented block
        if strings.HasPrefix(str_arr[i],"注释：") || strings.HasPrefix(str_arr[i],"注释:") {
          var j int
          // Loop through lines to search for end of commented block
          for j = i+1; j < len(str_arr); j++ {
            // If the end is found stop the loop
            if strings.HasSuffix(str_arr[j],"结束注释") {
              break
            }
          }
          // If the end was found, update the line counter to the current position
          if j < len(str_arr) {
            i = j
          }
        } else if strings.HasSuffix(str_arr[i],"结束注释") {
          fmt.Println(messages.CommentEndWithoutStart)
          os.Exit(0)
        } else {
          output_str_arr = append(output_str_arr,str_arr[i])
        }
    }
    return output_str_arr
}

/*
strip_whitespace - remove unnecessary whitespace from string array
Arguments:
  str_arr - Array of strings
Returns:
  []string - Array of strings with whitespace removed
*/
func strip_whitespace( str_arr []string ) []string {
  var output_str_arr []string
  // Loop through array and remove unneeded spaces line by line
  for i := 0; i < len(str_arr); i++ {
    input_line := []rune(str_arr[j])
    strip_whitespace()
  }
}

func strip_whitespace( input_rune_arr []rune ) []rune {
  var output_rune_arr []rune
  for i := 0; i < len(input_rune_arr); i++ {
    if input_line[j] == '#' {
      output_rune_arr = append(output_rune_arr,input_rune_arr[i])
      i++
      if input_line[i] == '(' {
        output_rune_arr = append(output_rune_arr,input_rune_arr[i])
        var temp_rune_arr []rune
        parenth_count = 1
        for i < len(input_rune_arr) &&  {
          i++
          temp_rune_arr == append(temp_rune_arr,input_rune_arr[i])
          if input_rune_arr[i] == '(' {
            parenth_count++
          }
          if input_rune_arr[i] == ')' {
            parenth_count--
          }
          if parenth_count == 0 {
            break
          }
          temp_rune_arr == append(temp_rune_arr,input_rune_arr[i])
        }
        if i >= len(input_rune_arr) {
          fmt.Println(messages.)
          os.Exit(0)
        }
        output_rune_arr == append(output_rune_arr,strip_whitespace(temp_rune_arr))
        output_rune_arr == append(output_rune_arr,')')
      }
    } else if input_rune_arr[i] == '"' || input_rune_arr[i] == '\'' || input_rune_arr[i] == '”' || input_rune_arr[i] == '“' || input_rune_arr[i] == '‘' || input_rune_arr[i] == '’' {
      in_quotes = !in_quotes
      output_rune_arr = append(output_rune_arr,input_rune_arr[i])
    } else if input_rune_arr[i] == ' ' || input_rune_arr[i] == '\t' && in_quotes {
      continue
    } else {
      output_rune_arr = append(output_rune_arr,input_rune_arr[i])
    }
  }
}
