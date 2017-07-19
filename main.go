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
    "errors"
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

    workspace := []map[string]variable.Variable{}
    workspace = append(workspace, map[string]variable.Variable{})
    lines = strip_whitespace(lines)
    lines = find_comments(lines)
    lines, workspace[0], err = find_functions(lines,workspace[0])
    if err != nil {
      fmt.Println(err)
      os.Exit(0)
    }
    if strings.HasPrefix(lines[0],"#") {
      lines = append(lines[:0],lines[1:]...)
    }

    for i := 0; i < len(lines); i++ {
      fmt.Println(lines[i])
    }

    parser.ParseScript(lines,workspace)
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
    input_line := []rune(str_arr[i])
    in_quotes := false
    var output_line []rune
    // Loop through characters in line and remove whitespace
    for j := 0; j < len(input_line); j++ {
      // Cases for in_quotes and quotation preceded by an escape character
      if input_line[j] == '#' {
        output_line = append(output_line,input_line[j])
        j++
        // Check that the '#' is not the final character in a line
        if j >= len(input_line) {
          fmt.Println(messages.LineEndsWithPoundSign)
          os.Exit(0)
        }
        output_line = append(output_line,input_line[j])
      } else if input_line[j] == '"' || input_line[j] == '\'' || input_line[j] == '”' || input_line[j] == '“' || input_line[j] == '‘' || input_line[j] == '’' {
        in_quotes = !in_quotes
        output_line = append(output_line,input_line[j])
      } else if ( input_line[j] == ' ' || input_line[j] == '\t' ) && !in_quotes {
        continue
      } else {
        output_line = append(output_line,input_line[j])
      }
    }
    if string(output_line) != "" {
      output_str_arr = append(output_str_arr,string(output_line))
    }
  }
  return output_str_arr
}

/*
find_functions - find functions, store them to the workspace, and remove from line array
Arguments:
  str_arr - Array of strings
  map[string]variable.Variable - Workspace to which to add functions
Returns:
  []string - Modified Array of strings with functions removed
  map[string]variable.Variable - Workspace with functions added
  error - Fires if an error occurs while parsing the array
*/
func find_functions( str_arr []string, workspace map[string]variable.Variable ) ( []string, map[string]variable.Variable, error ) {
  var out_str_arr []string
  for i := 0; i < len(str_arr); i++ {
    if strings.HasPrefix(str_arr[i],"函数") {
      func_definition := strings.TrimPrefix(str_arr[i],"函数")
      name_and_args := strings.Split(func_definition,"要")
      if len(name_and_args) > 2 {
        return str_arr, workspace, errors.New(messages.InvalidFunctionDeclaration) //TODO include line number in error message
      } //TODO Check args for invalid characters
      arg_list := strings.Split(name_and_args[1],string(',')) //TODO make this work with the other kind of comma
      new_function := variable.Variable{}
      new_function.TypeCode = 10
      i++
      function_content := []string{}
      for i < len(str_arr) {
        if str_arr[i] == "结束函数" {
          break
        }
        if strings.HasPrefix(str_arr[i],"函数") {
          return str_arr, workspace, errors.New(messages.FunctionWithinFunction)
        }
        function_content = append(function_content,str_arr[i])
        i++
      }
      if i >= len(str_arr) {
        return str_arr, workspace, errors.New(messages.EndFunctionNotFound)
      }
      new_function.FuncVal = function_content
      new_function.FuncArgs = arg_list
      if _, ok := workspace[name_and_args[0]]; ok {
        return str_arr, workspace, errors.New(messages.DuplicateName) //TODO include function name in error message
      }
      workspace[name_and_args[0]] = new_function
    } else {
      out_str_arr = append(out_str_arr,str_arr[i])
    }
  }
  return out_str_arr, workspace, nil
}

/*
find_classes - find functions, store them to the workspace, and remove from line array
Arguments:
  str_arr - Array of strings
  map[string]variable.Variable - Workspace to which to add classes
Returns:
  []string - Modified Array of strings with classes removed
  map[string]variable.Variable - Workspace with classes added
  error - Fires if an error occurs while parsing the array
*/
func find_classes( str_arr []string, workspace map[string]variable.Variable ) ( []string, map[string]variable.Variable, error ) {
  var out_str_arr []string
  for i := 0; i < len(str_arr); i++ {
    if strings.HasPrefix(str_arr[i],"类") {
      class_definition := strings.TrimPrefix(str_arr[i],"类")
      name_and_super_class := strings.Split(class_definition,"是")
      if len(name_and_super_class) > 2 {
        return str_arr, workspace, errors.New(messages.InvalidClassDeclaration) //TODO include line number in error message
      } //TODO Check args for invalid characters
      new_class := variable.Variable{}
      new_class.TypeCode = 11
      i++
      class_content := []string{}
      for i < len(str_arr) {
        if str_arr[i] == "结束类" {
          break
        }
        if strings.HasPrefix(str_arr[i],"类") {
          return str_arr, workspace, errors.New(messages.ClassWithinClass)
        }
        class_content = append(class_content,str_arr[i])
        i++
      }
      if i >= len(str_arr) {
        return out_str_arr, workspace, errors.New(messages.EndClassNotFound)
      }
      /*
      new_class.ClassVal = content
      new_function.FuncArgs = name_and_args[1]
      if workspace[name_and_args[0]] != nil {
        return str_arr, workspace, errors.New(messges.DuplicateName) //TODO include function name in error message
      }
      workspace[name_and_args[0]] = new_function
      */
    }
  }
  return str_arr, workspace, nil
}
