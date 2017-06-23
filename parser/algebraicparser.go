package parser

import (
  "variable"
  "strconv"
  "strings"
  //"fmt"
)

func AlgebraicParser(expression string, variableMap []map[string]variable.Variable ) variable.Variable  {
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
      part1 := AlgebraicParser(firsthalf, variableMap)
      return part1.Add(AlgebraicParser(lasthalf,variableMap))
    } else if optype == 2 {
      part1 := AlgebraicParser(firsthalf,variableMap)
      return part1.Sub(AlgebraicParser(lasthalf,variableMap))
    }
  } else if mulDivIndex != -1 {
    firsthalf := strings.TrimSpace(expression[:currIndex])
    lasthalf := strings.TrimSpace(expression[currIndex+1:])
    if optype == 3 {
      part1 := AlgebraicParser(firsthalf,variableMap)
      return part1.Mul(AlgebraicParser(lasthalf,variableMap))
    } else if optype == 4 {
      part1 := AlgebraicParser(firsthalf,variableMap)
      return part1.Div(AlgebraicParser(lasthalf,variableMap))
    } else if optype == 5 {
      part1 := AlgebraicParser(firsthalf,variableMap)
      return part1.Mod(AlgebraicParser(lasthalf,variableMap))
    }
  } else if expIndex != -1 {
    firsthalf := strings.TrimSpace(expression[:currIndex])
    lasthalf := strings.TrimSpace(expression[currIndex+1:])
    part1 := AlgebraicParser(firsthalf,variableMap)
    return part1.Exp(AlgebraicParser(lasthalf,variableMap))
  } else if expression[0] == '(' && expression[len(expression)-1] == ')' {
    return AlgebraicParser(strings.TrimSpace(expression[1:len(expression)-1]),variableMap)
  } else {
    return EvaluateAtom(expression,variableMap)
  }
  return variable.Variable{}
}

func EvaluateAtom(expression string, variableMap []map[string]variable.Variable) variable.Variable  {
  returnVar := variable.Variable{}
  value, err := strconv.ParseInt(expression,10,64)
  if err != nil {
    value, err := strconv.ParseInt(expression,64)
    if err != nil {
      if ( strings.HasPrefix(expression,"\"") || strings.HasPrefix(expression,"“") || strings.HasPrefix(expression,"”") ) && ( strings.HasSuffix(expression,"\"") || strings.HasSuffix(expression,"“") || strings.HasSuffix(expression,"”") ) {
        returnVar.TypeCode = 4
        returnVar.StringVal = expression[1:len(expression)-1]
        return returnVar
      } else {
        for _, vmap := range variableMap {
          if val, exists := vmap[expression]; exists {
            return val
          }
        }
      }
    }
    returnVar.TypeCode = 3
    returnVar.IntVal = value
    return returnVar
  }
  returnVar.TypeCode = 2
  returnVar.IntVal = value
  return returnVar
}
