// State Machine for Interpreting a Program
// (C) 2016 Michael Yuhas

#if os(Linux)
  import Glibc
#else
  import Darwin
#endif

import Foundation

func stateMachine( lineArray: [String], CurrVarList: Dictionary<String,VarObject> ) {
  var progCounter = 0
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
        try basicOutput( output_text: tempString.trimmingCharacters(in: CharacterSet.whitespaces), varList:CurrVarList )
      } catch {
        print("错误：句法不对 （第\(progCounter)句）")
        print(lineArray[progCounter])
        break
      }
      progCounter+=1
    } else if thisLine.hasPrefix("如果") {
      let tempString = thisLine.substring(from: thisLine.index(thisLine.startIndex, offsetBy:2))
      do {
        if try evaluateBoolean( input_string: tempString.trimmingCharacters(in: CharacterSet.whitespaces), varList:CurrVarList ) {
          progCounter+=1
          var ifLevel = 0
          var subRoutineArray = [String]()
          while true {
            if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("如果") {
              ifLevel += 1
            } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "否则" && ifLevel == 0 {
              progCounter += 1
              while true {
                if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("如果") {
                  ifLevel += 1
                } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" && ifLevel == 0 {
                  progCounter += 1
                  break
                } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" {
                  ifLevel -= 1
                }
                progCounter += 1
              }
              break
            } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" && ifLevel == 0 {
              progCounter += 1
              break
            } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" {
              ifLevel -= 1
            }
            subRoutineArray.append(lineArray[progCounter])
            progCounter += 1
          }
          stateMachine( lineArray:subRoutineArray, CurrVarList:CurrVarList)  //TODO Have a way to store variables
          continue
        } else {
          progCounter += 1
          var ifLevel = 0
          while true {
            if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("如果") {
              ifLevel += 1
            } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "否则" && ifLevel == 0 {
              progCounter += 1
              var subRoutineArray = [String]()
              while true {
                if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("如果") {
                  ifLevel += 1
                } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" && ifLevel == 0 {
                  progCounter += 1
                  break
                } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" {
                  ifLevel -= 1
                }
                subRoutineArray.append(lineArray[progCounter])
                progCounter += 1
              }
              stateMachine( lineArray:subRoutineArray, CurrVarList:CurrVarList)
              break
            } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束支" && ifLevel == 0 {
              progCounter += 1
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
    } else if thisLine == "否则" || thisLine == "结束支" || thisLine == "结束圈" {
      print("错误：命令不清楚（第\(progCounter)句）")
      print(lineArray[progCounter])
      break
    } else if progCounter == 0 && thisLine.hasPrefix("#") {
      progCounter+=1
      continue
    } else {
      print("错误：命令不清楚（第\(progCounter)句）")
      print(lineArray[progCounter])
      break
    }
  }

}
