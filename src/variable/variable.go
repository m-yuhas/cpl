package variable

type Variable struct {
  Type_code uint8
  IntVal int64
}

func (v *Variable) SetType(type_code uint8) {
  v.Type_code = type_code
}

func (v *Variable) SetValue(value int64) {
  v.IntVal = value
}

func (v *Variable) Add(addend Variable) Variable {
  returnVar := Variable{}
  returnVar.SetType(1)
  returnVar.SetValue(v.IntVal+addend.IntVal)
  return returnVar
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
