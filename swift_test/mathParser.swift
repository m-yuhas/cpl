// An algebraic parser in swift, relies on Access to a dictionary of Variables
// Definition of Variable object available in varObject.swift
// (C) 2016 Michael Yuhas

enum ExpressionError: Error {
  case invalidParenthesisUse
  case invalidSyntax
  case unknownVarName
}

func parseExpression(expression : String) throws -> VarObject {
  var parenthCount = 0
  var addSubIndex = -1
  var mulDivIndex = -1
  var expIndex = -1
  //var facIndex = -1
  var optype = -1
  var currIndex = -1
  for i in 0..<expression.characters.count {
    let index = expression.index(expression.startIndex, offsetBy: i)
    if expression[index] == "(" {
      parenthCount += 1
      continue
    } else if expression[index] == ")" {
      parenthCount -= 1
      continue
    } else if parenthCount == 0 {
      if expression[index] == "+" {
        addSubIndex = i
        optype = 1
        currIndex = i
        break
      } else if expression[index] == "-" {
        addSubIndex = i
        optype = 2
        currIndex = i
        break
      } else if expression[index] == "*" {
        mulDivIndex = i
        optype = 3
        currIndex = i
        break
      } else if expression[index] == "/" {
        mulDivIndex = i
        optype = 4
        currIndex = i
        break
      } else if expression[index] == "%" {
        mulDivIndex = i
        optype = 5
        currIndex = i
        break
      } else if expression[index] == "^" {
        expIndex = i
        optype = 6
        currIndex = i
        break
      } else if expression[index] == "!" {
        //facIndex = i
        optype = 7
        currIndex = i
        break
      }
    }
  }
  //TODO Make this more efficient and only use optype
  //TODO Implement boolean operators
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

func evaluateAtom(expression : String) throws -> VarObject {
  if Int(expression) != nil {
    return VarObject(initial_value: Int(expression)! )
  } else if Float(expression) != nil {
    return VarObject(initial_value: Float(expression)! )
  } else if varList[expression] != nil {
    return varList[expression]!
  } else {
    throw ExpressionError.unknownVarName
  }
}
