/*
variable.go - 'Object' for storing and performing operations on genearic variables
(C) 2017 Michael Yuhas
*/

/*
This is part of the variable package
*/
package variable

/*
error - return type errors from operations
strings - aids in string operations
strconv - conversion to and from string type
math - operations not defined in standar go
bytes - storing variable data in a buffer
*/
import (
  "errors"
  "strings"
  "strconv"
  "math"
  //"fmt"
)

/*
Enum for Variable Type codes
NULL - null
BOOL - boolean
INT - integer
FLOAT - floating point
STRING - string TODO: Remove this and implement string as Object
FUNC - function
CLASS - class
OBJ - object
*/

type TYPE_CODE uint8

const (
  NULL TYPE_CODE = iota + 1
  BOOL
  INT
  FLOAT
  STRING
  ARRAY
  FUNC
  CLASS
  OBJ
)

/*
Variable Structure
TypeCode - type code as defined in enum
Value - array of bytes to be read out as the respective type
*/
type Variable struct {
  TypeCode TYPE_CODE
  BoolVal bool
  IntVal int64
  FloatVal float64
  StringVal string
  ArrayVal []Variable
  FuncVal []string
  FuncLines []int
  FuncArgs []string
  ClassVal []Variable
}



func (v *Variable) Add(addend Variable) (Variable, error) {
  returnVar := Variable{}
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch addend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = BOOL
      returnVar.BoolVal = v.BoolVal || addend.BoolVal
        return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      if v.BoolVal {
        returnVar.IntVal = addend.IntVal + 1
      } else {
        returnVar.IntVal = addend.IntVal
      }
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      if v.BoolVal {
        returnVar.FloatVal = addend.FloatVal + 1
      } else {
        returnVar.FloatVal = addend.FloatVal
      }
      return returnVar, nil
    case STRING:
      returnVar.TypeCode = STRING
      var s_arr []string
      if v.BoolVal {
        s_arr[0] = "真"
      } else {
        s_arr[0] = "假"
      }
      s_arr[1] = addend.StringVal
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar, nil
    }
  case INT:
    switch addend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = INT
      if addend.BoolVal {
        returnVar.IntVal = v.IntVal + 1
      } else {
        returnVar.IntVal = v.IntVal
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      returnVar.IntVal = v.IntVal + addend.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = float64(v.IntVal) + addend.FloatVal
      return returnVar, nil
    case STRING:
      returnVar.TypeCode = STRING
      var s_arr []string
      s_arr[0] = strconv.FormatInt(v.IntVal, 10)
      s_arr[1] = addend.StringVal
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar, nil
    }
  case FLOAT:
    switch addend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = FLOAT
      if addend.BoolVal {
        returnVar.FloatVal = v.FloatVal + 1
      } else {
        returnVar.FloatVal = v.FloatVal
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal + float64(addend.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal + addend.FloatVal
      return returnVar, nil
    case STRING:
      returnVar.TypeCode = STRING
      var s_arr []string
      s_arr[0] = strconv.FormatFloat(v.FloatVal,'f',-1,64)
      s_arr[1] = addend.StringVal
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar, nil
    }
  case STRING:
    switch addend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = STRING
      var s_arr []string
      s_arr[0] = v.StringVal
      if addend.BoolVal {
        s_arr[1] = "真"
      } else {
        s_arr[1] = "假"
      }
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar, nil
    case INT:
      returnVar.TypeCode = STRING
      var s_arr []string
      s_arr = append(s_arr,v.StringVal)
      s_arr = append(s_arr,strconv.FormatInt(addend.IntVal, 10))
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = STRING
      var s_arr []string
      s_arr = append(s_arr,v.StringVal)
      s_arr = append(s_arr,strconv.FormatFloat(addend.FloatVal,'f',-1,64))
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar, nil
    case STRING:
      returnVar.TypeCode = STRING
      var s_arr []string
      s_arr = append(s_arr,v.StringVal)
      s_arr = append(s_arr,addend.StringVal)
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar, nil
    }
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Sub(subtrahend Variable) (Variable, error) {
  returnVar := Variable{}
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch subtrahend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      return returnVar, errors.New("错误：变量不无初始化")
    case INT:
      returnVar.TypeCode = INT
      if v.BoolVal {
        returnVar.IntVal = 1 - subtrahend.IntVal
      } else {
        returnVar.IntVal = 0 - subtrahend.IntVal
      }
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      if v.BoolVal {
        returnVar.FloatVal = 1 - subtrahend.FloatVal
      } else {
        returnVar.FloatVal = 0 - subtrahend.FloatVal
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch subtrahend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = INT
      if subtrahend.BoolVal {
        returnVar.IntVal = v.IntVal - 1
      } else {
        returnVar.IntVal = v.IntVal
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      returnVar.IntVal = v.IntVal - subtrahend.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = float64(v.IntVal) - subtrahend.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch subtrahend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = FLOAT
      if subtrahend.BoolVal {
        returnVar.FloatVal = 1 - v.FloatVal
      } else {
        returnVar.FloatVal = 0 - v.FloatVal
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal - float64(subtrahend.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal - subtrahend.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    switch subtrahend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = STRING
      if subtrahend.BoolVal {
        returnVar.StringVal = strings.Replace(v.StringVal,"真","",-1)
      } else {
        returnVar.StringVal = strings.Replace(v.StringVal,"假","",-1)
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = STRING
      returnVar.StringVal = strings.Replace(v.StringVal,strconv.FormatInt(subtrahend.IntVal, 10),"",-1)
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = STRING
      returnVar.StringVal = strings.Replace(v.StringVal,strconv.FormatFloat(subtrahend.FloatVal,'f',-1,64),"",-1)
      return returnVar, nil
    case STRING:
      returnVar.TypeCode = STRING
      returnVar.StringVal = strings.Replace(v.StringVal,subtrahend.StringVal,"",-1)
      return returnVar, nil
    }
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Mul(factor Variable) (Variable, error) {
  returnVar := Variable{}
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch factor.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = BOOL
      returnVar.BoolVal = v.BoolVal && factor.BoolVal
      return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      if v.BoolVal {
        returnVar.IntVal = factor.IntVal
      } else {
        returnVar.IntVal = 0
      }
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      if v.BoolVal {
        returnVar.FloatVal = factor.FloatVal
      } else {
        returnVar.FloatVal = 0
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch factor.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = INT
      if factor.BoolVal {
        returnVar.IntVal = v.IntVal
      } else {
        returnVar.IntVal = 0
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      returnVar.IntVal = v.IntVal * factor.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = float64(v.IntVal) * factor.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch factor.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = FLOAT
      if factor.BoolVal {
        returnVar.FloatVal = v.FloatVal
      } else {
        returnVar.FloatVal = 0
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal * float64(factor.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal * factor.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Div(dividend Variable) (Variable, error) {
  returnVar := Variable{}
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch dividend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      return returnVar, errors.New("错误：变量不无初始化")
    case INT:
      returnVar.TypeCode = INT
      if v.BoolVal {
        returnVar.IntVal = 1 / dividend.IntVal
      } else {
        returnVar.IntVal = 0
      }
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      if v.BoolVal {
        returnVar.FloatVal = 1 / dividend.FloatVal
      } else {
        returnVar.FloatVal = 0
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch dividend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = INT
      if dividend.BoolVal {
        returnVar.IntVal = v.IntVal
      } else {
        return returnVar, errors.New("Divide by Zero Error")
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      returnVar.IntVal = v.IntVal / dividend.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = float64(v.IntVal) / dividend.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch dividend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = FLOAT
      if dividend.BoolVal {
        returnVar.FloatVal = v.FloatVal
      } else {
        return returnVar, errors.New("Error Divide By Zero")
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal / float64(dividend.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = v.FloatVal / dividend.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Mod(dividend Variable) (Variable, error) {
  returnVar := Variable{}
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch dividend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      return returnVar, errors.New("错误：变量不无初始化")
    case INT:
      returnVar.TypeCode = INT
      if v.BoolVal {
        returnVar.IntVal = 1 % dividend.IntVal
      } else {
        returnVar.IntVal = 0
      }
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      if v.BoolVal {
        returnVar.FloatVal = math.Mod(1.0,dividend.FloatVal)
      } else {
        returnVar.FloatVal = 0
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch dividend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = INT
      if dividend.BoolVal {
        returnVar.IntVal = 0
      } else {
        return returnVar, errors.New("Divide by Zero Error")
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      returnVar.IntVal = v.IntVal % dividend.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = math.Mod(float64(v.IntVal),dividend.FloatVal)
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch dividend.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = FLOAT
      if dividend.BoolVal {
        returnVar.FloatVal = 0
      } else {
        return returnVar, errors.New("Error Divide By Zero")
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = math.Mod(v.FloatVal,float64(dividend.IntVal))
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = math.Mod(v.FloatVal,dividend.FloatVal)
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Exp(exponent Variable) (Variable, error) {
  returnVar := Variable{}
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch exponent.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = BOOL
      returnVar.BoolVal = v.BoolVal != exponent.BoolVal
      return returnVar, errors.New("错误：变量不无初始化")
    case INT:
      returnVar.TypeCode = INT
      if v.BoolVal {
        returnVar.IntVal = 1
      } else {
        returnVar.IntVal = 0
      }
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      if v.BoolVal {
        returnVar.FloatVal = 1
      } else {
        returnVar.FloatVal = 0
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch exponent.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = INT
      if exponent.BoolVal {
        returnVar.IntVal = v.IntVal
      } else {
        returnVar.IntVal = 1
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = INT
      returnVar.IntVal = int64(math.Pow(float64(v.IntVal),float64(exponent.IntVal)))
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = math.Pow(float64(v.IntVal),exponent.FloatVal)
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch exponent.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.TypeCode = FLOAT
      if exponent.BoolVal {
        returnVar.FloatVal = v.FloatVal
      } else {
        returnVar.FloatVal = 1
      }
      return returnVar, nil
    case INT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = math.Pow(v.FloatVal,float64(exponent.FloatVal))
      return returnVar, nil
    case FLOAT:
      returnVar.TypeCode = FLOAT
      returnVar.FloatVal = math.Pow(v.FloatVal,exponent.FloatVal)
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Fac() (Variable, error) {
  returnVar := Variable{}
  if v.TypeCode == 2 {
    returnVar.TypeCode = INT
    returnVal := int64(1)
    for i := v.IntVal; i > 1; i++ {
      returnVal = returnVal * i
    }
    returnVar.IntVal = returnVal
    return returnVar, nil
  }
  return returnVar, errors.New("Can only do factorial on integer")
}

func (v *Variable) Eq(operand Variable) (Variable, error) {
  returnVar := Variable{}
  returnVar.TypeCode = BOOL
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.BoolVal = v.BoolVal == operand.BoolVal
        return returnVar, nil
    case INT:
      if v.BoolVal {
        returnVar.BoolVal = operand.IntVal == 1
      } else {
        returnVar.BoolVal = operand.IntVal == 0
      }
      return returnVar, nil
    case FLOAT:
      if v.BoolVal {
        returnVar.BoolVal = math.Abs(operand.FloatVal - 1) < 0.000001
      } else {
        returnVar.BoolVal = math.Abs(operand.FloatVal) < 0.000001
      }
      return returnVar, nil
    case STRING:
      if v.BoolVal {
        returnVar.BoolVal = operand.StringVal == "真"
      } else {
        returnVar.BoolVal = operand.StringVal == "假"
      }
      return returnVar, nil
    }
  case INT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.IntVal == 1
      } else {
        returnVar.BoolVal = v.IntVal == 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.IntVal == operand.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = math.Abs(float64(v.IntVal) - operand.FloatVal) < 0.000001
      return returnVar, nil
    case STRING:
      returnVar.BoolVal = strconv.FormatInt(v.IntVal, 10) == operand.StringVal
      return returnVar, nil
    }
  case FLOAT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = math.Abs(v.FloatVal - 1) < 0.000001
      } else {
        returnVar.BoolVal = math.Abs(v.FloatVal) < 0.000001
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = math.Abs(v.FloatVal - float64(operand.IntVal)) < 0.000001
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = math.Abs(v.FloatVal - operand.FloatVal) < 0.000001
      return returnVar, nil
    case STRING:
      returnVar.BoolVal = strconv.FormatFloat(v.FloatVal,'f',-1,64) == operand.StringVal
      return returnVar, nil
    }
  case STRING:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.StringVal == "真"
      } else {
        returnVar.BoolVal = v.StringVal == "假"
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = strconv.FormatInt(operand.IntVal, 10) == v.StringVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = strconv.FormatFloat(operand.FloatVal,'f',-1,64) == v.StringVal
      return returnVar, nil
    case STRING:
      returnVar.BoolVal = v.StringVal == operand.StringVal
      return returnVar, nil
    }
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Neq(operand Variable) (Variable, error) {
  returnVar := Variable{}
  returnVar.TypeCode = BOOL
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      returnVar.BoolVal = v.BoolVal != operand.BoolVal
        return returnVar, nil
    case INT:
      if v.BoolVal {
        returnVar.BoolVal = operand.IntVal != 1
      } else {
        returnVar.BoolVal = operand.IntVal != 0
      }
      return returnVar, nil
    case FLOAT:
      if v.BoolVal {
        returnVar.BoolVal = math.Abs(operand.FloatVal - 1) > 0.000001
      } else {
        returnVar.BoolVal = math.Abs(operand.FloatVal) > 0.000001
      }
      return returnVar, nil
    case STRING:
      if v.BoolVal {
        returnVar.BoolVal = operand.StringVal != "真"
      } else {
        returnVar.BoolVal = operand.StringVal != "假"
      }
      return returnVar, nil
    }
  case INT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.IntVal != 1
      } else {
        returnVar.BoolVal = v.IntVal != 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.IntVal != operand.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = math.Abs(float64(v.IntVal) - operand.FloatVal) > 0.000001
      return returnVar, nil
    case STRING:
      returnVar.BoolVal = strconv.FormatInt(v.IntVal, 10) != operand.StringVal
      return returnVar, nil
    }
  case FLOAT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = math.Abs(v.FloatVal - 1) > 0.000001
      } else {
        returnVar.BoolVal = math.Abs(v.FloatVal) > 0.000001
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = math.Abs(v.FloatVal - float64(operand.IntVal)) > 0.000001
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = math.Abs(v.FloatVal - operand.FloatVal) > 0.000001
      return returnVar, nil
    case STRING:
      returnVar.BoolVal = strconv.FormatFloat(v.FloatVal,'f',-1,64) != operand.StringVal
      return returnVar, nil
    }
  case STRING:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.StringVal != "真"
      } else {
        returnVar.BoolVal = v.StringVal != "假"
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = strconv.FormatInt(operand.IntVal, 10) != v.StringVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = strconv.FormatFloat(operand.FloatVal,'f',-1,64) != v.StringVal
      return returnVar, nil
    case STRING:
      returnVar.BoolVal = v.StringVal != operand.StringVal
      return returnVar, nil
    }
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Lt(operand Variable) (Variable, error) {
  returnVar := Variable{}
  returnVar.TypeCode = BOOL
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if v.BoolVal == false && operand.BoolVal == true {
        returnVar.BoolVal = true
      } else {
        returnVar.BoolVal = false
      }
      return returnVar, nil
    case INT:
      if v.BoolVal {
        returnVar.BoolVal = operand.IntVal > 1
      } else {
        returnVar.BoolVal = operand.IntVal > 0
      }
      return returnVar, nil
    case FLOAT:
      if v.BoolVal {
        returnVar.BoolVal = 1 < operand.FloatVal
      } else {
        returnVar.BoolVal = 0 < operand.FloatVal
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.IntVal < 1
      } else {
        returnVar.BoolVal = v.IntVal < 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.IntVal < operand.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = float64(v.IntVal) < operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.FloatVal < 1
      } else {
        returnVar.BoolVal = v.FloatVal < 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.FloatVal < float64(operand.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = v.FloatVal < operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Lte(operand Variable) (Variable, error) {
  returnVar := Variable{}
  returnVar.TypeCode = BOOL
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if v.BoolVal == false {
        returnVar.BoolVal = true
      } else {
        returnVar.BoolVal = false
      }
      return returnVar, nil
    case INT:
      if v.BoolVal {
        returnVar.BoolVal = 1 <= operand.IntVal
      } else {
        returnVar.BoolVal = 0 <= operand.IntVal
      }
      return returnVar, nil
    case FLOAT:
      if v.BoolVal {
        returnVar.BoolVal = 1 <= operand.FloatVal
      } else {
        returnVar.BoolVal = 0 <= operand.FloatVal
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.IntVal <= 1
      } else {
        returnVar.BoolVal = v.IntVal <= 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.IntVal <= operand.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = float64(v.IntVal) <= operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.FloatVal <= 1
      } else {
        returnVar.BoolVal = v.FloatVal <= 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.FloatVal <= float64(operand.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = v.FloatVal <= operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Gt(operand Variable) (Variable, error) {
  returnVar := Variable{}
  returnVar.TypeCode = BOOL
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if v.BoolVal == true && operand.BoolVal == false {
        returnVar.BoolVal = true
      } else {
        returnVar.BoolVal = false
      }
      return returnVar, nil
    case INT:
      if v.BoolVal {
        returnVar.BoolVal = 1 > operand.IntVal
      } else {
        returnVar.BoolVal = 0 > operand.IntVal
      }
      return returnVar, nil
    case FLOAT:
      if v.BoolVal {
        returnVar.BoolVal = 1 > operand.FloatVal
      } else {
        returnVar.BoolVal = 0 > operand.FloatVal
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.IntVal > 1
      } else {
        returnVar.BoolVal = v.IntVal > 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.IntVal > operand.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = float64(v.IntVal) > operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.FloatVal > 1
      } else {
        returnVar.BoolVal = v.FloatVal > 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.FloatVal > float64(operand.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = v.FloatVal > operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Gte(operand Variable) (Variable, error) {
  returnVar := Variable{}
  returnVar.TypeCode = BOOL
  switch v.TypeCode {
  case NULL:
    return returnVar, errors.New("错误：变量不无初始化")
  case BOOL:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if v.BoolVal == true {
        returnVar.BoolVal = true
      } else {
        returnVar.BoolVal = false
      }
      return returnVar, nil
    case INT:
      if v.BoolVal {
        returnVar.BoolVal = 1 >= operand.IntVal
      } else {
        returnVar.BoolVal = 0 >= operand.IntVal
      }
      return returnVar, nil
    case FLOAT:
      if v.BoolVal {
        returnVar.BoolVal = 1 >= operand.FloatVal
      } else {
        returnVar.BoolVal = 0 >= operand.FloatVal
      }
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case INT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.IntVal >= 1
      } else {
        returnVar.BoolVal = v.IntVal >= 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.IntVal >= operand.IntVal
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = float64(v.IntVal) >= operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case FLOAT:
    switch operand.TypeCode {
    case NULL:
      return returnVar, errors.New("错误：变量不无初始化")
    case BOOL:
      if operand.BoolVal {
        returnVar.BoolVal = v.FloatVal >= 1
      } else {
        returnVar.BoolVal = v.FloatVal >= 0
      }
      return returnVar, nil
    case INT:
      returnVar.BoolVal = v.FloatVal >= float64(operand.IntVal)
      return returnVar, nil
    case FLOAT:
      returnVar.BoolVal = v.FloatVal >= operand.FloatVal
      return returnVar, nil
    case STRING:
      return returnVar, errors.New("错误：变量不无初始化")
    }
  case STRING:
    return returnVar, errors.New("错误：变量不无初始化")
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) And(operand Variable) (Variable, error) {
  returnVar := Variable{}
  if v.TypeCode != 1 && operand.TypeCode != 1 {
    return returnVar, errors.New("错误：未知的错误")
  }
  returnVar.TypeCode = BOOL
  returnVar.BoolVal = v.BoolVal && operand.BoolVal
  return returnVar, nil
}

func (v *Variable) Or(operand Variable) (Variable, error) {
  returnVar := Variable{}
  if v.TypeCode != 1 && operand.TypeCode != 1 {
    return returnVar, errors.New("错误：未知的错误")
  }
  returnVar.TypeCode = BOOL
  returnVar.BoolVal = v.BoolVal || operand.BoolVal
  return returnVar, nil
}

func (v *Variable) Not() (Variable, error) {
  returnVar := Variable{}
  if v.TypeCode != 1 {
    return returnVar, errors.New("错误：未知的错误")
  }
  returnVar.TypeCode = BOOL
  returnVar.BoolVal = !v.BoolVal
  return returnVar, nil
}

func (v *Variable) ToString() (string, error) {
  switch v.TypeCode {
  case NULL:
    return "", errors.New("错误：变量有不对的类")
  case BOOL:
    if v.BoolVal {
      return "真", nil
    } else {
      return "假", nil
    }
  case INT:
    return strconv.FormatInt(v.IntVal,10), nil
  case FLOAT:
    return strconv.FormatFloat(v.FloatVal,'f',-1,64), nil
  case STRING:
    return v.StringVal, nil
  case ARRAY:
    out_string := "[ "
    for _, array_entry := range v.ArrayVal {
      array_entry_string, err := array_entry.ToString()
      //fmt.Println(array_entry)
      if err != nil {
        return "", err
      }
      out_string += array_entry_string //TODO: Improve this and do error handling
      out_string += " "
    }
    out_string += "]"
    return out_string, nil
  }
  return "", errors.New("错误：变量有不对的类")
}
