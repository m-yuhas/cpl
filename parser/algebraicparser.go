package parser

import (
  "cpl/variable"
  "strconv"
  "strings"
  //"fmt"
)

func AlgebraicParser(expression string, variableMap []map[string]variable.Variable ) (variable.Variable, error)  {
  parenthCount := 0
  addSubIndex := -1
  mulDivIndex := -1
  expIndex := -1
  optype := -1
  currIndex := -1
  var currRune rune
  for pos, _ := range expression {
    currRune = []rune(expression)[pos]
    if currRune == rune('(') {
      parenthCount += 1
      continue
    } else if currRune == rune(')') {
      parenthCount -= 1
      continue
    } else if parenthCount == 0 {
      switch currRune {
      case '+':
        addSubIndex = pos
        optype = 1
        currIndex = pos
        break
      case '-':
         addSubIndex = pos
         optype = 2
         currIndex = pos
         break
      case '*':
        mulDivIndex = pos
        optype = 3
        currIndex = pos
        break
      case '/':
        mulDivIndex = pos
        optype = 4
        currIndex = pos
        break
      case '%':
        mulDivIndex = pos
        optype = 5
        currIndex = pos
        break
      case '^':
        expIndex = pos
        optype = 6
        currIndex = pos
        break
      }
    }
  }
  //fmt.Println(expression)
  if addSubIndex != -1 {
    firsthalf := strings.TrimSpace(expression[:currIndex])
    lasthalf := strings.TrimSpace(expression[currIndex+1:])
    if optype == 1 {
      part1, err := AlgebraicParser(firsthalf, variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(lasthalf, variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Add(part2)
    } else if optype == 2 {
      part1, err := AlgebraicParser(firsthalf,variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(lasthalf,variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Sub(part2)
    }
  } else if mulDivIndex != -1 {
    firsthalf := strings.TrimSpace(expression[:currIndex])
    lasthalf := strings.TrimSpace(expression[currIndex+1:])
    if optype == 3 {
      part1, err := AlgebraicParser(firsthalf,variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(lasthalf,variableMap)
      if err != nil {
        return part2, err
      }
      return part1.Mul(part2)
    } else if optype == 4 {
      part1, err := AlgebraicParser(firsthalf,variableMap)
      if err != nil {
        return part1, err
      }
      part2, err := AlgebraicParser(lasthalf,variableMap)
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
