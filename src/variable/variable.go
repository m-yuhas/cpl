package variable

type Variable struct {
  type_code uint8
  intVal int64
}

func (v *Variable) SetType(type_code uint8) {
  v.type_code = type_code
}

func (v *Variable) SetValue(value int64) {
  v.intVal = value
}
