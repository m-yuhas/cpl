// State Machine for Interpreting a Program
// (C) 2016 Michael Yuhas

#if os(Linux)
  import Glibc
#else
  import Darwin
#endif

import Foundation

func stateMachine( lineArray: [String] ) -> Int {
  varList.append([String: VarObject]())
  var progCounter = 0
  while progCounter < lineArray.count {
    let thisLine = lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces)
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
      do {
        if try evaluateBoolean( input_string: tempString.trimmingCharacters(in: CharacterSet.whitespaces) ) {
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
          if stateMachine( lineArray:subRoutineArray ) == 1 {
            return 1
          }  //TODO Have a way to store variables
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
              if stateMachine( lineArray:subRoutineArray ) == 1 {
                return 1
              }
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
        break
      }
    } else if thisLine.hasPrefix("从") {
      let tempString = thisLine.substring(from: thisLine.index(thisLine.startIndex, offsetBy:1)).trimmingCharacters(in: CharacterSet.whitespaces)
      let initCondEndIndex = tempString.range(of: " 直到")
      if initCondEndIndex == nil {
        print("错误：句法不对 （第\(progCounter)句）")
        print(lineArray[progCounter])
        break
      }
      //print(tempString)
      //let initConditionString = tempString.substring(to: initCondEndIndex!.lowerBound)
      let endConditionString = tempString.substring(from: tempString.index(initCondEndIndex!.lowerBound, offsetBy: 3))
      progCounter += 1
      var loopLevel = 0
      var subRoutineArray = [String]()
      while true {
        if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("当") || lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("从") {
          loopLevel += 1
        } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束圈" && loopLevel == 0 {
          progCounter += 1
          break
        } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束圈" {
          loopLevel -= 1
        }
        subRoutineArray.append(lineArray[progCounter])
        progCounter += 1
      }
      //print( endConditionString )
      do {
        let CounterVarString = try storeVar( expression: tempString.substring(to: initCondEndIndex!.lowerBound) ).trimmingCharacters(in: CharacterSet.whitespaces)
        //print(CounterVarString)
        let FinalCondition = try parseExpression( expression: endConditionString.trimmingCharacters(in: CharacterSet.whitespaces) )
        outerLoop: while true {
          for i in 0..<varList.count {
            if varList[i][CounterVarString] != nil {
              if try compare( val1: varList[i][CounterVarString]!, val2: FinalCondition, optype: 3 ) {
                break outerLoop
              }
            }
          }
          if stateMachine( lineArray: subRoutineArray ) == 1 {
            break
          }
        }
        continue
      } catch {
        print("错误：命令不清楚（第\(progCounter)句）")
        print(lineArray[progCounter])
        break
      }
    } else if thisLine.hasPrefix("当") {
      let tempString = thisLine.substring(from: thisLine.index(thisLine.startIndex, offsetBy:1)).trimmingCharacters(in: CharacterSet.whitespaces)
      progCounter += 1
      var loopLevel = 0
      var subRoutineArray = [String]()
      while true {
        if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("当") || lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces).hasPrefix("从") {
          loopLevel += 1
        } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束圈" && loopLevel == 0 {
          progCounter += 1
          break
        } else if lineArray[progCounter].trimmingCharacters(in: CharacterSet.whitespaces) == "结束圈" {
          loopLevel -= 1
        }
        subRoutineArray.append(lineArray[progCounter])
        progCounter += 1
      }
      do {
        while try evaluateBoolean( input_string: tempString ) {
          if stateMachine( lineArray: subRoutineArray ) == 1 {
            break
          }
        }
        continue
      } catch {
        print("错误：命令不清楚（第\(progCounter)句）")
        print(lineArray[progCounter])
        break
      }
    } else if thisLine == "否则" || thisLine == "结束支" || thisLine == "结束圈" {
      print("错误：命令不清楚（第\(progCounter)句）")
      print(lineArray[progCounter])
      break
    } else if thisLine == "跳出" {
      return 1
    } else if progCounter == 0 && thisLine.hasPrefix("#") {
      progCounter+=1
      continue
    } else {
      let charset = CharacterSet(charactersIn: "=")
      if thisLine.rangeOfCharacter(from: charset) != nil {
        do {
          _ = try storeVar( expression: thisLine )
          progCounter+=1
          continue
        } catch {
          print("错误：命令不清楚（第\(progCounter)句）")
          print(lineArray[progCounter])
          break
        }
      }
      print("错误：命令不清楚（第\(progCounter)句）")
      print(lineArray[progCounter])
      break
    }
  }
  varList.removeLast()
  return 0
}
