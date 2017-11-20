package parser

import (
  "cpl/variable"
  "errors"
  //"strconv"
  //"strings"
  //"fmt"
)

func StringParser( expression string ) ( variable.Variable, error ) {
  returnVar := variable.Variable{}
  expression_arr := []rune(expression)
  var output_arr []rune
  for i := 0; i < len(expression_arr); i++ {
    if expression_arr[i] == '"' || expression_arr[i] == '\'' || expression_arr[i] == '”' || expression_arr[i] == '“' || expression_arr[i] == '‘' || expression_arr[i] == '’' {
      return returnVar, errors.New("错误：变量有不对的类")
    } else if expression_arr[i] == '#' {
      if i >= len(expression_arr) {
        return returnVar, errors.New("Fill in Later")
      }
      i++
      if expression_arr[i] == '#' || expression_arr[i] == '"' || expression_arr[i] == '\'' || expression_arr[i] == '”' || expression_arr[i] == '“' || expression_arr[i] == '‘' || expression_arr[i] == '’' {
        output_arr = append(output_arr,expression_arr[i])
      } else if expression_arr[i] == '换' && i <= len(expression_arr) {
        i++
        if expression_arr[i] == '行' {
          output_arr = append(output_arr,'\n')
        }
      }
    } else {
      output_arr = append(output_arr,expression_arr[i])
    }
  }
  returnVar.TypeCode = variable.STRING
  returnVar.StringVal = string(output_arr)
  return returnVar, nil
}
