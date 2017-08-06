package parser

import (
  "fmt"
  "strings"
  "cpl/variable"
  "os"
  "errors"
  "cpl/messages"
  //"strconv"
)

func ParseScript( script []string, workspace []map[string]variable.Variable ) ([]map[string]variable.Variable, error) {
  iflevel := 0
  index := 0
  localWorkspace := map[string]variable.Variable{}
  workspace = append(workspace,localWorkspace)
  for index < len(script) {
    if strings.HasPrefix(script[index],"输出") {
      Output(script[index],workspace)
    } else if strings.HasPrefix(script[index],"如果") {
      text := strings.TrimPrefix(script[index],"如果")
      trueFalse, err := BooleanParser(text,workspace)
      if err != nil {
        fmt.Println(err)
      }
      iflevel++
      starting_iflevel := iflevel
      if !trueFalse.BoolVal {
        index++
        for index < len(script) {
          if strings.HasPrefix(script[index],"如果") {
            iflevel++
          } else if strings.HasPrefix(script[index],"否则") && iflevel == starting_iflevel {
            break
          } else if strings.HasPrefix(script[index],"结束支") {
            iflevel--
          }
          if iflevel == starting_iflevel - 1 {
            break
          }
          index++
        }
      }
    } else if strings.HasPrefix(script[index],"否则") {
      if iflevel < 0 {
        fmt.Println("ERROR 否则")
        os.Exit(1)
      }
      index++
      starting_iflevel := iflevel
      for index < len(script) {
        if strings.HasPrefix(script[index],"如果") {
          iflevel++
        } else if strings.HasPrefix(script[index],"结束支") {
          iflevel--
        }
        if iflevel == starting_iflevel - 1 {
          break
        }
        index++
      }
    } else if strings.HasPrefix(script[index],"结束支") {
      if iflevel < 0 {
        fmt.Println("ERROR 否则")
        os.Exit(1)
      }
    } else if strings.HasPrefix(script[index],"从") {
      //TODO: Check for nested loops
      expression := strings.TrimPrefix(script[index],"从")
      expressionArray := strings.Split(expression,"直到")
      initCondArray := strings.Split(expressionArray[0],"=")
      initCondArray[0] = strings.TrimSpace(initCondArray[0])
      initCondArray[1] = strings.TrimSpace(initCondArray[1])
      var pos2 int
      for pos, vmap := range workspace {
        if _, exists := vmap[initCondArray[0]]; exists {
          var err error
          vmap[initCondArray[0]], err = AlgebraicParser(initCondArray[1],workspace)
          if err != nil {
            fmt.Println(err)
          }
          pos2=pos
          break
        }
      }
      var loopContents []string
      index++
      for index < len(script) {
        if strings.HasPrefix(script[index],"结束圈") {
          break
        }
        loopContents = append(loopContents,script[index])
        index++
      }
      var1 := workspace[pos2][initCondArray[0]]
      temp0, err := AlgebraicParser(strings.TrimSpace(expressionArray[1]),workspace)
      if err != nil {
        fmt.Println(err)
      }
      eval, err := var1.Eq(temp0)
      if err != nil {
        fmt.Println(err)
      }
      for !eval.BoolVal {
        //fmt.Println(variableMap)
        //fmt.Println(variableMap[pos2][initCondArray[0]].IntVal)
        //fmt.Println("HERE")
        workspace, _ = ParseScript(loopContents,workspace)
        //fmt.Println()
        var1 = workspace[pos2][initCondArray[0]]

        temp0, err := AlgebraicParser(strings.TrimSpace(expressionArray[1]),workspace)
        if err != nil {
          fmt.Println(err)
        }
        eval, err = var1.Eq(temp0)
        if err != nil {
          fmt.Println(err)
        }
      }
      //index++
    } else if strings.HasPrefix(script[index],"当") {
      expression := strings.TrimPrefix(script[index],"当")
      expression = strings.TrimSpace(expression)
      var loopContents []string
      index++
      for index < len(script) {
        if strings.HasPrefix(script[index],"结束圈") {
          break
        }
        loopContents = append(loopContents,script[index])
        index++
      }
      var1, err := BooleanParser(expression,workspace)
      if err != nil {
        fmt.Println(err)
      }
      for var1.BoolVal {
        workspace, _ = ParseScript(loopContents,workspace)
        var1, err = BooleanParser(expression,workspace)
      }
      //fmt.Println("while loop")
    } else if strings.HasPrefix(script[index],"结束圈") {
      fmt.Println("end loop")
    } else if strings.HasPrefix(script[index],"跳出") {
      fmt.Println("break")
    } else if strings.HasPrefix(script[index],"函数") {
      fmt.Println("function")
    } else if strings.HasPrefix(script[index],"结束函数") {
      fmt.Println("end function")
    } else if strings.HasPrefix(script[index],"离去") {
      os.Exit(1)
    } else if strings.Contains(script[index],"=") {
      expressionArray := strings.SplitN(script[index],"=",-1)
      if len(expressionArray) > 2 {
        fmt.Println("ERROR too many =")
        continue
      }
      expressionArray[0] = strings.TrimSpace(expressionArray[0])
      expressionArray[1] = strings.TrimSpace(expressionArray[1])
      if strings.Contains(expressionArray[0],"?/\\][{}()*&^%$#@!~`]") {
        fmt.Println("ERROR invalid character")
      }
      modified := false
      for _, vmap := range workspace {
        if _, exists := vmap[expressionArray[0]]; exists {
          var err error
          vmap[expressionArray[0]], err = AlgebraicParser(strings.SplitN(script[index],"=",-1)[1],workspace)
          if err != nil {
            fmt.Println(err)
          }
          modified = true
          break
        }
      }
      if !modified {
        var err error
        workspace[len(workspace)-1][expressionArray[0]], err = AlgebraicParser(strings.SplitN(script[index],"=",-1)[1],workspace)
        if err != nil {
          fmt.Println(err)
        }
        //fmt.Println("Setting "+expressionArray[0]+"to"+string(AlgebraicParser(strings.SplitN(script[index],"=",-1)[1],variableMap).IntVal))
      }
    } else {
/*
      tmp := variable.Variable{}
      tmp = AlgebraicParser(script[index])
      fmt.Println(tmp.IntVal)*/
    }
    index++
  }
  return workspace[:len(workspace)-1], nil
}

func Output( text string, workspace []map[string]variable.Variable ) error {
  text_to_parse := []rune(strings.TrimPrefix(text,"输出"))
  if ( text_to_parse[0] == '(' || text_to_parse[0] == '（' ) && ( text_to_parse[len(text_to_parse)-1] == ')' || text_to_parse[len(text_to_parse)-1] == '）' ) {
    text_to_parse = text_to_parse[1:len(text_to_parse)-1]
  } else {
    return errors.New(messages.OutputCommandSyntaxError)
  }
  evaluated_expression, err := AlgebraicParser(string(text_to_parse),workspace)
  if err != nil {
    return err //TODO Find way to append current line to error
  }
  output_str, err := evaluated_expression.ToString()
  if err != nil {
    return err  // TODO see above
  }
  fmt.Println(output_str)
  return nil
}
