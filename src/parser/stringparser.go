package parser

import (
  "variable"
  "strconv"
  //"strings"
  //"fmt"
)

func StringParser(expression string, variableMap []map[string]variable.Variable ) string {
  //var startQuoteType rune
  inQuotes := false
  var outputSlice string
  for i := 0; i < len(expression)-1; i++ {
    if expression[i] == '"' {
      //startQuoteType = '"'
      inQuotes = !inQuotes
      continue
    }
    if inQuotes && expression[i] == '#' {
      sliceToParse := ""
      i++
      for i < len(expression)-1 {
        if expression[i] == ' ' {
          i++
          break
        }
        if expression[i] == ')' {
          //sliceToParse = append(sliceToParse,string(expression[i]))
          sliceToParse = sliceToParse + string(expression[i])
          i++
          break
        }
        //sliceToParse = append(sliceToParse,string(expression[i]))
        sliceToParse = sliceToParse + string(expression[i])
        i++
      }
      //outputSlice = append(outputSlice,string(AlgebraicParser(string(sliceToParse),variableMap).IntVal))
      //fmt.Println(string(sliceToParse))
      //outputSlice = outputSlice + AlgebraicParser(string(sliceToParse),variableMap).IntVal
      outputSlice = string(strconv.AppendInt([]byte(outputSlice),AlgebraicParser(string(sliceToParse),variableMap).IntVal,10))
      continue
    }
    //outputSlice = append(outputSlice,string(expression[i]))
    outputSlice = outputSlice + string(expression[i])
  }
  return string(outputSlice)
}
