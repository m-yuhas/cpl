// This function evaluates a Boolean Expression
// Definition of Variable object available in varObject.swift
// (C) 2016 Michael Yuhas

enum BooleanExpressionError: Error {
  case invalidSyntax
}

func evaluateBoolean( input_string: String ) throws -> Bool {
  var expression = input_string
  if expression[expression.startIndex] == "(" && expression[expression.index(expression.endIndex, offsetBy:-1)] == ")" {
    expression.remove(at: expression.startIndex)
    expression.remove(at: expression.index(expression.endIndex, offsetBy:-1))
  }
  var firsthalf = ""
  var lasthalf = ""
  var optype = -1
  var charcount = -1


  for char in expression.characters {
    if char == "=" {
      optype = 1
      break
    } else if char == "<" {
      optype = 2
      break
    } else if char == ">" {
      optype = 3
      break
    }
    firsthalf.append(char)
    charcount+=1
  }
  if optype == -1 {
    do {
      let temp = try parseExpression( expression: firsthalf )
      if temp.getType() == 1 {
        if temp.getIntegerValue() != 0 {
          return true
        } else {
          return false
        }
      } else if temp.getType() == 2 {
        if temp.getFloatValue() != 0 {
          return true
        } else {
          return false
        }
      } else if temp.getType() == 3 {
        return true
      } else {
        throw BooleanExpressionError.invalidSyntax
      }
    } catch {
      return false
    }
  }
  for char in expression.substring(to: expression.index(expression.startIndex, offsetBy: charcount+1)).characters {
    if char == "=" || char == "<" || char == ">" {
      throw BooleanExpressionError.invalidSyntax
    }
    lasthalf.append(char)
  }
  do {
    let evalfirst = try parseExpression( expression: firsthalf )
    let evallast = try parseExpression( expression: lasthalf )
    return try compare( val1: evalfirst, val2: evallast, optype: optype )
  } catch {
    throw BooleanExpressionError.invalidSyntax
  }
}

func compare( val1: VarObject, val2: VarObject, optype: Int ) throws -> Bool {
  if optype == 1 {
    if val1.getType() == 1 {
      if val2.getType() == 1 {
        return val1.getIntegerValue() == val2.getIntegerValue()
      } else if val2.getType() == 2 {
        return Float(val1.getIntegerValue()) == val2.getFloatValue()
      } else if val2.getType() == 3 {
        return String(val1.getIntegerValue()) == val2.getStringValue()
      }
    } else if val1.getType() == 2 {
      if val2.getType() == 1 {
        return val1.getFloatValue() == Float(val2.getIntegerValue())
      } else if val2.getType() == 2 {
        return val1.getFloatValue() == val2.getFloatValue()
      } else if val2.getType() == 3 {
        return String(val1.getFloatValue()) == val2.getStringValue()
      }
    } else if val1.getType() == 3 {
      if val2.getType() == 1 {
        return val1.getStringValue() == String(val2.getIntegerValue())
      } else if val2.getType() == 2 {
        return val1.getStringValue() == String(val2.getFloatValue())
      } else if val2.getType() == 3 {
        return val1.getStringValue() == val2.getStringValue()
      }
    }
  } else if optype == 2 {
    if val1.getType() == 1 {
      if val2.getType() == 1 {
        return val1.getIntegerValue() < val2.getIntegerValue()
      } else if val2.getType() == 2 {
        return Float(val1.getIntegerValue()) < val2.getFloatValue()
      } else if val2.getType() == 3 {
        return String(val1.getIntegerValue()).characters.count < val2.getStringValue().characters.count
      }
    } else if val1.getType() == 2 {
      if val2.getType() == 1 {
        return val1.getFloatValue() < Float(val2.getIntegerValue())
      } else if val2.getType() == 2 {
        return val1.getFloatValue() < val2.getFloatValue()
      } else if val2.getType() == 3 {
        return String(val1.getFloatValue()).characters.count < val2.getStringValue().characters.count
      }
    } else if val1.getType() == 3 {
      if val2.getType() == 1 {
        return val1.getStringValue().characters.count < String(val2.getIntegerValue()).characters.count
      } else if val2.getType() == 2 {
        return val1.getStringValue().characters.count < String(val2.getFloatValue()).characters.count
      } else if val2.getType() == 3 {
        return val1.getStringValue().characters.count < val2.getStringValue().characters.count
      }
    }
  } else if optype == 3 {
    if val1.getType() == 1 {
      if val2.getType() == 1 {
        return val1.getIntegerValue() > val2.getIntegerValue()
      } else if val2.getType() == 2 {
        return Float(val1.getIntegerValue()) > val2.getFloatValue()
      } else if val2.getType() == 3 {
        return String(val1.getIntegerValue()).characters.count > val2.getStringValue().characters.count
      }
    } else if val1.getType() == 2 {
      if val2.getType() == 1 {
        return val1.getFloatValue() > Float(val2.getIntegerValue())
      } else if val2.getType() == 2 {
        return val1.getFloatValue() > val2.getFloatValue()
      } else if val2.getType() == 3 {
        return String(val1.getFloatValue()).characters.count > val2.getStringValue().characters.count
      }
    } else if val1.getType() == 3 {
      if val2.getType() == 1 {
        return val1.getStringValue().characters.count > String(val2.getIntegerValue()).characters.count
      } else if val2.getType() == 2 {
        return val1.getStringValue().characters.count > String(val2.getFloatValue()).characters.count
      } else if val2.getType() == 3 {
        return val1.getStringValue().characters.count > val2.getStringValue().characters.count
      }
    }
  }
  throw BooleanExpressionError.invalidSyntax
}
