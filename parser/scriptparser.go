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
  index := 0
  localWorkspace := map[string]variable.Variable{}
  workspace = append(workspace,localWorkspace)
  for index < len(script) {
    if strings.HasPrefix(script[index],"输出") {
      err := Output(script[index],workspace)
      if err != nil {
        return workspace, err
      }
    } else if strings.HasPrefix(script[index],"如果") {
      gap, workspace, err := If(script[index:],workspace)
      if err != nil {
        fmt.Println("ERROR FROM IF STATEMENT")
        return workspace, err
      }
      //fmt.Println(gap+1)
      index+=gap
    } else if strings.HasPrefix(script[index],"否则") {
      fmt.Println("ERROR FROM ELSE STATEMENT")
      return workspace, errors.New("This can't get used here")
    } else if strings.HasPrefix(script[index],"结束分支") {
      fmt.Println("ERROR FROM END STATEMENT")
      return workspace, errors.New("This can't get used here")
    } else if strings.HasPrefix(script[index],"当") {
      gap, workspace, err := While(script[index:],workspace)
      if err != nil {
        return workspace, err
      }
      index+=gap+1
    } else if strings.HasPrefix(script[index],"结束循环") {
      return workspace, errors.New("This can't get used here")
    } else if strings.HasPrefix(script[index],"跳出") {
      if len(workspace) > 1 {
        return workspace, errors.New("This can't get used here")
      }
      return workspace, nil
    } else if strings.HasPrefix(script[index],"离去") {
      os.Exit(1)
    } else {
      workspace, err := StoreVariable(script[index],workspace)
      if err != nil {
        return workspace, err
      }
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

func If( script []string, workspace []map[string]variable.Variable ) (int, []map[string]variable.Variable, error) {
  iflevel := 1
  var case_arr []bool
  var codelet_arr [][]string
  var codelet []string
  text_to_parse := strings.TrimPrefix(script[0],"如果")
  evaluated_expression, err := AlgebraicParser(text_to_parse,workspace)
  if err != nil || evaluated_expression.TypeCode != 1 {
    return 0, workspace, err
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
      if err != nil || evaluated_expression.TypeCode != 1 {
        return 0, workspace, err
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
    return 0, workspace, err
  }
  for i := 0; i < len(case_arr); i++ {
    if case_arr[i] {
      workspace, err := ParseScript( codelet_arr[i], workspace )
      return index, workspace, err
      break
    }
  }
  return index, workspace, err
}

func While( script []string, workspace []map[string]variable.Variable ) (int, []map[string]variable.Variable, error) {
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
    return 0, workspace, errors.New("No end loop")
  }
  true_false, err := AlgebraicParser(expression,workspace)
  if err != nil {
    return  0, workspace, err
  }
  if true_false.TypeCode != 1 {
    return 0, workspace, errors.New("Invalid Expression (Must Evaluate to Boolean)")
  }
  for true_false.BoolVal {
    workspace, err = ParseScript(loop_contents,workspace)
    if err != nil {
      return loop_count, workspace, err
    }
    true_false, err = AlgebraicParser(expression,workspace)
    if err != nil {
      return  0, workspace, err
    }
    if true_false.TypeCode != 1 {
      return 0, workspace, errors.New("Invalid Expression")
    }
  }
  return len(loop_contents), workspace, nil
}

func StoreVariable(text string, workspace []map[string]variable.Variable) ([]map[string]variable.Variable, error) {
  text_arr := []rune(text)
  in_quotes := false
  for i := 0; i < len(text_arr); i++ {
    if text_arr[i] == '"' || text_arr[i] == '\'' || text_arr[i] == '”' || text_arr[i] == '“' || text_arr[i] == '‘' || text_arr[i] == '’' {
      in_quotes = !in_quotes
    } else if in_quotes && text_arr[i] == '#' {
      i++
    } else if !in_quotes && text_arr[i] == '《' {
      place_to_store := string(text_arr[:i])
      expression := string(text_arr[i+1:])
      if strings.Contains(place_to_store,"+!@#$%^&*()_-=][{}|/?><.,") {
        return workspace, errors.New("Invalid Variable Name")
      }
      temp, err := AlgebraicParser(expression,workspace)
      if err != nil {
        return workspace, nil
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
      place_to_store := string(text_arr[i+1:])
      expression := string(text_arr[:i])
      if strings.Contains(place_to_store,"+`!@#$%^&*()_-=][{}|/?><.,\\\"'") {
        return workspace, errors.New("Invalid Variable Name")
      }
      temp, err := AlgebraicParser(expression,workspace)
      if err != nil {
        return workspace, nil
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
  return workspace, nil
}
