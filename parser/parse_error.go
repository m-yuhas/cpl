package parser

type ParseError struct {
  msg string
  line_number int64
  previous_error *ParseError
}

func (e *ParseError) Error() string {
  return e.msg + "（句" + string(line_number) + "）\n";
}

func NewParseError( m string, ln int64, pe *ParseError) ParseError {
  return ParseError{msg: m, line_number: ln, previous_error: &pe}
}
