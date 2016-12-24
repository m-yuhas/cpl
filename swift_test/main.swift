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
let data = Data(bytes: i!, count: 100)
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
var progCounter = 0
while progCounter < lineArray.count {
  var thisLine = lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces)
  if thisLine.isEmpty {
    progCounter+=1
    continue
  } else if thisLine.hasPrefix("注解：") {
    progCounter+=1
    continue
  } else if thisLine.hasPrefix("宣") {
    thisLine.remove(at: thisLine.startIndex)
    do {
      try basicOutput( output_text: thisLine.trimmingCharacters(in: CharacterSet.whitespaces) )
    } catch {
      print("错误：句法不对 （第\(progCounter)句）")
      print(lineArray[progCounter])
      break
    }
    progCounter+=1
  } else {
    print("错误：命令不清楚（第\(progCounter)句）")
    print(lineArray[progCounter])
    break
  }
}
