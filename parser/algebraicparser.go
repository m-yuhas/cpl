package parser

import (
  "cpl/variable"
  "strconv"
  "strings"
  //"fmt"
)

type OpType int

const (
  ADD OpType = 1 << iota
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

func AlgebraicParser(expression string, variableMap []map[string]variable.Variable ) (variable.Variable, error)  {
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
  optype := -1
  currIndex := -1
  inquotes := false
  expression_arr := []rune(expression)
  for i := 0; i < len(expression_arr); i++ {
    if expression_arr[i] == '(' || expression_arr[i] == '（' {
      parenthCount += 1
      continue
    } else if expression_arr[i] == ')' || expression_arr[i] == '）' {
      parenthCount -= 1
      continue
    } else if input_line[i] == '"' || input_line[j] == '\'' || input_line[j] == '”' || input_line[j] == '“' || input_line[j] == '‘' || input_line[j] == '’' {
      inQuotes = !inQuotes
    } else if inQuotes && input_line[i] == '#' {
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
              opytpe = LTE4
            }
          }
        }
      }
    }
  }
  //fmt.Println(expression)
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
      part1, err := AlgebraicParser(string(expression_arr[:addSubIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[addSubIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Mul(part2)
    } else if optype == DIV {
      part1, err := AlgebraicParser(string(expression_arr[:addSubIndex]),variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(string(expression_arr[addSubIndex+1:]),variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Div(part2)
    } else if optype == 5 {
      part1, err := AlgebraicParser(firsthalf,variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(lasthalf,variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Mod(part2)
    }
  } else if expIndex != -1 {
    firsthalf := strings.TrimSpace(expression[:currIndex])
    lasthalf := strings.TrimSpace(expression[currIndex+1:])
    part1, err := AlgebraicParser(firsthalf,variableMap)
    if err != nil {
      return part1, err
    }
    part2, err := AlgebraicParser(lasthalf,variableMap)
    if err != nil {
      return part2, err
    }
    return part1.Exp(part2)
  } else if expression[0] == '(' && expression[len(expression)-1] == ')' {
    return AlgebraicParser(strings.TrimSpace(expression[1:len(expression)-1]),variableMap)
  } else {
    return EvaluateAtom(expression,variableMap)
  }
  return variable.Variable{}, nil
}

func EvaluateAtom(expression string, variableMap []map[string]variable.Variable) (variable.Variable, error)  {
  returnVar := variable.Variable{}
  value, err := strconv.ParseInt(expression,10,64)
  if err != nil {
    value, err := strconv.ParseInt(expression,10,64)
    if err != nil {
      if ( strings.HasPrefix(expression,"\"") || strings.HasPrefix(expression,"“") || strings.HasPrefix(expression,"”") ) && ( strings.HasSuffix(expression,"\"") || strings.HasSuffix(expression,"“") || strings.HasSuffix(expression,"”") ) {
        returnVar.TypeCode = 4
        returnVar.StringVal = expression[1:len(expression)-1]
        return returnVar, nil
      } else {
        for _, vmap := range variableMap {
          if val, exists := vmap[expression]; exists {
            return val, nil
          }
        }
      }
    }
    returnVar.TypeCode = 3
    returnVar.IntVal = value
    return returnVar, nil
  }
  returnVar.TypeCode = 2
  returnVar.IntVal = value
  return returnVar, nil
}
