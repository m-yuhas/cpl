#if os(Linux)
  import Glibc
#else
  import Darwin
#endif

import Foundation

print("Hello")

let s = "test.txt"
let cs = "test.txt"
//var filePath=UnsafeMutablePointer<Int8>(cs)
var filePath = "test.txt"
var i = readFile( filePath )
let data = Data(bytes: i!, count: 100)
var str = String(data: data, encoding: String.Encoding.utf8)!
//str = str!

//print("From C with love \(str)\n")
//print("The length of the string is \(str.characters.count)")

//var result: [String] = []
//enumerateLines { line, _ in result.append(line) }
var lineArray = str.components(separatedBy : "\n")
//for line in lineArray {
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
//for var lineNumber = 0; lineNumber < lineArray.count;
//  print("\(line)")
//}
