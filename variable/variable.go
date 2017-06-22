package variable

type Variable struct {
  TypeCode uint8
  BoolVal bool
  IntVal int64
  FloatVal float64
  StringVal string
  FuncVal []string
  FuncArgs []string
}

func (v *Variable) Add(addend Variable) (Variable, error) {
  returnVar := Variable{}
  switch v.TypeCode {
  case 0:
    return returnVar, errors.New("错误：变量不无初始化")
  case 1:
    switch addend.TypeCode {
    case 0:
      return returnVar, errors.New("错误：变量不无初始化")
    case 1:
      returnVar.TypeCode = 1
      returnVar.BoolVal = v.BoolVal || addend.BoolVal
        return returnVar
    case 2:
      returnVar.TypeCode = 2
      if v.BoolVal {
        returnVar.IntVal = addend.IntVal + 1
      } else {
        returnVar.IntVal = addend.IntVal
      }
      return returnVar
    case 3:
      returnVar.TypeCode = 3
      if v.BoolVal {
        returnVar.FloatVal = addend.FloatVal + 1
      } else {
        returnVar.FloatVal = addend.FloatVal
      }
      return returnVar
    case 4:
      returnVar.TypeCode = 4
      s_arr []string
      if v.BoolVal {
        s_arr[0] = "真"
      } else {
        s_arr[0] = "假"
      }
      s_arr[1] = addend.StringVal
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar
    }
  case 2:
    switch addend.TypeCode {
    case 0:
      return returnVar, errors.New("错误：变量不无初始化")
    case 1:
      returnVar.TypeCode = 2
      if addend.BoolVal {
        returnVar.IntVal = v.IntVal + 1
      } else {
        returnVar.IntVal = v.IntVal
      }
      return returnVar
    case 2:
      returnVar.TypeCode = 2
      returnVar.IntVal = v.IntVal + addend.IntVal
      return returnVar
    case 3:
      returnVar.TypeCode = 3
      returnVar.FloatVal = float64(v.IntVal) + addend.FloatVal
      return returnVar
    case 4:
      returnVar.TypeCode = 4
      s_arr []string
      s_arr[0] = strconv.Itoa(v.IntVal)
      s_arr[1] = addend.StringVal
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar
    }
  case 3:
    switch addend.TypeCode {
    case 0:
      return returnVar, errors.New("错误：变量不无初始化")
    case 1:
      returnVar.TypeCode = 3
      if addend.BoolVal {
        returnVar.FloatVal = v.FloatVal + 1
      } else {
        returnVar.FloatVal = v.FloatVal
      }
      return returnVar
    case 2:
      returnVar.typeCode = 3
      returnVar.FloatVal = v.FloatVal + float64(addend.FloatVal)
      return returnVar
    case 4:
      returnVar.TypeCode = 4
      s_arr []string
      s_arr[0] = strconv.FormatFloat(v.FloatVal,'f',-1,64)
      s_arr[1] = addend.StringVal
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar
    }
  case 4:
    switch addend.TypeCode {
    case 0:
      return returnVar, errors.New("错误：变量不无初始化")
    case 1:
      returnVar.TypeCode = 4
      s_arr []string
      s_arr[0] = v.StringVal
      if addend.BoolVal {
        s_arr[1] = "真"
      } else {
        s_arr[1] = "假"
      }
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar
    case 2:
      returnVar.TypeCode = 4
      s_arr []string
      s_arr[0] = v.StringVal
      s_arr[1] = strconv.Itoa(addend.IntVal)
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar
    case 3:
      returnVar.TypeCode = 4
      s_arr []string
      s_arr[0] = v.StringVal
      s_arr[1] = strconv.FormatFloat(v.FloatVal,'f',-1,64)
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar
    case 4:
      returnVar.TypeCode = 4
      s_arr []string
      s_arr[0] = v.StringVal
      s_arr[1] = addend.StringVal
      returnVar.StringVal = strings.Join(s_arr,"")
      return returnVar
    }
  }
  return returnVar, errors.New("错误：未知的错误")
}

func (v *Variable) Sub(subtrahend Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(1)
  returnVar.SetValue(v.IntVal-subtrahend.IntVal)
  return returnVar
}

func (v *Variable) Mul(factor Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(1)
  returnVar.SetValue(v.IntVal*factor.IntVal)
  return returnVar
}

func (v *Variable) Div(dividend Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(1)
  returnVar.SetValue(v.IntVal/dividend.IntVal)
  return returnVar
}

func (v *Variable) Mod(dividend Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(1)
  returnVar.SetValue(v.IntVal%dividend.IntVal)
  return returnVar
}

func (v *Variable) Exp(exponent Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(1)
  returnVar.SetValue(v.IntVal^exponent.IntVal)
  return returnVar
}

func (v *Variable) Eq(operand Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(2)
  returnVar.BoolVal = v.IntVal == operand.IntVal
  return returnVar
}

func (v *Variable) Lt(operand Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(2)
  returnVar.BoolVal = v.IntVal < operand.IntVal
  return returnVar
}

func (v *Variable) Gt(operand Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(2)
  returnVar.BoolVal = v.IntVal > operand.IntVal
  return returnVar
}

func (v *Variable) And(operand Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(2)
  returnVar.BoolVal = v.BoolVal && operand.BoolVal
  return returnVar
}

func (v *Variable) Or(operand Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(2)
  returnVar.BoolVal = v.BoolVal || operand.BoolVal
  return returnVar
}

func (v *Variable) Not() Variable {
  returnVar := Variable{}
  returnVar.SetType(2)
  returnVar.BoolVal = !v.BoolVal
  return returnVar
}
