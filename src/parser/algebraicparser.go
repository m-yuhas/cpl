package parser

import (
  "fmt"
)

func AlgebraicParser(expression string) {
  parenthCount := 0
  addSubIndex := -1
  mulDivIndex := -1
  expIndex := -1
  opType := -1
  currIndex := -1
  currRune rune
  for pos, char := range expression {
    currRune = []rune(expression)[pos]
    if currRune == "(" {
      parenthCount += 1
      continue
    } else if currRune == ")" {
      parenthCount -= 1
      continue
    } else if parenthCount == 0 {
      switch currRune {
      case "+":
        addSubIndex = pos
        optype = 1
        currIndex = pos
      case "-":
         addSubIndex = pos
         optype = 2
         currIndex = pos
      case "*":
        mulDivIndex = pos
        optype = 3
        currIndex = pos
      case "/":
        mulDivIndex = pos
        optype = 4
        currIndex = pos
      case "%":
        mulDivIndex = pos
        optype = 5
        currIndex = pos
      case "^":
        expIndex = pos
        optype = 6
        currIndex = pos
      }
    }
  }
  if addSubIndex != -1 {
    firsthalf = 
  }
  fmt.Println("HERE")
}

if addSubIndex != -1 {
  let firsthalf = expression.substring(to: expression.index(expression.startIndex, offsetBy: currIndex))
  let lasthalf = expression.substring(from: expression.index(expression.startIndex, offsetBy: currIndex+1))
  if optype == 1 {
    do {
      return try parseExpression( expression:firsthalf ).add( addend:parseExpression( expression:lasthalf ) )
    } catch {
      throw ExpressionError.invalidSyntax
    }
  }
  if optype == 2 {
    do {
      return try parseExpression( expression:firsthalf ).sub( subtrahend:parseExpression( expression:lasthalf ) )
    } catch {
      throw ExpressionError.invalidSyntax
    }
  }
} else if mulDivIndex != -1 {
  let firsthalf = expression.substring(to: expression.index(expression.startIndex, offsetBy: currIndex))
  let lasthalf = expression.substring(from: expression.index(expression.startIndex, offsetBy: currIndex+1))  //TODO check for last character in array
  if optype == 3 {
    do {
      return try parseExpression( expression:firsthalf ).mul( factor:parseExpression( expression:lasthalf ) )
    } catch {
      throw ExpressionError.invalidSyntax
    }
  }
  if optype == 4 {
    do {
      return try parseExpression( expression:firsthalf ).div( divisor:parseExpression( expression:lasthalf ) )
    } catch {
      throw ExpressionError.invalidSyntax
    }
  }
  if optype == 5 {
    do {
      return try parseExpression( expression:firsthalf ).mod( divisor:parseExpression( expression:lasthalf ) )
    } catch {
      throw ExpressionError.invalidSyntax
    }
  }
} else if expIndex != -1 {
  let firsthalf = expression.substring(to: expression.index(expression.startIndex, offsetBy: currIndex))
  let lasthalf = expression.substring(from: expression.index(expression.startIndex, offsetBy: currIndex+1))
  do {
    return try parseExpression( expression:firsthalf ).exp( exponent:parseExpression( expression:lasthalf ) )
  } catch {
    throw ExpressionError.invalidSyntax
  }
} else if expression[expression.startIndex] == "(" && expression[expression.endIndex] == ")" {
  let range = expression.index(expression.startIndex, offsetBy: 1)..<expression.index(expression.endIndex, offsetBy: -1)
  let removedParenths = expression.substring(with: range)
  do {
    return try parseExpression( expression:removedParenths )
  } catch {
    throw ExpressionError.invalidSyntax
  }
} else {
  do {
    return try evaluateAtom( expression:expression )
  } catch {
    throw ExpressionError.invalidSyntax
  }
}
throw ExpressionError.invalidSyntax
}
