// An algebraic parser in swift, relies on Access to a dictionary of Variables
// Definition of Variable object available in varObject.swift
// (C) 2016 Michael Yuhas

func parseExpression(expression : String) -> VarObject {
  addSubArray = expression.components(separatedBy : "+")
  var parenthCount = 0
  var addSubIndex = -1
  var mulDivIndex = -1
  var expIndex = -1
  var facIndex = -1
  var optype = -1
  for var i = 0; i < expression.characters.count; i+=1 {
    if expression[i] == '(' {
      parenthCount += 1
      continue
    } else if expression[i] == ')' {
      parenthCount -= 1
      continue
    } else if parenthCount = 0 {
      if expression[i] == '+' {
        addSubIndex = i
        optype = 1
        break
      } else if expression[i] == '-' {
        addSubIndex = i
        optype = 2
        break
      } else if expression[i] == '*' {
        mulDivIndex = i
        optype = 3
        break
      } else if expression[i] == '/' {
        mulDivIndex = i
        optype = 4
        break
      } else if expression[i] == '%' {
        mulDivIndex = i
        optype = 5
        break
      } else if expression[i] == '^' {
        expIndex = i
        optype = 6
        break
      } else if expression[i] == '!' {
        facIndex = i
        optype = 7
      }
    }
  }
  if ( addSubIndex != -1 ) {
    var firsthalf = expression.substring(to: i)
    var lasthalf = expression.substring(from: i+1)
    if optype == 1 {
      return parseExpression( firsthalf ) + parseExpression( lasthalf )
    }
    if optype == 2 {
      return
    }
  }
  return
}



  if ( muldivindex != -1 ) {
    printf("Multiply Divide Reached\n");
    char firsthalf[strlen(expression)];
    char lasthalf[strlen(expression)];
    strncpy(firsthalf,expression,muldivindex);
    strncpy(lasthalf,expression+muldivindex+1,strlen(expression)-muldivindex+1);
    int returnVal1;
    int returnVal2;
    int error;
    error = evaluate_int_expression( firsthalf, pointerToHashMap, &returnVal1 );
    error = evaluate_int_expression( lasthalf, pointerToHashMap, &returnVal2 );
    if ( optype == 1 ) {
      *value = returnVal1 * returnVal2;
      return error;
    } else {
      *value = returnVal1 / returnVal2;
      return error;
    }
  }
  return eval_atom( expression, pointerToHashMap, value );
}

int eval_atom( char *expression, struct Hashmap *pointerToHashMap, int *value ) {
  if ( isdigit(expression[0]) ) {
    *value = atoi(expression);
    return 0;
  } else {
    return get_int_at_key(pointerToHashMap,expression,value);
  }
}
