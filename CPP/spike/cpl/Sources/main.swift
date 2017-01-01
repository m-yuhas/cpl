#if os(Linux)
  import Glibc
#else
  import Darwin
#endif

import Foundation
import testFunct

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
var progCounter = 0
while progCounter < lineArray.count {
  print(lineArray[progCounter])
  testFunct()
  progCounter+=1
}
//for var lineNumber = 0; lineNumber < lineArray.count;
//  print("\(line)")
//}
