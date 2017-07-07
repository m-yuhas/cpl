package parser

import (
  //"fmt"
  "strings"
  "cpl/variable"
  //"os"
)

func BooleanParser(expression string, variableMap []map[string]variable.Variable ) (variable.Variable, error) {
  parenthCount := 0
  opIndex := -1
  notIndex := -1
  eqIndex := -1
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
      case '与':
        opIndex = pos
        optype = 1
        currIndex = pos
        break
      case '或':
        opIndex = pos
        optype = 2
        currIndex = pos
        break
      case '非':
        notIndex = pos
        optype = 3
        currIndex = pos
        break
      case '=':
        eqIndex = pos
        optype = 4
        currIndex = pos
        break
      case '<':
        eqIndex = pos
        optype = 5
        currIndex = pos
        break
      case '>':
        eqIndex = pos
        optype = 6
        currIndex = pos
        break
      }
    }
  }
  if opIndex != -1 {
    firsthalf := strings.TrimSpace(expression[:currIndex])
    lasthalf := strings.TrimSpace(expression[currIndex+1:])
    part1, err := BooleanParser(firsthalf,variableMap)
    if err != nil {
      return part1, err
    }
    part2, err := BooleanParser(lasthalf,variableMap)
    if err != nil {
      return part2, err
    }
    if optype == 1 {
      return part1.And(part2)
    } else if optype == 2 {
      return part1.Or(part2)
    }
  } else if notIndex != -1 {
    part, err := BooleanParser(strings.TrimSpace(expression[currIndex+1:]),variableMap)
    if err != nil {
      return part, err
    }
    return part.Not()
  } else if eqIndex != -1 {
    firsthalf := strings.TrimSpace(expression[:currIndex])
    lasthalf := strings.TrimSpace(expression[currIndex+1:])
    part1, err := BooleanParser(firsthalf,variableMap)
    if err != nil {
      return part1, err
    }
    part2, err := BooleanParser(lasthalf,variableMap)
    if err != nil {
      return part2, err
    }
    if optype == 4 {
      return part1.Eq(part2)
    } else if optype == 5 {
      return part1.Lt(part2)
    } else if optype == 6 {
      return part1.Gt(part2)
    }
  } else {
    return AlgebraicParser(expression,variableMap)
  }
  return variable.Variable{}, nil
}
