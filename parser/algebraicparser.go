package parser

import (
  "cpl/variable"
  "strconv"
  "strings"
  "errors"
  "fmt"
  "os"
  "bufio"
)

type OpType int

const (
  NULL OpType = iota + 1
  ADD
  SUB
  MUL
  DIV
  MOD
  EXP
  FAC
  AND
  OR
  NOT
  EQU
  GT
  LT
  GTE
  LTE
  NEQ
  EQU2
  GT2
  LT2
  GTE4
  LTE4
)

func AlgebraicParser(expression string, variableMap []map[string]variable.Variable ) (variable.Variable,  error)  {
  inQuotes := false
  parenthCount := 0
  addSubIndex := -1
  mulDivIndex := -1
  expIndex := -1
  facIndex := -1
  andIndex := -1
  orIndex := -1
  notIndex := -1
  equIndex := -1
  optype := NULL
  //currIndex := -1
  //inquotes := false
  expression_arr := []rune(expression)
  if len(expression_arr) == 0 {
    fmt.Println("Expression is Null")
    returnVar := variable.Variable{}
    return returnVar, errors.New("Expression is Null")
  }
  for i := 0; i < len(expression_arr); i++ {
    if expression_arr[i] == '(' || expression_arr[i] == '（' {
      parenthCount += 1
      continue
    } else if expression_arr[i] == ')' || expression_arr[i] == '）' {
      parenthCount -= 1
      continue
    } else if expression_arr[i] == '"' || expression_arr[i] == '\'' || expression_arr[i] == '”' || expression_arr[i] == '“' || expression_arr[i] == '‘' || expression_arr[i] == '’' {
      //fmt.Println("HERE in this one")
      inQuotes = !inQuotes
    } else if inQuotes && expression_arr[i] == '#' {
      i++
      continue
    } else if parenthCount == 0 && !inQuotes {
      switch expression_arr[i] {
      case '+':
        addSubIndex = i
        optype = ADD
        break
      case '-':
        addSubIndex = i
        optype = SUB
        break
      case '*':
        mulDivIndex = i
        optype = MUL
        break
      case '/':
        mulDivIndex = i
        optype = DIV
        break
      case '%':
        mulDivIndex = i
        optype = MOD
        break
      case '^':
        expIndex = i
        optype = EXP
        break
      case '!':
        facIndex = i
        optype = FAC
        break
      case '与':
        andIndex = i
        optype = AND
        break
      case '或':
        orIndex = i
        optype = OR
        break
      case '非':
        notIndex = i
        optype = NOT
        break
      case '=':
        equIndex = i
        optype = EQU
        break
      case '>':
        equIndex = i
        optype = GT
        break
      case '<':
        equIndex = i
        optype = LT
        break
      }
      if i < len(expression_arr)-1 {
        if expression_arr[i] == '等' && expression_arr[i+1] == '于' {
          equIndex = i
          optype = EQU2
        } else if expression_arr[i] == '大' && expression_arr[i+1] == '于' {
          equIndex = i
          optype = GT2
        } else if expression_arr[i] == '小' && expression_arr[i+1] == '于' {
          equIndex = i
          optype = LT2
        } else if i < len(expression_arr)-2 {
          if expression_arr[i] == '不' && expression_arr[i+1] == '等' && expression_arr[i+2] == '于' {
            equIndex = i
            optype = NEQ
          } else if i < len(expression_arr)-3 {
            if expression_arr[i] == '大' && expression_arr[i+1] == '于' && expression_arr[i+2] == '等' && expression_arr[i+3] == '于' {
              equIndex = i
              optype = GTE4
            } else if expression_arr[i] == '小' && expression_arr[i+1] == '于' && expression_arr[i+2] == '等' && expression_arr[i+3] == '于' {
              equIndex = i
              optype = LTE4
            }
          }
        }
      }
    }
  }
  //fmt.Println(expression)
  if orIndex != -1 {
    part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
    if err != nil {
      return part1, err
    }
    part2, err := AlgebraicParser(string(expression_arr[equIndex+1:]),variableMap)
    if err != nil {
      return part2, err
    }
    return part1.Or(part2)
  }
  if andIndex != -1 {
    part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
    if err != nil {
      return part1, err
    }
    part2, err := AlgebraicParser(string(expression_arr[equIndex+1:]),variableMap)
    if err != nil {
      return part2, err
    }
    return part1.And(part2)
  }
  if notIndex != -1 {
    part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
    if err != nil {
      return part1, err
    }
    return part1.Not()
  }
  if equIndex != -1 {
    if optype == EQU {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Eq(part2)
    }
    if optype == EQU2 {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+2:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Eq(part2)
    }
    if optype == NEQ {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+3:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Neq(part2)
    }
    if optype == GT {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Gt(part2)
    }
    if optype == LT {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Lt(part2)
    }
    if optype == GT2 {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+2:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Gt(part2)
    }
    if optype == LT2 {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+2:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Lt(part2)
    }
    if optype == GTE4 {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+4:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Gte(part2)
    }
    if optype == LTE4 {
      part1, err := AlgebraicParser(string(expression_arr[:equIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[equIndex+4:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Lte(part2)
    }
  }
  if addSubIndex != -1 {
    if optype == ADD {
      part1, err := AlgebraicParser(string(expression_arr[:addSubIndex]), variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[addSubIndex+1:]), variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Add(part2)
    } else if optype == SUB {
      part1, err := AlgebraicParser(string(expression_arr[:addSubIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[addSubIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Sub(part2)
    }
  } else if mulDivIndex != -1 {
    if optype == MUL {
      part1, err := AlgebraicParser(string(expression_arr[:mulDivIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[mulDivIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Mul(part2)
    } else if optype == DIV {
      part1, err := AlgebraicParser(string(expression_arr[:mulDivIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[mulDivIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Div(part2)
    } else if optype == MOD {
      part1, err := AlgebraicParser(string(expression_arr[:mulDivIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[mulDivIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Mod(part2)
    }
  } else if expIndex != -1 {
    part1, err := AlgebraicParser(string(expression_arr[:expIndex]),variableMap)
    if err != nil {
      return part1, err
    }
    part2, err := AlgebraicParser(string(expression_arr[expIndex+1:]),variableMap)
    if err != nil {
      return part2, err
    }
    return part1.Exp(part2)
  } else if facIndex != -1 {
    part1, err := AlgebraicParser(string(expression_arr[:expIndex]),variableMap)
    if err != nil {
      return part1, err
    }
    return part1.Fac()
  } else if expression_arr[0] == '(' && expression_arr[len(expression_arr)-1] == ')' {
    return AlgebraicParser(string(expression_arr[1:len(expression_arr)-1]),variableMap)
  } else if (expression_arr[0] == '"' || expression_arr[0] == '\'' || expression_arr[0] == '”' || expression_arr[0] == '“' || expression_arr[0] == '‘' || expression_arr[0] == '’') &&  (expression_arr[len(expression_arr)-1] == '"' || expression_arr[len(expression_arr)-1] == '\'' || expression_arr[len(expression_arr)-1] == '”' || expression_arr[len(expression_arr)-1] == '“' || expression_arr[len(expression_arr)-1] == '‘' || expression_arr[len(expression_arr)-1] == '’') {
    return StringParser(string(expression_arr[1:len(expression_arr)-1]))
  } else {
    return EvaluateAtom(string(expression_arr),variableMap)
  }
  return variable.Variable{}, nil
}

func EvaluateAtom(expression string, variableMap []map[string]variable.Variable) (variable.Variable, error)  {
  returnVar := variable.Variable{}

  ////////////////////
  // ARRAY []
  ////////////////////
  if ( strings.HasSuffix(expression,"]") || strings.HasSuffix(expression,"】") ) && ( strings.HasPrefix(expression,"[") || strings.HasPrefix(expression,"【") ) {
    returnVar.TypeCode = variable.ARRAY
    expr_arr := []rune(expression)
    expr_arr = expr_arr[1:len(expr_arr)-1]
    array_elements_arr := strings.FieldsFunc(string(expr_arr),SplitByCommas)
    returnVar.ArrayVal = nil
    for _, el := range array_elements_arr {
      tempVar, err := AlgebraicParser(el, variableMap)
      if err != nil {
        //fmt.Println("Error Occurred")
        return returnVar, err
      }
      //fmt.Println("Added Element")
      returnVar.ArrayVal = append(returnVar.ArrayVal,tempVar)
    }
    //fmt.Println(len(returnVar.ArrayVal))
    return returnVar, nil
  }





  ////////////////////
  // ARRAYS and STRINGS
  ////////////////////
  if strings.HasSuffix(expression,"]") || strings.HasSuffix(expression,"】") {
    expr_arr := []rune(expression)
    var name []rune
    var i int
    for i = 0; i < len(expr_arr); i++ {
      if expr_arr[i] == '[' || expr_arr[i] == '【' {
        break
      } else {
        name = append(name, expr_arr[i])
      }
    }
    if i >= len(expr_arr) {
      return returnVar, errors.New("No opening [")
    }
    expr_arr = expr_arr[i+1:len(expr_arr)-1]
    arg_arr := strings.Split(string(expr_arr),"][") //TODO Also split by other type of brackets
    for _, vmap := range variableMap {
      if val, exists := vmap[string(name)]; exists {
        if val.TypeCode != variable.STRING && val.TypeCode != variable.ARRAY {
          return returnVar, errors.New("Variable Cannot Be Indexed")
        }
        if val.TypeCode == variable.ARRAY {
          for _, arg := range arg_arr {
            tempVar, err := AlgebraicParser(arg,variableMap)
            if err != nil {
                return returnVar, err
            }
            //fmt.Println("Index Parsed")
            //fmt.Println(tempVar)
            if tempVar.TypeCode == variable.FLOAT {
              //TODO Round to nearest Int
            }
            if tempVar.TypeCode != variable.INT {
              return returnVar, errors.New("Index must evaluate to int or float")
            }
            //fmt.Println("Type Code is integer")
            //fmt.Println(tempVar.IntVal)
            returnVar = val.ArrayVal[tempVar.IntVal]
            return returnVar, nil
          }
        }
        if val.TypeCode ==variable.STRING {
          if len(arg_arr) > 1 {
            return returnVar, errors.New("String only has one dimension")
          }
          tempVar, err := AlgebraicParser(arg_arr[0],variableMap)
          if err != nil {
            return returnVar, err
          }
          if tempVar.TypeCode == variable.FLOAT {
            //TODO Round to nearest Int
          }
          if tempVar.TypeCode != variable.INT {
            return returnVar, errors.New("Index must evaluate to int or float")
          }
          tempSlice := []rune(val.StringVal)
          returnVar.StringVal = string(tempSlice[tempVar.IntVal])
          returnVar.TypeCode = variable.STRING
          return returnVar, nil
        }
      }
    }
  }




  if strings.HasSuffix(expression,")") || strings.HasSuffix(expression,"）") {
    expr_arr := []rune(expression)
    var name []rune
    var i int
    for i = 0; i < len(expr_arr); i++ {
      if expr_arr[i] == '(' || expr_arr[i] == '（' {
        break
      } else {
        name = append(name,expr_arr[i])
      }
    }
    if i >= len(expr_arr) {
      return returnVar, errors.New("Improperly Formed Function Call")
    }
    expr_arr = expr_arr[i+1:len(expr_arr)-1]
    arg_arr := strings.FieldsFunc(string(expr_arr),SplitByCommas)
    ////////////////////////
    // INPUT
    ////////////////////////
    if string(name) == "输入" {
      if len(arg_arr) != 1 {
        return returnVar, errors.New("Incorrect Number of Args")
      }
      tempVar, err := AlgebraicParser(arg_arr[0],variableMap)
      if err != nil {
        return returnVar, err
      }
      inputBuffer := bufio.NewReader(os.Stdin)
      output_str, err := tempVar.ToString()
      if err != nil {
        return returnVar, err
      }
      fmt.Printf(output_str)
      input_text, err := inputBuffer.ReadString('\n')
      if err != nil {
        return returnVar, err
      }
      input_text = strings.Replace(input_text,"\r\n","",-1)
      input_text = strings.Replace(input_text,"\n","",-1)
      float_val, err := strconv.ParseFloat(input_text,64)
      if err == nil {
        returnVar.TypeCode = variable.FLOAT
        returnVar.FloatVal = float_val
        return returnVar, nil
      }
      int_val, err := strconv.ParseInt(input_text,10,64)
      if err == nil {
        returnVar.TypeCode = variable.INT
        returnVar.IntVal = int_val
        return returnVar, nil
      }
      //fmt.Printf("String Type Generated by Input") //DEBUG
      //fmt.Printf(error.Error(err))
      returnVar.TypeCode = variable.STRING
      returnVar.StringVal = input_text
      return returnVar, nil
    }
    for _, vmap := range variableMap {
      if val, exists := vmap[string(name)]; exists {
        if val.TypeCode == variable.FUNC {
          workspace := []map[string]variable.Variable{}
          workspace = append(workspace, map[string]variable.Variable{})
          if len(arg_arr) != len(val.FuncArgs) {
            return returnVar, errors.New("Incorrect Number of Args")
          }
          for j := 0; j < len(val.FuncArgs); j++ {
            tempVar, err := AlgebraicParser(arg_arr[j],variableMap)
            workspace[0][val.FuncArgs[j]] = tempVar
            if err != nil {
              return returnVar, errors.New("Errors evaluating function arg")
            }
          }
          workspace, _, err := ParseScript(val.FuncVal,val.FuncLines,workspace)
          if err != nil {
            return returnVar, err
          }
          return workspace[0]["+返回价值"], nil
        }
      }
    }
    return returnVar, errors.New("Function Name Not Found")
  }
  value, err := strconv.ParseInt(expression,10,64)
  if err != nil {
    value, err := strconv.ParseInt(expression,10,64)
    if err != nil {
      if ( strings.HasPrefix(expression,"\"") || strings.HasPrefix(expression,"“") || strings.HasPrefix(expression,"”") ) && ( strings.HasSuffix(expression,"\"") || strings.HasSuffix(expression,"“") || strings.HasSuffix(expression,"”") ) {
        returnVar.TypeCode = variable.STRING
        returnVar.StringVal = expression[1:len(expression)-1]
        return returnVar, nil
      } else {
        for _, vmap := range variableMap {
          if val, exists := vmap[expression]; exists {
            if val.TypeCode == variable.FUNC {
              workspace := []map[string]variable.Variable{}
              workspace = append(workspace, map[string]variable.Variable{})
              workspace, _, err := ParseScript(val.FuncVal,val.FuncLines,workspace)
              if err != nil {
                return returnVar, err
              }
              return workspace[0]["+返回价值"], nil
            } else {
              return val, nil
            }
          }
        }
      }
    }
    returnVar.TypeCode = variable.FLOAT
    returnVar.IntVal = value
    return returnVar, nil
  }
  returnVar.TypeCode = variable.INT
  returnVar.IntVal = value
  return returnVar, nil
}

func SplitByCommas( r rune ) bool {
  return r == ',' || r == '、'
}
