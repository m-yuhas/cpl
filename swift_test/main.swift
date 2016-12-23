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

let s = "test.txt"
let cs = "test.txt"
var filePath = "test.txt"
var i = readFile( filePath )
let data = Data(bytes: i!, count: 100)
var str = String(data: data, encoding: String.Encoding.utf8)!
var lineArray = str.components(separatedBy : "\n")

print("2+2")
do {
  try print(parseExpression(expression:"2+2").getIntegerValue())
} catch ExpressionError.unknownVarName {
  print("Error Unknown var name")
}


var varList = [String : VarObject]()
var progCounter = 0
while progCounter < lineArray.count {
  lineArray[progCounter] = lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces)
  if lineArray[progCounter].isEmpty {
    progCounter+=1
    continue
  } else if lineArray[progCounter].hasPrefix("注解：") {
    progCounter+=1
    continue
  } else {
    print("错误：命令不清楚（第\(progCounter)句）")
    print(lineArray[progCounter])
    break
  }
}
