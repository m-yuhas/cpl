#if os(Linux)
  import Glibc
#else
  import Darwin
#endif

import Foundation

print("Hello")

var i = getInt()
let data = Data(bytes: i!, count: 10)
var str = String(data: data, encoding: String.Encoding.utf8)
str = str!

print("From C with love \(str)\n")
print("The length of the string is \(str!.characters.count)")

let filePath="test.txt"
var fileSize : NSNumber
do {
  let attr = try FileManager.default.attributesOfItem(atPath: filePath)
  fileSize = attr[FileAttributeKey.size] as! NSNumber
  print("File Size \(fileSize)\n")
  //var fs = [fileSize intValue]
  //print("File Size as Int \(fs)")
} catch {
  print("Error Could Not Read file")
}
