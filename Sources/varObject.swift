// This swift file defines the Variable Class that is used for storing Variables
// (C) 2016 Michael Yuhas

enum OperationError: Error {
  case multiplyString
  case divideString
  case moduloDivideString
  case raiseStringToPower
  case stringRaisedToPower
  case factorialOfString
  case invalidState
}

class VarObject {
  private var intVal: Int
  private var floatVal: Float
  private var stringVal: String
  private var type: Int

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

  public func getType() -> Int {
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

  public func add(addend : VarObject) throws -> VarObject {
    if self.type == 1 {
      if addend.getType() == 1 {
        return VarObject(initial_value: self.intVal + addend.getIntegerValue())
      } else if addend.getType() == 2 {
        return VarObject(initial_value: Float(self.intVal) + addend.getFloatValue())
      } else if addend.getType() == 3 {
        return VarObject(initial_value: String(self.intVal) + addend.getStringValue())
      }
    } else if self.type == 2 {
      if addend.getType() == 1 {
        return VarObject(initial_value: self.floatVal + Float(addend.getIntegerValue()))
      } else if addend.getType() == 2 {
        return VarObject(initial_value: self.floatVal + addend.getFloatValue())
      } else if addend.getType() == 3 {
        return VarObject(initial_value: String(self.floatVal) + addend.getStringValue())
      }
    } else if self.type == 3 {
      if addend.getType() == 1 {
        return VarObject(initial_value: self.stringVal + String(addend.getIntegerValue()))
      } else if addend.getType() == 2 {
        return VarObject(initial_value: self.stringVal + String(addend.getFloatValue()))
      } else if addend.getType() == 3 {
        return VarObject(initial_value: self.stringVal + addend.stringVal)
      }
    }
    throw OperationError.invalidState
  }

  public func sub(subtrahend : VarObject) throws -> VarObject {
    if self.type == 1 {
      if subtrahend.getType() == 1 {
        return VarObject(initial_value: self.intVal - subtrahend.getIntegerValue())
      } else if subtrahend.getType() == 2 {
        return VarObject(initial_value: Float(self.intVal) - subtrahend.getFloatValue())
      } else if subtrahend.getType() == 3 {
        return VarObject(initial_value: String(self.intVal).replacingOccurrences(of: subtrahend.getStringValue(), with: ""))
      }
    } else if self.type == 2 {
      if subtrahend.getType() == 1 {
        return VarObject(initial_value: self.floatVal - Float(subtrahend.getIntegerValue()))
      } else if subtrahend.getType() == 2 {
        return VarObject(initial_value: self.floatVal - subtrahend.getFloatValue())
      } else if subtrahend.getType() == 3 {
        return VarObject(initial_value: String(self.floatVal).replacingOccurrences(of: subtrahend.getStringValue(), with: ""))
      }
    } else if self.type == 3 {
      if subtrahend.getType() == 1 {
        return VarObject(initial_value: self.stringVal.replacingOccurrences(of: String(subtrahend.getIntegerValue()), with: ""))
      } else if subtrahend.getType() == 2 {
        return VarObject(initial_value: self.stringVal.replacingOccurrences(of: String(subtrahend.getFloatValue()), with: ""))
      } else if subtrahend.getType() == 3 {
        return VarObject(initial_value: self.stringVal.replacingOccurrences(of: subtrahend.getStringValue(), with: ""))
      }
    }
    throw OperationError.invalidState
  }

  public func mul(factor : VarObject) throws -> VarObject {
    if self.type == 1 {
      if factor.getType() == 1 {
        return VarObject(initial_value: self.intVal * factor.getIntegerValue())
      } else if factor.getType() == 2 {
        return VarObject(initial_value: Float(self.intVal) * factor.getFloatValue())
      } else if factor.getType() == 3 {
        throw OperationError.multiplyString
      }
    } else if self.type == 2 {
      if factor.getType() == 1 {
        return VarObject(initial_value: self.floatVal * Float(factor.getIntegerValue()))
      } else if factor.getType() == 2 {
        return VarObject(initial_value: self.floatVal * factor.getFloatValue())
      } else if factor.getType() == 3 {
        throw OperationError.multiplyString
      }
    } else {
      throw OperationError.multiplyString
    }
    throw OperationError.invalidState
  }

  public func div(divisor : VarObject) throws -> VarObject {
    if self.type == 1 {
      if divisor.getType() == 1 {
        return VarObject(initial_value: Float(self.intVal) / Float(divisor.getIntegerValue()))
      } else if divisor.getType() == 2 {
        return VarObject(initial_value: Float(self.intVal) / divisor.getFloatValue())
      } else if divisor.getType() == 3 {
        throw OperationError.divideString
      }
    } else if self.type == 2 {
      if divisor.getType() == 1 {
        return VarObject(initial_value: self.floatVal / Float(divisor.getIntegerValue()))
      } else if divisor.getType() == 2 {
        return VarObject(initial_value: self.floatVal / divisor.getFloatValue())
      } else if divisor.getType() == 3 {
        throw OperationError.divideString
      }
    } else {
      throw OperationError.divideString
    }
    throw OperationError.invalidState
  }

  public func mod(divisor : VarObject) throws -> VarObject {
    if self.type == 1 {
      if divisor.getType() == 1 {
        return VarObject(initial_value: self.intVal % divisor.getIntegerValue())
      } else if divisor.getType() == 2 {
        //return VarObject(initial_value: Float(self.intVal) % divisor.getFloatValue())
        throw OperationError.moduloDivideString
      } else if divisor.getType() == 3 {
        throw OperationError.moduloDivideString
      }
    } else if self.type == 2 {
      if divisor.getType() == 1 {
        //return VarObject(initial_value: self.floatVal % Float(divisor.getIntegerValue()))
        throw OperationError.moduloDivideString
      } else if divisor.getType() == 2 {
        //return VarObject(initial_value: self.floatVal % divisor.getFloatValue())
        throw OperationError.moduloDivideString
      } else if divisor.getType() == 3 {
        throw OperationError.moduloDivideString
      }
    } else {
      throw OperationError.moduloDivideString
    }
    throw OperationError.invalidState
  }

  public func exp(exponent : VarObject) throws -> VarObject {
    throw OperationError.raiseStringToPower
  }

  public func fac() throws -> VarObject {
    throw OperationError.factorialOfString
  }
}
