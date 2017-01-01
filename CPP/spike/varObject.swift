// This swift file defines the Variable Class that is used for storing Variables
// (C) 2016 Michael Yuhas

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
}
