// This swift file defines the Variable Class that is used for storing Variables
// (C) 2016 Michael Yuhas

enum OperationError: Error {
  case multiplyString
  case divideString
  case moduloDivideString
  case raiseStringToPower
  case stringRaisedToPower
  case factorialOfString
}

class VarObject {
  private var intVal: Int
  private var floatVal: Float
  private var stringVal: String
  private var type: UInt8

  init() {
    self.intVal = 0
    self.floatVal = 0
    self.stringVal = ""
    self.type = 0
  }

  init(initial_value : Int) {
    self.intVal = initial_value
    self.floatVal = 0
    self.stringVal = ""
    self.type = 1
  }

  init(initial_value : Float) {
    self.intVal = 0
    self.floatVal = initial_value
    self.stringVal = ""
    self.type = 2
  }

  init(initial_value : String) {
    self.intVal = 0
    self.floatVal = 0
    self.stringVal = initial_value
    self.type = 3
  }

  public func setValue(value : Int) {
    self.intVal = value
    self.type = 1
  }

  public func setValue(value : Float) {
    self.floatVal = value
    self.type = 2
  }

  public func setValue(value : String) {
    self.stringVal = value
    self.type = 3
  }

  public func getType() -> UInt8 {
    return self.type
  }

  public func getIntegerValue() -> Int {
    return self.intVal
  }

  public func getFloatValue() -> Float {
    return self.floatVal
  }

  public func getStringValue() -> String {
    return self.stringVal
  }

  public func add(addend : VarObject) -> VarObject {
    if self.type = 1 {
      if addend.getType = 1 {
        return VarObject(self.intVal + addend.getIntValue())
      } else if addend.getType = 2 {
        return VarObject(Float(self.intVal) + addend.getFloatValue())
      } else if addend.getType = 3 {
        return VarObject(String(self.intVal) + addend.getStringValue())
      }
    } else if self.type = 2 {
      if addend.getType = 1 {
        return VarObject(self.floatVal + Float(addend.getIntValue()))
      } else if addend.getType = 2 {
        return VarObject(self.floatVal + addend.getFloatValue())
      } else if addend.getType = 3 {
        return VarObject(String(self.floatVal) + addend.getStringValue())
      }
    } else if self.type = 3 {
      if addend.getType = 1 {
        return VarObject(self.stringVal + String(addend.getIntValue()))
      } else if addend.getType = 2 {
        return VarObject(self.stringVal + String(addend.getFloatValue()))
      } else if addend.getType = 3 {
        return VarObject(self.stringVal + addend.stringVal)
      }
    }
  }

  public func sub(subtrahend : VarObject) -> VarObject {
    if self.type = 1 {
      if subtrahend.getType = 1 {
        return VarObject(self.intVal - subtrahend.getIntValue())
      } else if subtrahend.getType = 2 {
        return VarObject(Float(self.intVal) - subtrahend.getFloatValue())
      } else if subtrahend.getType = 3 {
        return VarObject(String(self.intVal).replacingOccurences(of: subtrahend.getStringValue(), with: ""))
      }
    } else if self.type = 2 {
      if subtrahend.getType = 1 {
        return VarObject(self.floatVal - Float(subtrahend.getIntValue()))
      } else if subtrahend.getType = 2 {
        return VarObject(self.floatVal - subtrahend.getFloatValue())
      } else if subtrahend.getType = 3 {
        return VarObject(String(self.floatVal).replacingOccurences(of: subtrahend.getStringValue(), with: ""))
      }
    } else if self.type = 3 {
      if subtrahend.getType = 1 {
        return VarObject(self.stringVal.replacingOccurences(of: String(subtrahend.getIntValue(), with: "")))
      } else if subtrahend.getType = 2 {
        return VarObject(self.stringVal.replacingOccurences(of: String(subtrahend.getFloatValue(), with: "")))
      } else if subtrahend.getType = 3 {
        return VarObject(self.stringVal.replacingOccurences(of: subtrahend.getStringValue(), with: ""))
      }
    }
  }

  public func mul(factor : VarObject) throws -> VarObject {
    if self.type = 1 {
      if factor.getType = 1 {
        return VarObject(self.intVal * factor.getIntValue())
      } else if factor.getType = 2 {
        return VarObject(Float(self.intVal) * factor.getFloatValue())
      } else if factor.getType = 3 {
        throw OperationError.multiplyString
      }
    } else if self.type = 2 {
      if factor.getType = 1 {
        return VarObject(self.floatVal * Float(factor.getIntValue()))
      } else if factor.getType = 2 {
        return VarObject(self.floatVal * factor.getFloatValue())
      } else if factor.getType = 3 {
        throw OperationError.multiplyString
      }
    } else {
      throw OperationError.multiplyString
    }
  }

  public func div(divisor : VarObject) throws -> VarObject {
    if self.type = 1 {
      if divisor.getType = 1 {
        return VarObject(Float(self.intVal) / Float(divisor.getIntValue()))
      } else if divisor.getType = 2 {
        return VarObject(Float(self.intVal) / divisor.getFloatValue())
      } else if divisor.getType = 3 {
        throw OperationError.divideString
      }
    } else if self.type = 2 {
      if divisor.getType = 1 {
        return VarObject(self.floatVal / Float(divisor.getIntValue()))
      } else if divisor.getType = 2 {
        return VarObject(self.floatVal / divisor.getFloatValue())
      } else if divisor.getType = 3 {
        throw OperationError.divideString
      }
    } else {
      throw OperationError.divideString
    }
  }

  public func mod(divisor : VarObject) throws -> VarObject {
    if self.type = 1 {
      if divisor.getType = 1 {
        return VarObject(self.intVal % divisor.getIntValue())
      } else if divisor.getType = 2 {
        return VarObject(Float(self.intVal) % divisor.getFloatValue())
      } else if divisor.getType = 3 {
        throw OperationError.moduloDivideString
      }
    } else if self.type = 2 {
      if divisor.getType = 1 {
        return VarObject(self.floatVal % Float(divisor.getIntValue))
      } else if divisor.getType = 2 {
        return VarObject(self.floatVal % divisor.getFloatValue())
      } else if divisor.getType = 3 {
        throw OperationError.moduloDivideString
      }
    } else {
      throw OperationError.moduloDivideString
    }
  }

  public func exp(exponent : VarObject) throws -> VarObject {

  }

  public func fac() throws -> VarObject {
    
  }
}
