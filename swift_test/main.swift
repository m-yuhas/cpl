#if os(Linux)
  import Glibc
#else
  import Darwin
#endif

import Foundation

// TODO: Add Error Checking For Open FIle Errors
// TODO: Add C Function to get File length
// TODO: Finish Math parser
// TODO: 0.0 Release: Print, Input, Store, Loop, Branch

/*
if C_ARGC <= 1 {
  print("用法：cpl 《程序地址》")
  exit(0)
} else if C_ARG == 2 {
  filePath = C_ARGV[1]
} else {
  print("还不支持那么多args")
  exit(0)
}*/


if CommandLine.arguments.count <= 1 {
  print("用法：cpl 《程序地址》")
  exit(0)
} else if CommandLine.arguments.count > 2 {
  print("还不支持那么多args")
  exit(0)
}

var filePath = CommandLine.arguments[1]
//var filePath = "test.txt"
var i = readFile( filePath )
let data = Data(bytes: i!, count: 1000)
var str = String(data: data, encoding: String.Encoding.utf8)!
if str == "ERR001" || str == "ERR002" {
  print("错误：不能开文件")
  exit(0)
}
var lineArray = str.components(separatedBy : "\n")

//print("2*5")
//do {
//  try print(parseExpression(expression:"2-5").getIntegerValue())
//} catch ExpressionError.unknownVarName {
//  print("Error Unknown var name")
//}

var varList = [String : VarObject]()
stateMachine( lineArray:lineArray, CurrVarList:varList)

/*var progCounter = 0
var ifLevel = 0
var loopLevel = 0
while progCounter < lineArray.count {
  var thisLine = lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces)
  if thisLine.isEmpty {
    progCounter+=1
    continue
  } else if thisLine.hasPrefix("注解：") || thisLine.hasPrefix("注解:") {
    progCounter+=1
    continue
  } else if thisLine.hasPrefix("输出:") || thisLine.hasPrefix("输出：") {
    let tempString = thisLine.substring(from: thisLine.index(thisLine.startIndex, offsetBy:3))
    do {
      try basicOutput( output_text: tempString.trimmingCharacters(in: CharacterSet.whitespaces) )
    } catch {
      print("错误：句法不对 （第\(progCounter)句）")
      print(lineArray[progCounter])
      break
    }
    progCounter+=1
  } else if thisLine.hasPrefix("如果") {
    let tempString = thisLine.substring(from: thisLine.index(thisLine.startIndex, offsetBy:2))
    ifLevel += 1
    do {
      if try evaluateBoolean( input_string: tempString.trimmingCharacters(in: CharacterSet.whitespaces) ) {
        progCounter+=1
        continue
      } else {
        progCounter += 1
        var startingIfLevel = ifLevel
        while true {
          if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("如果") {
            ifLevel += 1
          } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "否则" && startingIfLevel == ifLevel {
            break
          } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" && startingIfLevel == ifLevel {
            break
          } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" {
            ifLevel -= 1
          }
          progCounter += 1
        }
        continue
      }
    } catch {
      print("错误：句法不对 （第\(progCounter)句）")
      print(lineArray[progCounter])
    }
    continue
  } else if thisLine.hasPrefix("从") {
    continue
  } else if thisLine.hasPrefix("当") {
    continue
  } else if thisLine == "否则" {
    if ifLevel < 1 {
      print("错误：突然“否则”（第\(progCounter)句）")
      print(lineArray[progCounter])
    }
    progCounter += 1
    continue
  } else if thisLine == "结束支" {
    ifLevel -= 1
    if ifLevel < 0 {
      print("错误：突然“结束支”（第\(progCounter)句）")
      print(lineArray[progCounter])
    }
    progCounter += 1
    continue
  } else if thisLine == "结束圈" {
  } else if false {
    continue
  } else if progCounter == 0 && thisLine.hasPrefix("#") {
    progCounter+=1
    continue
  } else {
    print("错误：命令不清楚（第\(progCounter)句）")
    print(lineArray[progCounter])
    break
  }
}*/
