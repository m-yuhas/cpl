// This function stores a Variable
// (C) 2016 Michael Yuhas

#if os(Linux)
  import Glibc
#else
  import Darwin
#endif

import Foundation

enum VariableStorageError: Error {
  case invalidSyntax
}

func storeVar( expression : String ) throws {
  var firsthalf = ""
  var lasthalf = ""
  var charcount = 0
  for char in expression.characters {
    if char == "=" {
      charcount += 1
      break
    }
    firsthalf.append(char)
    charcount += 1
  }
  for char in expression.substring(from: expression.index(expression.startIndex, offsetBy: charcount)).characters {
    if char == "=" {
      throw VariableStorageError.invalidSyntax
    }
    lasthalf.append(char)
  }
  do {
    let evalLast = try parseExpression( expression: lasthalf )
    let charset = CharacterSet(charactersIn: "+=-/*^!#\"():")
    if firsthalf.rangeOfCharacter(from: charset) != nil {
      throw VariableStorageError.invalidSyntax
    }
    for i in 0..<varList.count {
      if varList[i][firsthalf] != nil {
        varList[i][firsthalf] = evalLast
        return
      }
    }
    varList[varList.count - 1][firsthalf] = evalLast
    return
  } catch {
    throw VariableStorageError.invalidSyntax
  }
}
