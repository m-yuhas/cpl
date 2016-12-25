// this implements the Input/Output Command Text Display
// Definition of Variable object available in varObject.swift
// (C) 2016 Michael Yuhas

enum OutputError: Error {
  case invalidSyntax
}

func basicOutput( output_text: String, varList:Dictionary<String,VarObject> ) throws {
  var expression = output_text
  if expression[expression.startIndex] != "\"" || expression[expression.index(expression.endIndex, offsetBy:-1)] != "\"" {
    throw OutputError.invalidSyntax
  }
  expression.remove(at: expression.startIndex)
  expression.remove(at: expression.index(expression.endIndex, offsetBy:-1))
  var i=0
  while i < expression.characters.count {
    if expression[expression.index(expression.startIndex, offsetBy:i)] == "#" {
      i+=1
      if expression[expression.index(expression.startIndex, offsetBy:i)] == " " {
        print("# ", terminator:"")
        i+=1
        continue
      } else if expression[expression.index(expression.startIndex, offsetBy:i)] == "#" {
        print("#", terminator: "")
        i+=1
        continue
      } else {
        var exprName = ""
        if expression[expression.index(expression.startIndex, offsetBy:i)] == "(" {
          i+=1
          var parenthCount = 0
          while i < expression.characters.count {
            if expression[expression.index(expression.startIndex, offsetBy:i)] == "(" {
              exprName.append(expression[expression.index(expression.startIndex, offsetBy:i)])
              parenthCount+=1
            } else if expression[expression.index(expression.startIndex, offsetBy:i)] == ")" && parenthCount == 0 {
              i+=1
              break
            } else if expression[expression.index(expression.startIndex, offsetBy:i)] == ")" {
              exprName.append(expression[expression.index(expression.startIndex, offsetBy:i)])
              parenthCount-=1
            } else {
              exprName.append(expression[expression.index(expression.startIndex, offsetBy:i)])
            }
            i+=1
          }
        } else {
          while i < expression.characters.count {
            if expression[expression.index(expression.startIndex, offsetBy:i)] != " " {
              exprName.append(expression[expression.index(expression.startIndex, offsetBy:i)])
              i+=1
            } else {
              break
            }
          }
        }
        do {
          let varToPrint = try parseExpression(expression:exprName, varList:varList)
          if varToPrint.getType() == 1 {
            print(varToPrint.getIntegerValue(), terminator: "")
          } else if varToPrint.getType() == 2 {
            print(varToPrint.getFloatValue(), terminator: "")
          } else if varToPrint.getType() == 3 {
            print(varToPrint.getStringValue(), terminator: "")
          }
        } catch {
          throw OutputError.invalidSyntax
        }
      }
    } else {
      print(expression[expression.index(expression.startIndex, offsetBy:i)], terminator:"")
      i+=1
    }
  }
  print("", terminator: "\n")
}
