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

func ParseScript( script []string, workspace []map[string]variable.Variable ) ([]map[string]variable.Variable, int, error) {
  index := 0
  localWorkspace := map[string]variable.Variable{}
  workspace = append(workspace,localWorkspace)
  for index < len(script) {
    if strings.HasPrefix(script[index],"输出") {
      err := Output(script[index],workspace)
      if err != nil {
        return workspace, -1, err
      }
    } else if strings.HasPrefix(script[index],"如果") {
      gap, status, workspace, err := If(script[index:],workspace)
      if err != nil {
        fmt.Println("ERROR FROM IF STATEMENT")
        return workspace, -1, err
      }
      if status == 1 || status == 2 || status == 3 {
        return workspace, status, nil
      }
      //fmt.Println(gap+1)
      index+=gap
    } else if strings.HasPrefix(script[index],"否则") {
      fmt.Println("ERROR FROM ELSE STATEMENT")
      return workspace, -1, errors.New("This can't get used here")
    } else if strings.HasPrefix(script[index],"结束分支") {
      fmt.Println("ERROR FROM END STATEMENT")
      return workspace, -1, errors.New("This can't get used here")
    } else if strings.HasPrefix(script[index],"当") {
      gap, status, workspace, err := While(script[index:],workspace)
      if err != nil {
        return workspace, -1, err
      }
      if status == 3 {
        return workspace, status, nil
      }
      index+=gap+1
    } else if strings.HasPrefix(script[index],"结束循环") {
      return workspace, -1, errors.New("This can't get used here")
    } else if strings.HasPrefix(script[index],"跳出") {
      if len(workspace) > 1 {
        return workspace, -1, errors.New("This can't get used here")
      }
      return workspace, 1, nil
    } else if strings.HasPrefix(script[index],"继续") {
      if len(workspace) > 1 {
        return workspace, -1, errors.New("This can't get used here")
      }
      return workspace, 2, nil
    } else if strings.HasPrefix(script[index],"返回") {
      if len(workspace) <= 1 {
        return workspace, -1, errors.New("This can't get used here")
      }
      workspace, err := Return(script[index], workspace)
      if err != nil {
        return workspace, -1, err
      }
      return workspace, 3, nil
    } else if strings.HasPrefix(script[index],"离去") {
      os.Exit(1)
    } else {
      workspace, err := EvaluateExpression(script[index],workspace)
      if err != nil {
        return workspace, -1, err
      }
    }
    index++
  }
  return workspace[:len(workspace)-1], 0, nil
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

func If( script []string, workspace []map[string]variable.Variable ) (int, int, []map[string]variable.Variable, error) {
  iflevel := 1
  var case_arr []bool
  var codelet_arr [][]string
  var codelet []string
  text_to_parse := strings.TrimPrefix(script[0],"如果")
  evaluated_expression, err := AlgebraicParser(text_to_parse,workspace)
  if err != nil || evaluated_expression.TypeCode != variable.BOOL {
    return 0, -1, workspace, err
  }
  case_arr = append(case_arr,evaluated_expression.BoolVal)
  var index int
  for index = 1; index < len(script); index++ {
    if strings.HasPrefix(script[index],"如果") {
      iflevel++
    } else if strings.HasPrefix(script[index],"结束分支") {
      iflevel--
    } else if strings.HasPrefix(script[index],"否则如果") && iflevel == 1 {
      text_to_parse = strings.TrimPrefix(script[index],"否则如果")
      evaluated_expression, err = AlgebraicParser(text_to_parse,workspace)
      if err != nil || evaluated_expression.TypeCode != variable.BOOL {
        return 0, -1, workspace, err
      }
      case_arr = append(case_arr,evaluated_expression.BoolVal)
      codelet_arr = append(codelet_arr,codelet)
      codelet = nil
      continue
    } else if strings.HasPrefix(script[index],"否则") && iflevel == 1 {
      case_arr = append(case_arr,true)
      codelet_arr = append(codelet_arr,codelet)
      codelet = nil
      continue
    }
    if iflevel == 0 {
      break
    }
    codelet = append(codelet,script[index])
  }
  codelet_arr = append(codelet_arr,codelet)
  if iflevel != 0 {
    return 0, -1, workspace, err
  }
  for i := 0; i < len(case_arr); i++ {
    if case_arr[i] {
      workspace, status, err := ParseScript( codelet_arr[i], workspace )
      return index, status, workspace, err
      break
    }
  }
  return index, 0, workspace, err
}

func While( script []string, workspace []map[string]variable.Variable ) (int, int, []map[string]variable.Variable, error) {
  expression := strings.TrimPrefix(script[0],"当")
  var loop_contents []string
  loop_count := 1
  var index int
  for index = 1; index < len(script); index++ {
    if strings.HasPrefix(script[index],"当") {
      loop_count++
    }
    if strings.HasPrefix(script[index],"结束循环") {
      loop_count--
    }
    if loop_count <= 0 {
      break
    }
    loop_contents = append(loop_contents,script[index])
  }
  if loop_count > 0 {
    return 0, 0, workspace, errors.New("No end loop")
  }
  true_false, err := AlgebraicParser(expression,workspace)
  if err != nil {
    return  0, 0, workspace, err
  }
  if true_false.TypeCode != variable.BOOL {
    return 0, 0, workspace, errors.New("Invalid Expression (Must Evaluate to Boolean)")
  }
  for true_false.BoolVal {
    workspace, status, err := ParseScript(loop_contents,workspace)
    if err != nil {
      return loop_count, 0, workspace, err
    }
    if status == 1 {
      return loop_count, 1, workspace, nil
    }
    if status == 3 {
      return loop_count, 3, workspace, nil
    }
    true_false, err = AlgebraicParser(expression,workspace)
    if err != nil {
      return  0, 0, workspace, err
    }
    if true_false.TypeCode != variable.BOOL {
      return 0, 0, workspace, errors.New("Invalid Expression")
    }
  }
  return len(loop_contents), 0, workspace, nil
}

func EvaluateExpression(text string, workspace []map[string]variable.Variable) ([]map[string]variable.Variable, error) {
  text_arr := []rune(text)
  in_quotes := false
  storage_statement := false
  for i := 0; i < len(text_arr); i++ {
    if text_arr[i] == '"' || text_arr[i] == '\'' || text_arr[i] == '”' || text_arr[i] == '“' || text_arr[i] == '‘' || text_arr[i] == '’' {
      in_quotes = !in_quotes
    } else if in_quotes && text_arr[i] == '#' {
      i++
    } else if !in_quotes && text_arr[i] == '《' {
      storage_statement = true
      place_to_store := string(text_arr[:i])
      expression := string(text_arr[i+1:])
      if strings.Contains(place_to_store,"+!@#$%^&*()_-=][{}|/?><.,") {
        return workspace, errors.New("Invalid Variable Name")
      }
      temp, err := AlgebraicParser(expression,workspace)
      if err != nil {
        return workspace, err
      }
      found_var_name := false
      for _, vmap := range workspace {
        if _, exists := vmap[place_to_store]; exists {
          vmap[place_to_store] = temp
          found_var_name = true
          break
        }
      }
      if !found_var_name {
        workspace[len(workspace)-1][place_to_store] = temp
      }
    } else if !in_quotes && text_arr[i] == '》' {
      storage_statement = true
      place_to_store := string(text_arr[i+1:])
      expression := string(text_arr[:i])
      if strings.Contains(place_to_store,"+`!@#$%^&*()_-=][{}|/?><.,\\\"'") {
        return workspace, errors.New("Invalid Variable Name")
      }
      temp, err := AlgebraicParser(expression,workspace)
      if err != nil {
        return workspace, err
      }
      found_var_name := false
      for _, vmap := range workspace {
        if _, exists := vmap[place_to_store]; exists {
          vmap[place_to_store] = temp
          found_var_name = true
          break
        }
      }
      if !found_var_name {
        workspace[len(workspace)-1][place_to_store] = temp
      }
    }
  }
  if !storage_statement {
    _, err := AlgebraicParser(text,workspace)
    if err != nil {
      return workspace, err
    }
  }
  return workspace, nil
}

func Return(text string, workspace []map[string]variable.Variable) ([]map[string]variable.Variable, error) {
  expression := strings.TrimPrefix(text,"返回")
  if len(expression) == 0 {
    return workspace, nil
  }
  variable, err := AlgebraicParser(expression, workspace)
  if err != nil {
    return workspace, err
  }
  workspace[0]["+返回价值"] = variable
  return workspace, nil
}
